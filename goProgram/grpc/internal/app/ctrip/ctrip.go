package ctrip

import (
	"context"
	"regexp"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

const (
	search_success             int32 = 0
	search_other_error         int32 = 3
	search_request_param_error int32 = 4
	search_system              int32 = 5
	search_network_error       int32 = -1
	search_response_error      int32 = -2
)

const (
	verify_seccess     int32 = 0
	verify_other_error int32 = 3
)

const (
	internal_error int32 = 500
	outer_error    int32 = 520
)

const (
	travel_itinerary string = "T" // 行程单
)

var errCodeRegexp *regexp.Regexp
var validate *validator.Validate

type CtripApiServer struct {
	search  handler.ApiServer
	booking hagdler.ApiServer
}

func NewCtripApiServer(ctx context.Context) *CtripApiServer {
	search := handler.NewApiSearchServer(ctx)
	booking := handler.NewApiBookingServer(ctx)
	validate = validator.New()
	errCodeRegexp = regexp.MustCompile(`\d{3}`)
	return &CtripApiServer{
		search:  search,
		booking: booking,
	}
}

func (api *CtripApiServer) Search(c *gin.Context, ctx corectx.TraceableContext) (api.ApiResponse, error) {
	ctripReq := &CtripSearchRequest{}
	defer func() {
		ctx.SetValue("cid:", ctripReq.Cid)
	}()
	if err := c.BindJSON(&ctripReq); err != nil {
		return nil, nil
	}
	appendCtripAdditionaReqInfo(ctx, ctripReq)
	if !config.IsPlatformOnline(ctripReq.Cid) {
		return transformErrorSearchResponse(search_system_error, transformer.NOT_ONLINE_ERROR), nil
	}
	req, err := traceformSearchRequest(ctripReq, ctx.GetTraceId())
	if err != nil {
		return transformErrorSearchResponse(search_system_error, transformer.REQUEST_PARAM_CONVERT_ERROR), nil
	}
	res, err := api.search.Search(ctx, req)
	if err != nil {
		return transformErrorSearchResponse(search_system_error, transformer.REQUEST_PARAM_CONVERT_ERROR), nil
	}
	ProcessResponse(res)
	return traceformSearchResponse(res, ctripReq)
}

func appendCtripAdditionalReqInfo(ctx corectx.TraceableContext, ctripReq *CtripSearchRequest) {
	traceTimers := ctx.GetValue(handler.CtxTraceTimerKey)
	if traceTimers != nil {
		self := traceTimers.(*service.TraceTimers).self
		self.AddtionReqInfo.CtripChannel = ctripReq.Channel
		self.AddtionReqInfo.CtripMainChannel = ctripReq.MainChannel
		self.AddtionReqInfo.CtripSunChanneId = ctripReq.SubChannelID
	}
}

func (api *CtripApiServer) Verify(c *fin.Context, ctx corectx.TraceableContext) (api.ApiResponse, error) {

}
