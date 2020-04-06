package kafka

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	pb "rr-factory.gloryholiday.com/yuetu/golang-core/kafka/bootybay"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

type Payloader func() ([]byte, error)
type MessageType int32

const (
	traceMessageTopic string = "TRACE_MESSAGE"
)
const (
	REQUEST MessageType = iota
	RESPONSE
	LOG
	OTHERS
)

func NewJsonPayloader(payload interface{}) Payloader {
	return func() ([]byte, error) {
		return json.Marshal(payload)
	}
}

func NewProtobufPayloader(message proto.Message) Payloader {
	return func() ([]byte, error) {
		return proto.Marshal(message)
	}
}

func NewBytePayloader(bytes []byte, err error) Payloader {
	return func() ([]byte, error) {
		return bytes, err
	}
}

type TraceMessenger struct {
	TraceId         string
	ReferredTraceId string
	Application     string
	MessageType     MessageType
	Operation       string
	Payloader       Payloader
	Removable       bool
	Tags            map[string][]byte
}

func (ptm *TraceMessenger) Topic() Topic {
	return Topic(traceMessageTopic)
}

func (ptm *TraceMessenger) PK() PK {
	return PK(ptm.TraceId)
}

func (ptm *TraceMessenger) Payload() ([]byte, error) {
	payload, err := ptm.Payloader()
	if err != nil {
		logger.WarnNt(logger.Message("Failed to marshal %s %s %d: %v", ptm.Application, ptm.Operation, ptm.MessageType, err))
		payload = []byte(err.Error())
	}

	gzippedPayload, err := utils.Gzip(payload)
	if err != nil {
		logger.Warn(ptm.TraceId, logger.Message("gzip %s res payload failed", ptm.Operation))
		return nil, fmt.Errorf("gzip %s res payload failed", ptm.Operation)
	}

	tags := map[string][]byte{
		"gzipped": []byte("t"),
	}
	if ptm.Removable || ptm.Operation == "search" {
		tags["removable"] = []byte("0")
	}

	for k, v := range ptm.Tags {
		tags[k] = v
	}

	message := &pb.TraceMessage{
		TraceId:         ptm.TraceId,
		ReferredTraceId: ptm.ReferredTraceId,
		Application:     ptm.Application,
		MessageType:     pb.TraceMessage_MessageType(ptm.MessageType),
		Operation:       ptm.Operation,
		Timestamp:       utils.ParseToGoogleTimestamp(time.Now()),
		Payload:         gzippedPayload,
		Tags:            tags,
		Gzipped:         true,
	}

	p, err := proto.Marshal(message)
	if err != nil {
		logger.ErrorNt("Failed to marsh TraceMessage", err)
		return nil, fmt.Errorf("failed to marsh TraceMessage")
	}

	return p, nil
}
