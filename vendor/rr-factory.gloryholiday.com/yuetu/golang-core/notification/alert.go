package notification

import (
	"rr-factory.gloryholiday.com/yuetu/golang-core/kafka"
	"rr-factory.gloryholiday.com/yuetu/golang-core/notification/monitor"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
	"time"
)

type AlertLevel int32
type AlertFor int32
type AlertObjectType int32

var alertEnabled = true

const (
	Low AlertLevel = iota
	Moderate
	Substantial
	Severe
	Critical
)

const (
	Devel AlertFor = iota
	Product
	Data
	Operation
	Boss
)

const (
	Na AlertObjectType = iota
	Provider
	Platform
)

type AlertMessage struct {
	TraceId         string
	Subject         string
	Content         string
	Level           AlertLevel
	AlertFor        []AlertFor
	AlertObjectType AlertObjectType
	AlertObjectName string
}

func transformAlertLevel(level AlertLevel) monitor.AlertMessage_AlertLevel {
	switch level {
	case Low:
		return monitor.AlertMessage_Low
	case Moderate:
		return monitor.AlertMessage_Moderate
	case Substantial:
		return monitor.AlertMessage_Substantial
	case Severe:
		return monitor.AlertMessage_Severe
	case Critical:
		return monitor.AlertMessage_Critical
	default:
		return monitor.AlertMessage_Low
	}
}

func transformAlertFor(alertFor AlertFor) monitor.AlertMessage_AlertFor {
	switch alertFor {
	case Devel:
		return monitor.AlertMessage_Devel
	case Product:
		return monitor.AlertMessage_Product
	case Data:
		return monitor.AlertMessage_Data
	case Operation:
		return monitor.AlertMessage_Operation
	case Boss:
		return monitor.AlertMessage_Boss
	default:
		return monitor.AlertMessage_Devel
	}
}

func transformAlertObjectType(alertObjectType AlertObjectType) monitor.AlertMessage_AlertObjectType {
	switch alertObjectType {
	case Provider:
		return monitor.AlertMessage_Provider
	case Platform:
		return monitor.AlertMessage_Platform
	default:
		return monitor.AlertMessage_Na
	}
}

func NewAlertMessage() *AlertMessage {
	return &AlertMessage{
		AlertFor: []AlertFor{},
	}
}

func (am *AlertMessage) WithSubject(subject string) *AlertMessage {
	am.Subject = subject
	return am
}

func (am *AlertMessage) WithLevel(level AlertLevel) *AlertMessage {
	am.Level = level
	return am
}

func (am *AlertMessage) For(alertFor AlertFor) *AlertMessage {
	am.AlertFor = append(am.AlertFor, alertFor)
	return am
}

func (am *AlertMessage) WithTid(tid string) *AlertMessage {
	am.TraceId = tid
	return am
}

func (am *AlertMessage) WithContent(content string) *AlertMessage {
	am.Content = content
	return am
}

func (am *AlertMessage) WithObject(objectType AlertObjectType, name string) *AlertMessage {
	am.AlertObjectType = objectType
	am.AlertObjectName = name
	return am
}

func SendDingdingNotification(pp *kafka.ProtobufMessengerProducer, alertMessage *AlertMessage) {
	tid := alertMessage.TraceId

	if len(tid) == 0 {
		tid = utils.UUID()
	}

	for _, alertFor := range alertMessage.AlertFor {
		message := &monitor.AlertMessage{
			Subject:    alertMessage.Subject,
			AlertFor:   transformAlertFor(alertFor),
			Level:      transformAlertLevel(alertMessage.Level),
			AlertBy:    monitor.AlertMessage_Dingding,
			Info:       alertMessage.Content,
			TraceId:    tid,
			Timestamp:  utils.ParseToGoogleTimestamp(time.Now()),
			ObjectType: transformAlertObjectType(alertMessage.AlertObjectType),
			ObjectName: alertMessage.AlertObjectName,
		}

		pp.SendAlertMessage(tid, message)
	}
}
