package encoders

import (
	"github.com/splunk/stef/go/pkg"
)

type BytesEncoder struct {
	StringEncoder
}

func (e *BytesEncoder) Init(limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet) error {
	return e.StringEncoder.Init(limiter, columns)
}

func (e *BytesEncoder) Encode(val pkg.Bytes) {
	e.StringEncoder.Encode(string(val))
}

type BytesDecoder struct {
	StringDecoder
}

func (d *BytesDecoder) Continue() {
	d.StringDecoder.Continue()
}

func (d *BytesDecoder) Decode(dst *pkg.Bytes) error {
	return d.StringDecoder.Decode((*string)(dst))
}

func (d *BytesDecoder) Init(columns *pkg.ReadColumnSet) error {
	return d.StringDecoder.Init(columns)
}
