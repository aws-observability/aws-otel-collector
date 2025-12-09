package encoders

import (
	"github.com/splunk/stef/go/pkg"
)

// BytesDictEncoder implements a dictionary-based bytes encoder.
type BytesDictEncoder struct {
	StringDictEncoder
}

// BytesDictEncoderDict maintains the dictionary for BytesDictEncoder.
type BytesDictEncoderDict StringDictEncoderDict

func (e *BytesDictEncoderDict) Init(limiter *pkg.SizeLimiter) {
	(*StringDictEncoderDict)(e).Init(limiter)
}

func (e *BytesDictEncoderDict) Reset() {
	(*StringDictEncoderDict)(e).Reset()
}

func (e *BytesDictEncoder) Init(
	dict *BytesDictEncoderDict, limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet,
) error {
	return e.StringDictEncoder.Init((*StringDictEncoderDict)(dict), limiter, columns)
}

func (e *BytesDictEncoder) Encode(val pkg.Bytes) {
	e.StringDictEncoder.Encode(string(val))
}

// BytesDictDecoder implements a dictionary-based bytes decoder.
type BytesDictDecoder struct {
	StringDictDecoder
}

// BytesDictDecoderDict maintains the dictionary for BytesDictDecoder.
type BytesDictDecoderDict StringDictDecoderDict

func (d *BytesDictDecoderDict) Init() {
	(*StringDictDecoderDict)(d).Init()
}

func (d *BytesDictDecoder) Continue() {
	d.StringDictDecoder.Continue()
}

func (d *BytesDictDecoder) Decode(dst *pkg.Bytes) error {
	return d.StringDictDecoder.Decode((*string)(dst))
}

func (d *BytesDictDecoder) Init(dict *BytesDictDecoderDict, columns *pkg.ReadColumnSet) error {
	return d.StringDictDecoder.Init((*StringDictDecoderDict)(dict), columns)
}
