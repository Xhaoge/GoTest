// protobuf_message.go  * Created on  2019/10/17
// Copyright (c) 2019 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package kafka

import (
	"github.com/golang/protobuf/proto"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

const (
	SearchAnalysisTopic = "SEARCH_REQ_NIGHTKING"
	VerifyAnalysisTopic = "VERIFY_DATA_NIGHTKING"
	OrderAnalysisTopic  = "ORDER_DATA_NIGHTKING"
	PayAnalysisTopic    = "PAY_DATA_NIGHTKING"
	TraceTimersTopic    = "TRACE_TIMERS"

	AiPolicyAdjustmentTopic   = "AI_POLICY_ADJUSTMENT"
	LowPriceLogTopic          = "LOW_PRICE_LOG"
	AlertTopic                = "ALERT"
	ShoppingPushTopic         = "SHOPPING_PUSH"
	spiderResponseStatusTopic = "SPIDER_REQ_TOPIC"
)

type ProtobufMessengerProducer struct {
	kafkaProducer KafkaProducer
}

func NewProtobufMessengerProducer(kp KafkaProducer) *ProtobufMessengerProducer {
	return &ProtobufMessengerProducer{kp}
}

type protobufMessenger struct {
	traceId string
	message proto.Message
}

func (pm *protobufMessenger) PK() PK {
	return PK(pm.traceId)
}

func (pm *protobufMessenger) Payload() ([]byte, error) {
	payload, err := proto.Marshal(pm.message)
	if err != nil {
		logger.WarnNt(logger.Message("Failed to marshal ProtobufMessenger %s", err.Error()))
		return nil, err
	}
	return payload, nil
}

type TraceTimerMessenger struct {
	protobufMessenger
}

func (tm *TraceTimerMessenger) Topic() Topic {
	return TraceTimersTopic
}

func (pp *ProtobufMessengerProducer) SendTraceTimers(traceId string, message proto.Message) {
	messenger := &TraceTimerMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}
	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending TraceTimers", err)
	})
}

type SearchAnalysisDataMessenger struct {
	protobufMessenger
}

func (sam *SearchAnalysisDataMessenger) Topic() Topic {
	return SearchAnalysisTopic
}

func (pp *ProtobufMessengerProducer) SendAnalysisSearchData(traceId string, message proto.Message) {
	messenger := &SearchAnalysisDataMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AnalysisSearchData", err)
	})
}

type VerifyAnalysisDataMessenger struct {
	protobufMessenger
}

func (sam *VerifyAnalysisDataMessenger) Topic() Topic {
	return VerifyAnalysisTopic
}

func (pp *ProtobufMessengerProducer) SendAnalysisVerifyData(traceId string, message proto.Message) {
	messenger := &VerifyAnalysisDataMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AnalysisVerifyData", err)
	})
}

type OrderAnalysisDataMessenger struct {
	protobufMessenger
}

func (sam *OrderAnalysisDataMessenger) Topic() Topic {
	return OrderAnalysisTopic
}

func (pp *ProtobufMessengerProducer) SendAnalysisOrderData(traceId string, message proto.Message) {
	messenger := &OrderAnalysisDataMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AnalysisOrderData", err)
	})
}

type PayAnalysisDataMessenger struct {
	protobufMessenger
}

func (payMsg *PayAnalysisDataMessenger) Topic() Topic {
	return PayAnalysisTopic
}

func (pp *ProtobufMessengerProducer) SendAnalysisPayData(traceId string, message proto.Message) {
	messenger := &PayAnalysisDataMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}
	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AnalysisPayData", err)
	})
}

type AIPolicyAdjustmentsMessenger struct {
	protobufMessenger
}

func (aim *AIPolicyAdjustmentsMessenger) Topic() Topic {
	return AiPolicyAdjustmentTopic
}

func (pp *ProtobufMessengerProducer) SendAIPolicyAdjustments(traceId string, message proto.Message) {
	messenger := &AIPolicyAdjustmentsMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AIPolicyAdjustmentsMessenger", err)
	})
}

type LowPriceLogsMessenger struct {
	protobufMessenger
}

func (lpm *LowPriceLogsMessenger) Topic() Topic {
	return LowPriceLogTopic
}

func (pp *ProtobufMessengerProducer) SendLowPriceLogs(traceId string, message proto.Message) {
	messenger := &LowPriceLogsMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending LowPriceLogsMessenger", err)
	})
}

type AlertMessenger struct {
	protobufMessenger
}

func (am *AlertMessenger) Topic() Topic {
	return AlertTopic
}

func (pp *ProtobufMessengerProducer) SendAlertMessage(traceId string, message proto.Message) {
	messenger := &AlertMessenger{
		protobufMessenger{
			traceId: traceId,
			message: message,
		},
	}
	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending AlertMessenger", err)
	})
}

type ShoppingPushMessenger struct {
	protobufMessenger
}

func (spm *ShoppingPushMessenger) Topic() Topic {
	return ShoppingPushTopic
}

func (pp *ProtobufMessengerProducer) SendShoppingPushMessage(traceId string, message proto.Message) {
	messenger := &ShoppingPushMessenger{
		protobufMessenger{
			traceId: traceId,
			message: message,
		},
	}
	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending ShoppingPushMessenger", err)
	})
}

type SpiderResponseStatusMessenger struct {
	protobufMessenger
}

func (sm *SpiderResponseStatusMessenger) Topic() Topic {
	return spiderResponseStatusTopic
}

func (pp *ProtobufMessengerProducer) SendSpiderResponseStatus(traceId string, message proto.Message) {
	messenger := &SpiderResponseStatusMessenger{
		protobufMessenger: protobufMessenger{traceId, message},
	}

	go utils.WithRecover(func() {
		pp.kafkaProducer.Send(messenger)
	}, func(err error) {
		logger.Error(traceId, "Panic while sending spiderResponseStatusMessenger", err)
	})
}
