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

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/klauspost/compress/zstd"

	splunksapm "github.com/signalfx/sapm-proto/gen"
)

type poolObj struct {
	// tempBuf holds bytes of the message in ProtoBuf format.
	// We keep the buffer in the pool object to reduce allocations.
	tempBuf *bytes.Buffer
}

type gzipPoolObj struct {
	poolObj
	gzipReader *gzip.Reader
}

type zstdPoolObj struct {
	poolObj
	zstdReader *zstd.Decoder
}

var (
	// ErrBadContentType indicates an incompatible content type was received
	ErrBadContentType = errors.New("bad content type")

	// Pool of buffers for copying non-compressed payloads.
	bufPool = &sync.Pool{
		New: func() interface{} {
			obj := newPoolObj()
			return &obj
		},
	}

	// Pool of gzip readers.
	gzipPool = &sync.Pool{
		New: func() interface{} {
			// create a new gzip gzipReader with a bytes gzipReader and array of bytes containing only the gzip header
			reader, _ := gzip.NewReader(
				bytes.NewReader(
					[]byte{
						31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 255, 255, 1, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0,
					},
				),
			)
			return &gzipPoolObj{
				gzipReader: reader,
				poolObj:    newPoolObj(),
			}
		},
	}

	// Pool of zstd readers.
	zstdPool = &sync.Pool{
		New: func() interface{} {
			reader, _ := zstd.NewReader(
				bytes.NewReader(
					[]byte{},
				),
				// Enable sync mode to avoid concurrency overheads. With http requests
				// we don't need any concurrent decompression for individual request
				// payload, it is pointless.
				zstd.WithDecoderConcurrency(1),
			)
			return &zstdPoolObj{
				zstdReader: reader,
				poolObj:    newPoolObj(),
			}
		},
	}
)

func newPoolObj() poolObj {
	return poolObj{
		tempBuf: &bytes.Buffer{},
	}
}

// ParseTraceV2Request processes an http request request into SAPM
func ParseTraceV2Request(req *http.Request) (*splunksapm.PostSpansRequest, error) {
	var sapm = &splunksapm.PostSpansRequest{}
	if err := ParseSapmRequest(req, sapm); err != nil {
		return nil, err
	}
	return sapm, nil
}

// ParseSapmRequest parses an http request request into an SAPM compatible proto definition.
func ParseSapmRequest(req *http.Request, into proto.Unmarshaler) error {
	// content type MUST be application/x-protobuf
	if req.Header.Get(ContentTypeHeaderName) != ContentTypeHeaderValue {
		return ErrBadContentType
	}

	var reader io.Reader

	// Temporary buffer to store the message in so that we can unmarshal it.
	var tempBuf *bytes.Buffer

	// Obtain the correct object from one of the pools based on the content encoding.
	switch req.Header.Get(ContentEncodingHeaderName) {
	case GZipEncodingHeaderValue:
		obj := gzipPool.Get().(*gzipPoolObj)
		defer gzipPool.Put(obj)
		tempBuf = obj.tempBuf
		// get the gzip reader
		// reset the reader with the request body
		if err := obj.gzipReader.Reset(req.Body); err != nil {
			return err
		}
		reader = obj.gzipReader

	case ZStdEncodingHeaderValue:
		obj := zstdPool.Get().(*zstdPoolObj)
		defer zstdPool.Put(obj)
		tempBuf = obj.tempBuf
		// get the zstd reader
		// reset the reader with the request body
		if err := obj.zstdReader.Reset(req.Body); err != nil {
			return err
		}
		reader = obj.zstdReader

	case "":
		// Not compressed. Just need a temporary buffer to read the entire Body into.
		obj := bufPool.Get().(*poolObj)
		defer bufPool.Put(obj)
		tempBuf = obj.tempBuf
		reader = req.Body
	}

	// Read ProtoBuf message bytes from the Reader into a temporary buffer.
	tempBuf.Reset()
	if _, err := io.Copy(tempBuf, reader); err != nil {
		return err
	}

	// Unmarshal the message from the buffer.
	return into.Unmarshal(tempBuf.Bytes())
}
