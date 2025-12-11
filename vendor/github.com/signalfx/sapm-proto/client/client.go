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
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	jaegerpb "github.com/jaegertracing/jaeger-idl/model/v1"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	idleConnTimeout     = 30 * time.Second
	tlsHandshakeTimeout = 10 * time.Second
	dialerTimeout       = 30 * time.Second
	dialerKeepAlive     = 30 * time.Second

	// default values
	defaultNumWorkers  uint = 8
	defaultMaxIdleCons      = 100
	defaultHTTPTimeout      = 10 * time.Second
)

type sendRequest struct {
	message []byte
	spans   int64
	batches int64
}

// CompressionMethod strings MUST match the Content-Encoding http header values.
type CompressionMethod string

const CompressionMethodGzip CompressionMethod = "gzip"
const CompressionMethodZstd CompressionMethod = "zstd"

// Client implements an HTTP sender for the SAPM protocol
type Client struct {
	tracerProvider trace.TracerProvider
	numWorkers     uint
	maxIdleCons    uint
	endpoint       string
	accessToken    string
	httpClient     *http.Client

	disableCompression bool
	// compressionMethod to use for payload. Ignored if disableCompression==true.
	compressionMethod CompressionMethod

	closeCh chan struct{}

	workers chan *worker
}

// New creates a new SAPM Client
func New(opts ...Option) (*Client, error) {

	c := &Client{
		numWorkers:        defaultNumWorkers,
		maxIdleCons:       defaultMaxIdleCons,
		compressionMethod: CompressionMethodGzip,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	if c.endpoint == "" {
		return nil, fmt.Errorf(
			"endpoint cannot be empty. WithEndpoint option must be called with a valid endpoint value",
		)
	}

	var clientTransport http.RoundTripper
	clientTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   dialerTimeout,
			KeepAlive: dialerKeepAlive,
		}).DialContext,
		MaxIdleConns:        int(c.maxIdleCons),
		MaxIdleConnsPerHost: int(c.maxIdleCons),
		IdleConnTimeout:     idleConnTimeout,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
	}

	if c.tracerProvider != nil {
		clientTransport = otelhttp.NewTransport(
			clientTransport,
			otelhttp.WithTracerProvider(c.tracerProvider),
			otelhttp.WithPropagators(otel.GetTextMapPropagator()),
		)
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout:   defaultHTTPTimeout,
			Transport: clientTransport,
		}
	}

	c.closeCh = make(chan struct{})
	c.workers = make(chan *worker, c.numWorkers)
	for i := uint(0); i < c.numWorkers; i++ {
		w, err := newWorker(
			c.httpClient, c.endpoint, c.accessToken, c.disableCompression, c.compressionMethod, c.tracerProvider,
		)
		if err != nil {
			return nil, err
		}
		c.workers <- w
	}

	return c, nil
}

// Export takes a Jaeger batches and uses one of the available workers to export it synchronously.
// It returns an error in case a request cannot be processed. It's up to the caller to retry.
func (sa *Client) Export(ctx context.Context, batches []*jaegerpb.Batch) error {
	return sa.ExportWithAccessToken(ctx, batches, "")
}

// ExportWithAccessToken takes a Jaeger batches and an SFx access token and uses one of the available
// workers to export it synchronously, preferentially using the provided token and defaulting to the
// worker's token if empty.
// It returns an error in case a request cannot be processed. It's up to the caller to retry.
func (sa *Client) ExportWithAccessToken(ctx context.Context, batches []*jaegerpb.Batch, accessToken string) error {
	_, err := sa.ExportWithAccessTokenAndGetResponse(ctx, batches, accessToken)
	return err
}

// ExportWithAccessTokenAndGetResponse does everything ExportWithAccessToken does and in addition will
// return a ResponseBody indicating the response returned from trace ingest. This can be used by consumers
// to get insights into partial drops of spans/traces from within a batch.
func (sa *Client) ExportWithAccessTokenAndGetResponse(ctx context.Context, batches []*jaegerpb.Batch, accessToken string) (*IngestResponse, error) {
	w := <-sa.workers

	ingestResponse, sendErr := w.export(ctx, batches, accessToken)
	sa.workers <- w
	if sendErr != nil {
		if sendErr.RetryDelaySeconds > 0 {
			go sa.pauseForDuration(time.Duration(sendErr.RetryDelaySeconds) * time.Second)
		}
		return ingestResponse, sendErr
	}
	return ingestResponse, nil

}

// Stop waits for all inflight requests to finish and then drains the worker pool so no more work can be done.
// It returns once all workers are drained from the pool. Note that the client can accept new requests while
// Stop() waits for other requests to finish.
func (sa *Client) Stop() {
	wg := sync.WaitGroup{}
	wg.Add(int(sa.numWorkers))
	close(sa.closeCh)
	for i := uint(0); i < sa.numWorkers; i++ {
		go func() {
			<-sa.workers
			wg.Done()
		}()
	}
	wg.Wait()
}

// pauseForDuration takes workers all workers from the pool and holds on to them until either the duration passes or
// the client is stopped.
func (sa *Client) pauseForDuration(d time.Duration) {
	if d <= 0 {
		return
	}

	done := make(chan struct{})
	workers := make([]*worker, 0, sa.numWorkers)
	ticker := time.NewTicker(d)

	// steal all workers from pool and hold them until time passes
	go func() {
	loop:
		for {
			select {
			case w := <-sa.workers:
				workers = append(workers, w)
			case <-ticker.C:
				break loop
			case <-sa.closeCh:
				break loop
			}
		}
		close(done)
	}()

	<-done
	// return held workers back to the pool
	for _, w := range workers {
		sa.workers <- w
	}
}
