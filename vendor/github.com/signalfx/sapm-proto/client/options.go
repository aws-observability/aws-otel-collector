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
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

// Option takes a reference to a Client and sets relevant config fields on it.
type Option func(*Client) error

// WithEndpoint takes an HTTP endpoint as a string in the format scheme://address:port/path and configures the
// client to export all requests to this endpoint.
func WithEndpoint(endpoint string) Option {
	return func(a *Client) error {
		a.endpoint = endpoint
		return nil
	}
}

// WithWorkers configures the client to use N number of workers.
func WithWorkers(n uint) Option {
	return func(a *Client) error {
		a.numWorkers = n
		return nil
	}
}

// WithMaxConnections allows to specify the max number of open HTTP connections the client should keep at any time.
func WithMaxConnections(n uint) Option {
	return func(a *Client) error {
		a.maxIdleCons = n
		return nil
	}
}

// WithHTTPClient allows to pass a custom HTTP Client instance to SAPM Client
func WithHTTPClient(c *http.Client) Option {
	return func(a *Client) error {
		a.httpClient = c
		return nil
	}
}

// WithAccessToken allows to pass an authentication token to the client. The auth token is set to X-SF-TOKEN HTTP header.
func WithAccessToken(t string) Option {
	return func(a *Client) error {
		a.accessToken = t
		return nil
	}
}

// WithDisabledCompression configures the client to not apply compression on the outgoing requests.
func WithDisabledCompression() Option {
	return func(a *Client) error {
		a.disableCompression = true
		return nil
	}
}

// WithCompressionMethod chooses the compression method for the outgoing requests.
// The default compression method is CompressionMethodGzip.
// This option is ignored if WithDisabledCompression() is used.
func WithCompressionMethod(compressionMethod CompressionMethod) Option {
	return func(a *Client) error {
		switch compressionMethod {
		case CompressionMethodGzip, CompressionMethodZstd:
			a.compressionMethod = compressionMethod
		default:
			return fmt.Errorf("invalid compression method %q", string(compressionMethod))
		}

		return nil
	}
}

// WithTracerProvider returns an Option to use the TracerProvider when
// creating a Tracer.
func WithTracerProvider(tracerProvider trace.TracerProvider) Option {
	return func(a *Client) error {
		a.tracerProvider = tracerProvider
		return nil
	}
}
