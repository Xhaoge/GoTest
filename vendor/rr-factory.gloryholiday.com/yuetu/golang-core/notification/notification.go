package notification

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"rr-factory.gloryholiday.com/yuetu/golang-core/config"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	coordinator "rr-factory.gloryholiday.com/yuetu/golang-core/notification/coordinator"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

type eventHandler func() bool
type eventCallback func(bool)

type subscription struct {
	id             string
	event          *coordinator.Event
	handlerResult  atomic.Value
	resubscribe    bool
	eventHandler   eventHandler
	eventCallbacks []eventCallback
}

func (sub *subscription) setResult(r bool) {
	sub.handlerResult.Store(r)
}

func (sub *subscription) getLatestResult() bool {
	r := sub.handlerResult.Load()
	if r == nil {
		return false
	}
	return r.(bool)
}

func (sub *subscription) recoverHandler(r interface{}) {
	event := sub.event
	switch r.(type) {
	case error:
		logger.Error(event.GetTraceId(), logger.Message("Recover from event[%s - %s] subscriber", event.Resource.String(), event.Action.String()), r.(error))
	default:
		logger.Warn(event.GetTraceId(), logger.Message("Recover from event[%s - %s] subscriber: %v", event.Resource.String(), event.Action.String(), r))
	}
}

func (sub *subscription) triggerEventHandler() {
	e := sub.event
	success := utils.Retry(fmt.Sprintf("Event[%s-%s]-handler", e.Resource.String(), e.Action.String()), sub.eventHandler, 5, 20*time.Second)
	sub.setResult(success)
	for _, callback := range sub.eventCallbacks {
		callback(success)
	}
}

func (sub *subscription) autoTriggerLoop() {
	event := sub.event
	logger.Info(event.GetTraceId(), logger.Message("Start handler watcher for event[%s - %s].", event.Resource.String(), event.Action.String()))
	for {
		time.Sleep(1 * time.Minute)
		done := make(chan struct{})
		go func() {
			defer func() {
				close(done)
				if r := recover(); r != nil {
					sub.recoverHandler(r)
				}
			}()

			sub.triggerEventHandler()
		}()

		<-done
	}
}

func (sub *subscription) subscribe() {
	event := sub.event
	for {
		done := make(chan struct{})
		go func() {
			defer func() {
				close(done)
				if config.IsProd() {
					if r := recover(); r != nil {
						sub.recoverHandler(r)
					}
				}
			}()

			var clientConn *grpc.ClientConn
			var subscribeClient coordinator.NotificationService_SubscribeClient
			coordinatorAddr := config.GetString("coordinator.service")

			defer func() {
				if clientConn != nil {
					clientConn.Close()
				}
			}()

			if len(coordinatorAddr) == 0 {
				logger.Fatal("Missing config [coordinator.service]")
			}

			if sub.resubscribe {
				// Only trigger event handler when re-subscribe
				sub.triggerEventHandler()
				sub.resubscribe = true
			}
			if success := utils.Retry(fmt.Sprintf("Subscribe Event[%s-%s]", event.Resource.String(), event.Action.String()), func() bool {
				conn, err := grpc.Dial(coordinatorAddr, grpc.WithInsecure())
				if err != nil {
					logger.Warn(event.GetTraceId(), logger.Message("Failed to subscribe event[%s - %s]: Failed to dial coordinator service: %s", event.Resource.String(), event.Action.String(), err.Error()))
					return false
				}
				client := coordinator.NewNotificationServiceClient(conn)

				res, err := client.Subscribe(context.Background(), &coordinator.Subscriber{
					Event:      event,
					Identifier: sub.id,
					Timestamp:  utils.ParseToGoogleTimestamp(time.Now()),
				})

				if err != nil {
					logger.Info(event.GetTraceId(), logger.Message("Failed to subscribe event[%s - %s]: %s", event.Resource.String(), event.Action.String(), err.Error()))
					return false
				}
				subscribeClient = res
				clientConn = conn
				return true
			}, 5, 2*time.Second); success {
				logger.Info(event.GetTraceId(), logger.Message("Subscribed event[%s - %s], start to recvieve events", event.Resource.String(), event.Action.String()))
				for {
					e, err := subscribeClient.Recv()
					if err == io.EOF {
						logger.Warn(e.GetTraceId(), logger.Message("Coordinator stopped sending event[%s - %s], this should not happen normally", event.Resource.String(), event.Action.String()))
						break
					}
					if err != nil {
						logger.Warn(e.GetTraceId(), logger.Message("Failed to recv event[%s - %s]: %s", event.Resource.String(), event.Action.String(), err.Error()))
						break
					}
					logger.Info(e.GetTraceId(), logger.Message("Recv event[%s - %s], trigger event handler", e.Resource.String(), e.Action.String()))
					sub.triggerEventHandler()
				}
			}

			logger.Warn(event.GetTraceId(), logger.Message("Event[%s - %s] subscriber finished, typically caused by coordinator side, please try re-subscribe", event.Resource.String(), event.Action.String()))
		}()
		<-done

		time.Sleep(1 * time.Minute)
	}
}

func SubscribeEvent(id string, event *coordinator.Event, eh eventHandler, callbacks ...eventCallback) {
	sub := &subscription{
		id:             id,
		event:          event,
		eventHandler:   eh,
		eventCallbacks: callbacks,
	}
	sub.setResult(true) // Do not trigger event handler at startup
	go sub.autoTriggerLoop()
	go sub.subscribe()
}

const maxRetryCount = 5

func newPbMessage(resource coordinator.Resource) *coordinator.Event {
	return &coordinator.Event{
		Resource:  resource,
		Action:    coordinator.Action_CHANGE,
		Timestamp: utils.ParseToGoogleTimestamp(time.Now()),
		TraceId:   utils.UUID(),
	}
}

func NotifyResourceUpdateEvent(resource coordinator.Resource, client coordinator.NotificationServiceClient) {
	go func() {
		var err error
		var response *coordinator.NotificationResponse
		event := newPbMessage(resource)
		if utils.Retry("NotifyResourceUpdateEvent", func() bool {
			timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			response, err = client.Notify(timeout, event)
			return nil != response && response.Status == coordinator.ResponseStatus_SUCCESS
		}, maxRetryCount, 10*time.Second) {
			logger.Info(event.GetTraceId(), logger.Message("event %s notified to coordinator service, response status: %s, msg: %s", resource.String(), response.Status.String(), response.Msg))
		} else {
			logger.Error(event.GetTraceId(), logger.Message("notify event %s to coordinator service failed", resource.String()), err)
		}
	}()
}

func NewCoordinatorServiceClient(cEndpointPort string) (coordinator.NotificationServiceClient, error) {
	conn, err := grpc.Dial(cEndpointPort, grpc.WithInsecure())
	if err != nil {
		logger.ErrorNt(logger.Message("connect to coordinator failed: %s", cEndpointPort), err)
		return nil, err
	}
	logger.InfoNt(logger.Message("connect to coordinator service on :%s", cEndpointPort))
	coordinatorServiceClient := coordinator.NewNotificationServiceClient(conn)
	return coordinatorServiceClient, nil
}
