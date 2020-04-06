package dingding

// doc https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq
import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"rr-factory.gloryholiday.com/yuetu/golang-core/utils"
)

type DingDingNotifyRequest struct {
	Msgtype  string           `json:"msgtype,omitempty"`
	Text     *DingDingContent `json:"text,omitempty"`
	Markdown *Markdown        `json:"markdown,omitempty"`
	At       *DingDingAt      `json:"at,omitempty"`
}

type DingDingAt struct {
	IsAtAll bool `json:"isAtAll,omitempty"`
}

type DingDingContent struct {
	Content *DingDingNotifyContent `json:"content,omitempty"`
}

type DingDingNotifyContent struct {
	Title       string `json:"title,omitempty"`
	OldTicketNo string `json:"oldTicketNo,omitempty"`
	NewTicketNo string `json:"newTicketNo,omitempty"`
	TtsOrderNo  string `json:"ttsOrderNo,omitempty"`
	UpdateTime  string `json:"updateTime,omitempty"`
}

type DingDingResponse struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type Markdown struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

func NotifyDingDingMarkdown(isAtAll bool, url, traceId, title, text string) *DingDingResponse {
	return NotifyDingDing(url, traceId, buildDingDingRequest(isAtAll, traceId, title, "markdown", text))
}

func NotifyDingDingDefault(isAtAll bool, url, traceId, title string) *DingDingResponse {
	return NotifyDingDing(url, traceId, buildDingDingRequest(isAtAll, traceId, title, "text", ""))
}

func NotifyDingDing(url, traceId string, request *DingDingNotifyRequest) *DingDingResponse {
	ddResponse := &DingDingResponse{
		ErrCode: 500,
		ErrMsg:  fmt.Sprintf("%s initErr %s", traceId, url),
	}
	utils.Retry(fmt.Sprintf("%s notifyDD", traceId), func() bool {
		_, err := utils.PostHttpJson(context.Background(), url, request, ddResponse)
		if err != nil {
			ddResponse.ErrCode = 500
			ddResponse.ErrMsg = errors.Wrap(err, ddResponse.ErrMsg).Error()
			return false
		}
		return true
	}, 2, 2*time.Second)
	return ddResponse
}

// param: IsAtAll 是否@所有人
func buildDingDingRequest(IsAtAll bool, traceId, title, msgtype, text string) *DingDingNotifyRequest {
	ddContent := transformDingDingRequestContent(traceId, title)
	return &DingDingNotifyRequest{
		Msgtype: msgtype,
		Text: &DingDingContent{
			Content: ddContent,
		},
		Markdown: &Markdown{
			Title: title,
			Text:  text,
		},
		At: &DingDingAt{
			IsAtAll: IsAtAll,
		},
	}
}

func transformDingDingRequestContent(traceId, content string) *DingDingNotifyContent {
	return &DingDingNotifyContent{Title: fmt.Sprintf("%s traceId[%s]", content, traceId)}
}
