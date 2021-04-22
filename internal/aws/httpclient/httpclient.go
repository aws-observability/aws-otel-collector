package httpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

const (
	defaultMaxRetries              = 3
	defaultTimeout                 = 1 * time.Second
	defaultBackoffRetryBaseInMills = 200
	maxHttpResponseLength          = 5 * 1024 * 1024 // 5MB
)

type HttpClient struct {
	maxRetries              int
	backoffRetryBaseInMills int
	client                  *http.Client
}

func New() *HttpClient {
	return &HttpClient{
		maxRetries:              defaultMaxRetries,
		backoffRetryBaseInMills: defaultBackoffRetryBaseInMills,
		client:                  &http.Client{Timeout: defaultTimeout},
	}
}

func (h *HttpClient) backoffSleep(currentRetryCount int) {
	backoffInMillis := int64(float64(h.backoffRetryBaseInMills) * math.Pow(2, float64(currentRetryCount)))
	sleepDuration := time.Millisecond * time.Duration(backoffInMillis)
	if sleepDuration > 60*1000 {
		sleepDuration = 60 * 1000
	}
	time.Sleep(sleepDuration)
}

func (h *HttpClient) Request(endpoint string) (body []byte, err error) {
	for i := 0; i < h.maxRetries; i++ {
		body, err = h.request(endpoint)
		if err != nil {
			log.Printf("W! retry [%d/%d], unable to get http response from %s, error: %v", i, h.maxRetries, endpoint, err)
			h.backoffSleep(i)
		}
	}
	return
}

func (h *HttpClient) request(endpoint string) ([]byte, error) {
	resp, err := h.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to get response from %s, error: %v", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to get response from %s, status code: %d", endpoint, resp.StatusCode)
	}

	if resp.ContentLength >= maxHttpResponseLength {
		return nil, fmt.Errorf("get response with unexpected length from %s, response length: %d", endpoint, resp.ContentLength)
	}

	var reader io.Reader
	//value -1 indicates that the length is unknown, see https://golang.org/src/net/http/response.go
	//In this case, we read until the limit is reached
	//This might happen with chunked responses from ECS Introspection API
	if resp.ContentLength == -1 {
		reader = io.LimitReader(resp.Body, maxHttpResponseLength)
	} else {
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body from %s, error: %v", endpoint, err)
	}

	if len(body) == maxHttpResponseLength {
		return nil, fmt.Errorf("response from %s, execeeds the maximum length: %v", endpoint, maxHttpResponseLength)
	}
	return body, nil
}
