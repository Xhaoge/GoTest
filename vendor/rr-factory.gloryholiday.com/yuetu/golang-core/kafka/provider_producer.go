package kafka

import (
	"encoding/json"
	"time"

	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

const (
	MessageTypeKey string = "messageType"
	OperationKey   string = "operation"
	QueryUrlKey    string = "queryurl"
)

const (
	Operation_search = "search"
	Operation_verify = "verify"
	Operation_order  = "order"
	Operation_pay    = "pay"
)

type AnalysisDataMessenger struct {
	TraceId       string
	AnalysisTopic string
	Data          *AnalysisReqData
}

type AnalysisBaseReqData struct {
	RequestStartAt  time.Time
	RequestFinishAt time.Time
	ResponseCode    int32
}

type AnalysisReqData struct {
	AnalysisBaseReqData
	RequestPayload  interface{}
	ResponsePayload interface{}
}

func (adm *AnalysisDataMessenger) Topic() Topic {
	return Topic(adm.AnalysisTopic)
}

func (adm *AnalysisDataMessenger) PK() PK {
	return PK(adm.TraceId)
}

func (adm *AnalysisDataMessenger) Payload() ([]byte, error) {
	res, err := json.Marshal(adm.Data)
	if err != nil {
		logger.WarnNt("marshal AnalysisReqData to json failed:" + err.Error())
		return nil, err
	}
	return res, nil
}

type ProviderProducerConfig struct {
	ApplicationName string
}

func NewAnalysisBaseReqData(request interface{}) *AnalysisReqData {
	return &AnalysisReqData{
		AnalysisBaseReqData: AnalysisBaseReqData{
			RequestStartAt: time.Now(),
		},
		RequestPayload: request,
	}
}

func (ar *AnalysisReqData) SetResponse(response interface{}, responseCode int32) {
	ar.RequestFinishAt = time.Now()
	ar.ResponseCode = responseCode
	ar.ResponsePayload = response
}
