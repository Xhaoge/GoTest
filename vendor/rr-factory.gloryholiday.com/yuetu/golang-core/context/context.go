package context

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

const (
	traceIdKey         string = "tid"
	referredTraceIdKey string = "rtid"
)

type TraceableContext interface {
	context.Context
	SetTraceId(string)
	SetReferredTraceId(string)
	GetTraceId() string
	GetReferredTraceId() string
	SetValue(string, interface{})
	GetValue(string) interface{}
	GetValueAsString(string) string
	StartTrackOperation(string)
	StopTrackOperation(string)
	TrackingMessage() string
}

type traceableCtx struct {
	parent      context.Context
	bucket      map[string]interface{}
	timeTracker *utils.TimeTracker
}

func WithCancelAndSignalHandler() (context.Context, context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-signals
		cancel()
	}()

	return ctx, cancel
}

func WithParent(parent context.Context) TraceableContext {
	return &traceableCtx{
		parent:      parent,
		bucket:      map[string]interface{}{},
		timeTracker: utils.NewTimeTracker(),
	}
}

func (tctx *traceableCtx) Deadline() (time.Time, bool) {
	return tctx.parent.Deadline()
}

func (ctx *traceableCtx) Done() <-chan struct{} {
	return ctx.parent.Done()
}

func (ctx *traceableCtx) Err() error {
	return ctx.parent.Err()
}

func (ctx *traceableCtx) Value(key interface{}) interface{} {
	return ctx.parent.Value(key)
}

func (ctx *traceableCtx) GetValue(key string) interface{} {
	value, ok := ctx.bucket[key]
	if !ok {
		value = ctx.parent.Value(key)
	}
	return value
}

func (ctx *traceableCtx) GetValueAsString(key string) string {
	v := ctx.GetValue(key)
	if v == nil {
		return ""
	}
	if sv, ok := v.(string); ok {
		return sv
	}
	return ""
}

func (ctx *traceableCtx) SetValue(key string, value interface{}) {
	ctx.bucket[key] = value
}

func (ctx *traceableCtx) GetTraceId() string {
	return ctx.GetValue(traceIdKey).(string)
}

func (ctx *traceableCtx) SetTraceId(tid string) {
	ctx.SetValue(traceIdKey, tid)
}

func (ctx *traceableCtx) GetReferredTraceId() string {
	if ctx.GetValue(referredTraceIdKey) == nil {
		return ""
	}
	return ctx.GetValue(referredTraceIdKey).(string)
}

func (ctx *traceableCtx) SetReferredTraceId(rtid string) {
	ctx.SetValue(referredTraceIdKey, rtid)
}

func (ctx *traceableCtx) StartTrackOperation(operationName string) {
	ctx.timeTracker.Start(operationName)
}

func (ctx *traceableCtx) StopTrackOperation(operationName string) {
	ctx.timeTracker.Stop(operationName)
}

func (ctx *traceableCtx) TrackingMessage() string {
	return ctx.timeTracker.String()
}

func PanicHandler(err *error, ctx TraceableContext, f func()) func() {
	traceId := ctx.GetTraceId()
	action := ctx.GetValue("action").(string)
	ctx.StartTrackOperation(action)
	return func() {
		ctx.StopTrackOperation(action)
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				*err = errors.New(t)
			case error:
				*err = t
			}

			logger.Error(traceId, logger.Message("[%s], PANIC", ctx.TrackingMessage()), *err)
			f()
		} else {
			if (*err) != nil {
				logger.Warn(traceId, logger.Message("[%s], ERROR: %s", ctx.TrackingMessage(), (*err).Error()))
			} else {
				logger.Debug(logger.Message("[%s], DONE", ctx.TrackingMessage()), zap.String("tracingID", traceId))
			}
		}
	}
}

func SimpleHandlePanic(traceId string, f func()) {
	var err error
	if r := recover(); r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		}
		logger.Error(traceId, "panic when send kafka message", err)
	}
	f()
}
