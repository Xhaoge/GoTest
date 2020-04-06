package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/google/uuid"
	"rr-factory.gloryholiday.com/yuetu/golang-core/config"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

type HttpOption func(*http.Request)

type httpRequest func(client *http.Client, ctx context.Context, req *http.Request, response interface{}) ([]byte, error)

func SetHttpHeader(headerKey string, headerValue string) HttpOption {
	return func(req *http.Request) {
		req.Header.Set(headerKey, headerValue)
	}
}

func doHttpJsonRequest(client *http.Client, ctx context.Context, req *http.Request, response interface{}) ([]byte, error) {
	var responseBody []byte
	resp, err := client.Do(req.WithContext(ctx))

	if err != nil {
		return responseBody, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bf := resp.Body
		responseBody, err = ioutil.ReadAll(bf)
		if err != nil {
			return responseBody, err
		}
		return responseBody, json.Unmarshal(responseBody, response)
	}

	return responseBody, errors.New(fmt.Sprintf("post json[%s] err res status[%d]", req.URL.Host, resp.StatusCode))
}

func doHttpXmlRequest(client *http.Client, ctx context.Context, req *http.Request, response interface{}) ([]byte, error) {
	var responseBody []byte
	resp, err := client.Do(req.WithContext(ctx))

	if err != nil {
		return responseBody, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bf := resp.Body
		responseBody, err = ioutil.ReadAll(bf)
		if err != nil {
			return responseBody, err
		}
		return responseBody, xml.Unmarshal(responseBody, response)
	}

	return responseBody, errors.New(fmt.Sprintf("post xml[%s] err res status[%d]", req.URL.Host, resp.StatusCode))
}

func GetHttpJson(ctx context.Context, url string, dt interface{}, options ...HttpOption) ([]byte, error) {
	transport := http.DefaultTransport.(*http.Transport)
	transport.MaxIdleConns = 4096
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	for _, option := range options {
		option(req)
	}

	return doHttpJsonRequest(client, ctx, req, dt)
}

func PostHttpXml(ctx context.Context, url string, payload interface{}, response interface{}, options ...HttpOption) ([]byte, error) {
	// data := []byte(payload.(string))  //  for test
	data, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return PostHttp(ctx, url, map[string]string{"Accept": "application/xml", "Content-Type": "application/xml"}, data, response,
		func(client *http.Client, ctx context.Context, req *http.Request, response interface{}) ([]byte, error) {
			return doHttpXmlRequest(client, ctx, req, response)
		}, options...)
}

func PostHttpJson(ctx context.Context, url string, payload interface{}, response interface{}, options ...HttpOption) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return PostHttp(ctx, url, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, data, response,
		func(client *http.Client, ctx context.Context, req *http.Request, response interface{}) ([]byte, error) {
			return doHttpJsonRequest(client, ctx, req, response)
		}, options...)
}

func PostHttp(ctx context.Context, url string, headers map[string]string, payload []byte, response interface{}, httpRequest httpRequest, options ...HttpOption) ([]byte, error) {
	transport := http.DefaultTransport.(*http.Transport)
	transport.MaxIdleConns = 4096
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{
		Transport: transport,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	for _, option := range options {
		option(req)
	}

	return httpRequest(client, ctx, req, response)
}

func UUID() string {
	return uuid.New().String()
}

func Retry(name string, f func() bool, times int, interval time.Duration) bool {
	return RetryWithCtx(context.Background(), name, f, times, interval)
}

func SyncRetry(name string, f func() bool, times int, interval time.Duration) {
	for i := 0; i < times; i++ {
		defer func() {
			if r := recover(); r != nil {
				logger.WarnNt(logger.Message("Recover from %s retry[%d/%d]: %v. Stacktrace: %s", name, i+1, times, r, debug.Stack()))
			}
		}()
		if f() {
			return
		} else {
			logger.InfoNt(logger.Message("The %d/%d try for %s failed, sleep %d seconds and try again", i+1, times, name, interval/time.Second))
			time.Sleep(interval)
		}
	}
}

func RetryWithCtx(ctx context.Context, name string, f func() bool, times int, interval time.Duration) bool {
	for i := 0; i < times; i++ {
		done := make(chan bool)
		go func(i int) {
			logger.InfoNt(logger.Message("Try %s the %d/%d time", name, i+1, times))
			defer func() {
				if r := recover(); r != nil {
					logger.WarnNt(logger.Message("Recover from %s retry[%d/%d]: %v. Stacktrace: %s", name, i+1, times, r, debug.Stack()))
					done <- false
				}
				close(done)
			}()
			if f() {
				done <- true
			} else {
				logger.InfoNt(logger.Message("The %d/%d try for %s failed, sleep %d seconds and try again", i+1, times, name, interval/time.Second))
				done <- false
			}
		}(i)
		if <-done {
			logger.InfoNt(logger.Message("%s successfully done at the %d/%d try", name, i+1, times))
			return true
		}
		if i == times-1 {
			// Avoid sleep if the last retry has been failed
			logger.InfoNt(logger.Message("All tries for %s failed ", name))
			return false
		}
		logger.InfoNt(logger.Message("The %d/%d try for %s failed, sleep %d seconds and try again", i+1, times, name, interval/time.Second))
		select {
		case <-ctx.Done():
			logger.WarnNt(logger.Message("The %d/%d try for %s failed due to context done ", i+1, times, name))
			return false
		case <-time.After(interval):
		}
	}

	return false
}

func AppId() string {
	applicationName := config.GetApplicationName()
	hostname, err := os.Hostname()

	if err != nil {
		return applicationName
	}

	if strings.HasPrefix(hostname, applicationName) {
		return hostname
	}

	return applicationName + "-" + hostname
}

func Gzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer := gzip.NewWriter(&buf)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	return buf.Bytes(), err
}

func Gunzip(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)

	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer

	_, err = result.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func TernaryOperatorInt(funResult bool, isTrue, isFalse int) int {
	return TernaryOperatorInterface(funResult, isTrue, isFalse).(int)
}

func TernaryOperatorFloat64(funResult bool, isTrue, isFalse float64) float64 {
	return TernaryOperatorInterface(funResult, isTrue, isFalse).(float64)
}

func TernaryOperatorString(funResult bool, isTrue, isFalse string) string {
	return TernaryOperatorInterface(funResult, isTrue, isFalse).(string)
}

func TernaryOperatorInterface(funResult bool, isTrue, isFalse interface{}) interface{} {
	if funResult {
		return isTrue
	}
	return isFalse
}

func WithRecover(fn func(), panicHandler func(err error)) {
	defer func() {
		var err error
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = fmt.Errorf("Recover result: %v", r)
			}
			panicHandler(err)
		}
	}()

	fn()
}
