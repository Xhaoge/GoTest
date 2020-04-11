package myhttp

import (
	"fmt"
	"testing"
)


func TestGetUrlBuild(t *testing.T) {
	h := NewHttpSend("http://prod-restful-api.gloryholiday.com/nightking/exchangeRate")
	data := map[string]string{"providerName":"sscts","cid":"ttn","originalCode":"USD","targetCode":"CNY"}
	h.Url = GetUrlBuild(h.Url,data)
	fmt.Println("url: ",h.Url)
	resp,_ := h.Get()
	fmt.Println("resp: ",string(resp))
}

func TestHttpSend_Post(t *testing.T) {
	h := NewHttpSend("http://test-restful-api.gloryholiday.com/currencyservice/getCurrency")
	data := map[string]string{
		"originalCode":"USD",
		"targetCode":"CNY",
		"publish_timestamp": "2020-03-28T00:00:00Z",
	}
	h.SetSendType(SENDTYPE_JSON)
	h.SetBody(data)
	resp, _ := h.Post()
	fmt.Println("resp: ",string(resp))
}