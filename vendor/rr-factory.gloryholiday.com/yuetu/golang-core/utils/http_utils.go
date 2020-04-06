package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	transport := http.DefaultTransport.(*http.Transport)
	transport.MaxIdleConns = 1014
	transport.MaxIdleConnsPerHost = 512
	transport.TLSHandshakeTimeout = time.Second * 15
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httpClient = &http.Client{
		Transport: transport,
		Timeout:   time.Second * 300,
	}
}

type RequestType int

const (
	GET RequestType = iota
	POST
)

type MsgTransform interface {
	Encode(v interface{}) ([]byte, error)
	Decode(bytes []byte, v interface{}) error
}

func NewJsonTransform() MsgTransform {
	return &jsonTransform{}
}

func NewXmlTransform() MsgTransform {
	return &xmlTransform{}
}

func NewUrlMapParamTransform() MsgTransform {
	return &urlMapParamTransform{}
}

type jsonTransform struct{}

func (jsonT *jsonTransform) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonT *jsonTransform) Decode(rsbs []byte, v interface{}) error {
	return json.Unmarshal(rsbs, v)
}

type xmlTransform struct{}

func (jsonT *xmlTransform) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (jsonT *xmlTransform) Decode(rsbs []byte, v interface{}) error {
	return xml.Unmarshal(rsbs, v)
}

type bytesTransform struct{}

func NewBytesTrasform() MsgTransform {
	return &bytesTransform{}
}

func (bt *bytesTransform) Encode(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

func (bt *bytesTransform) Decode(rsbs []byte, v interface{}) error {
	return nil
}

type urlMapParamTransform struct{}

func (bt *urlMapParamTransform) Encode(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("requst param wrong")
	}
	mapValue := v.(map[string]string)
	if mapValue == nil || len(mapValue) < 1 {
		return nil, errors.New("request param wrong")
	}
	paramValue := bytes.Buffer{}
	for key, value := range mapValue {
		if paramValue.Len() < 1 {
			paramValue.WriteString(key)
			paramValue.WriteString("=")
			paramValue.WriteString(value)
		} else {
			paramValue.WriteString("&")
			paramValue.WriteString(key)
			paramValue.WriteString("=")
			paramValue.WriteString(value)
		}
	}
	return paramValue.Bytes(), nil
}

func (bt *urlMapParamTransform) Decode(rsbs []byte, v interface{}) error {
	return nil
}

type HttpCall interface {
	Send() error
	GetResponse() ([]byte, interface{})
	GetRequest() ([]byte, interface{})
}

type httpCaller struct {
	Ctx            context.Context
	Url            string
	ReqBytes       []byte
	Request        interface{}
	Response       interface{}
	ResBytes       []byte
	Err            error
	RequestHandle  []HttpOption
	RType          RequestType
	ResponseDecode MsgTransform
}

func (hc *httpCaller) GetRequest() ([]byte, interface{}) {
	return hc.ReqBytes, hc.Request
}

func (hc *httpCaller) GetResponse() ([]byte, interface{}) {
	return hc.ResBytes, hc.Response
}

func NewHttpClient(ctx context.Context, url string, req interface{}, res interface{}, rt RequestType, reqDecode MsgTransform, resDecode MsgTransform, handles []HttpOption) HttpCall {
	reqBs, err := reqDecode.Encode(req)
	return &httpCaller{
		Ctx:            ctx,
		Url:            url,
		Request:        req,
		Response:       res,
		ReqBytes:       reqBs,
		Err:            err,
		ResponseDecode: resDecode,
		RequestHandle:  handles,
		RType:          rt,
	}

}

func NewJsonHttpRequest(ctx context.Context, url string, req interface{}, res interface{}, rt RequestType) HttpCall {
	transformJ := NewJsonTransform()
	RequestHandle := []HttpOption{func(req *http.Request) {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}}
	return NewHttpClient(ctx, url, req, res, rt, transformJ, transformJ, RequestHandle)
}

func NewJsonResByUrlParamJsonResClientHttpRequest(ctx context.Context, url string, req interface{}, res interface{}, rt RequestType) HttpCall {
	transformJ := NewUrlMapParamTransform()
	transformRes := NewJsonTransform()
	RequestHandle := []HttpOption{func(req *http.Request) {
		req.Header.Set("Accept", "text/xml,text/javascript")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}}
	return NewHttpClient(ctx, url, req, res, rt, transformJ, transformRes, RequestHandle)
}

func NewXmlHttpClient(ctx context.Context, url string, req interface{}, res interface{}, rt RequestType) HttpCall {
	transformJ := NewXmlTransform()
	RequestHandle := []HttpOption{func(req *http.Request) {
		req.Header.Set("Accept", "application/xml")
		req.Header.Set("Content-Type", "application/xml")
	}}
	return NewHttpClient(ctx, url, req, res, rt, transformJ, transformJ, RequestHandle)
}

func (hc *httpCaller) Send() error {
	if hc.Err != nil {
		return hc.Err
	}
	req, err := hc.newHttpRequest()
	if err != nil {
		return err
	}
	for _, option := range hc.RequestHandle {
		option(req)
	}
	resp, err := httpClient.Do(req.WithContext(hc.Ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytesBody := bytes.NewBuffer(make([]byte, 4096))
		bytesBody.Reset()
		_, err := io.Copy(bytesBody, resp.Body)
		hc.ResBytes = bytesBody.Bytes()
		if err != nil {
			return err
		}
		err = hc.ResponseDecode.Decode(hc.ResBytes, hc.Response)
		if err != nil {
		}
		return err
	}
	return errors.New(fmt.Sprintf("post json[%s] err res status[%d]", req.URL.Host, resp.StatusCode))
}

func (hc *httpCaller) newHttpRequest() (*http.Request, error) {
	switch hc.RType {
	case GET:
		return http.NewRequest("GET", hc.Url, nil)
	case POST:
		return http.NewRequest("POST", hc.Url, bytes.NewBuffer(hc.ReqBytes))
	default:
		return nil, errors.New("request type is wrong")
	}

}
