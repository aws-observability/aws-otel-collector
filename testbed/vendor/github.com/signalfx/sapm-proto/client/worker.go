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
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	jaegerpb "github.com/jaegertracing/jaeger-idl/model/v1"
	"github.com/klauspost/compress/zstd"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	sapmpb "github.com/signalfx/sapm-proto/gen"
)

// IngestResponse encapsulates the body of response returned by trace ingest and any error encountered
// by the worker while reading the body.
type IngestResponse struct {
	Body []byte
	Err  error
}

type resetWriteCloser interface {
	io.WriteCloser
	Reset(w io.Writer)
}

// worker is not safe to be called from multiple goroutines. Each caller must use locks to avoid races
// and data corruption. In case a caller needs to export multiple requests at the same time, it should
// use one worker per request.
type worker struct {
	tracer             trace.Tracer
	client             *http.Client
	accessToken        string
	endpoint           string
	compressWriter     resetWriteCloser
	disableCompression bool
	compressionMethod  CompressionMethod
}

func newWorker(
	client *http.Client,
	endpoint string,
	accessToken string,
	disableCompression bool,
	compressionMethod CompressionMethod,
	tracerProvider trace.TracerProvider,
) (*worker, error) {
	if tracerProvider == nil {
		tracerProvider = trace.NewNoopTracerProvider()
	}
	w := &worker{
		tracer:             tracerProvider.Tracer("github.com/signalfx/sapm-proto/client"),
		client:             client,
		accessToken:        accessToken,
		endpoint:           endpoint,
		disableCompression: disableCompression,
		compressionMethod:  compressionMethod,
	}

	if !disableCompression {
		switch w.compressionMethod {
		case CompressionMethodGzip:
			w.compressWriter = gzip.NewWriter(nil)
		case CompressionMethodZstd:
			var err error
			w.compressWriter, err = zstd.NewWriter(
				nil,
				// Enable sync mode to avoid concurrency overheads.
				zstd.WithEncoderConcurrency(1),
			)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown compression method %v", w.compressionMethod)
		}
	}

	return w, nil
}

func (w *worker) export(ctx context.Context, batches []*jaegerpb.Batch, accessToken string) (*IngestResponse, *ErrSend) {
	ctx, span := w.tracer.Start(ctx, "export")
	defer span.End()

	var spansCount int
	for _, batch := range batches {
		spansCount += len(batch.Spans)
	}

	span.SetAttributes(attribute.Int64("spans", int64(spansCount)))
	span.SetAttributes(attribute.Int64("batches", int64(len(batches))))

	if spansCount == 0 {
		return nil, nil
	}

	sr, err := w.prepare(batches, spansCount)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "")
		return nil, &ErrSend{Err: err, Permanent: true}
	}

	responseBody, serr := w.send(ctx, sr, accessToken)
	if serr == nil {
		return responseBody, nil
	}
	span.RecordError(err)
	span.SetStatus(codes.Error, "")

	return responseBody, serr
}

func (w *worker) send(ctx context.Context, r *sendRequest, accessToken string) (*IngestResponse, *ErrSend) {
	req, err := http.NewRequestWithContext(ctx, "POST", w.endpoint, bytes.NewBuffer(r.message))
	if err != nil {
		return nil, &ErrSend{Err: err, Permanent: true}
	}
	req.Header.Add(headerContentType, headerValueXProtobuf)

	if !w.disableCompression {
		req.Header.Add(headerContentEncoding, string(w.compressionMethod))
	}

	if accessToken == "" {
		accessToken = w.accessToken
	}

	if accessToken != "" {
		req.Header.Add(headerAccessToken, accessToken)
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, &ErrSend{Err: err}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	ingestResponse := &IngestResponse{bodyBytes, err}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return ingestResponse, nil
	}

	// Drop the batch if server thinks it is malformed in some way or client is not authorized
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnauthorized {
		msg := fmt.Sprintf("server responded with: %d", resp.StatusCode)
		return ingestResponse, &ErrSend{
			Err:        fmt.Errorf("dropping request: %s", msg),
			StatusCode: http.StatusBadRequest,
			Permanent:  true,
		}
	}

	// Check if server is overwhelmed and requested to pause sending for a while.
	// Pause from sending more data till the specified number of seconds in the Retry-After header.
	// Fallback to defaultRateLimitingBackoffSeconds if the header is not present
	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := defaultRateLimitingBackoffSeconds
		if val := resp.Header.Get(headerRetryAfter); val != "" {
			if seconds, err := strconv.Atoi(val); err == nil {
				retryAfter = seconds
			}
		}
		return ingestResponse, &ErrSend{
			Err:               errors.New("server responded with 429"),
			StatusCode:        resp.StatusCode,
			RetryDelaySeconds: retryAfter,
		}
	}

	// TODO: handle 301, 307, 308
	// redirects are not handled right now but should be to confirm with the spec.

	return ingestResponse, &ErrSend{
		Err:        fmt.Errorf("error exporting spans. server responded with status %d", resp.StatusCode),
		StatusCode: resp.StatusCode,
	}
}

// prepare takes a jaeger batches, converts them to a SAPM PostSpansRequest, compresses it and returns a request ready
// to be sent.
func (w *worker) prepare(batches []*jaegerpb.Batch, spansCount int) (*sendRequest, error) {
	psr := &sapmpb.PostSpansRequest{
		Batches: batches,
	}

	encoded, err := psr.Marshal()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	if w.disableCompression {
		return &sendRequest{
			message: encoded,
			batches: int64(len(batches)),
			spans:   int64(spansCount),
		}, nil
	}

	buf := bytes.NewBuffer([]byte{})
	w.compressWriter.Reset(buf)

	if _, err = w.compressWriter.Write(encoded); err != nil {
		return nil, fmt.Errorf("failed to compress request: %w", err)
	}

	if err = w.compressWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to compress request: %w", err)
	}
	sr := &sendRequest{
		message: buf.Bytes(),
		batches: int64(len(batches)),
		spans:   int64(spansCount),
	}
	return sr, nil
}
