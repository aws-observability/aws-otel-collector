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

package sapmprotocol

const (
	// TraceEndpointV2 is the endpoint used for SAPM v2 traces.  The SAPM protocol started with v2.  There is no v1.
	TraceEndpointV2 = "/v2/trace"
	// ContentTypeHeaderName is the http header name used for Content-Type
	ContentTypeHeaderName = "Content-Type"
	// ContentTypeHeaderValue is the value used for protobuf Content-Type http headers
	ContentTypeHeaderValue = "application/x-protobuf"

	// AcceptEncodingHeaderName is the http header name used for Accept-Encoding
	AcceptEncodingHeaderName = "Accept-Encoding"
	// ContentEncodingHeaderName is the http header name used for Content-Encoding
	ContentEncodingHeaderName = "Content-Encoding"
	// GZipEncodingHeaderValue is the value used for gzipped encoding http headers
	GZipEncodingHeaderValue = "gzip"
	// ZStdEncodingHeaderValue is the value used for zstd-compressed encoding http headers
	ZStdEncodingHeaderValue = "zstd"
)
