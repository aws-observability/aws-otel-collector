// Copyright 2019 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type request struct {
	r          *http.Request
	receivedAt time.Time
}

type mockTransport struct {
	sync.Mutex
	delay      time.Duration
	statusCode int
	headers    map[string]string
	err        error
	received   []*request
	body       string
}

func (m *mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if m.delay > 0 {
		time.Sleep(m.delay)
	}
	m.record(r)
	if m.err != nil {
		return nil, m.err
	}

	resp := &http.Response{
		StatusCode: m.statusCode,
		Body:       ioutil.NopCloser(strings.NewReader(m.body)),
	}
	resp.Header = http.Header{}
	for k, v := range m.headers {
		resp.Header.Add(k, v)
	}
	return resp, nil
}

func (m *mockTransport) reset(code int) {
	m.Lock()
	m.delay = time.Duration(0)
	m.statusCode = code
	m.received = []*request{}
	m.Unlock()
}

func (m *mockTransport) record(r *http.Request) {
	m.Lock()
	m.received = append(m.received, &request{r, time.Now()})
	m.Unlock()
}

func (m *mockTransport) requests() []*request {
	m.Lock()
	r := make([]*request, len(m.received))
	copy(r, m.received)
	defer m.Unlock()
	return r
}

func newMockHTTPClient(m *mockTransport) *http.Client {
	if m.statusCode == 0 {
		m.statusCode = 200
	}
	if m.received == nil {
		m.received = []*request{}
	}
	if m.headers == nil {
		m.headers = map[string]string{}
	}
	return &http.Client{Transport: m}
}
