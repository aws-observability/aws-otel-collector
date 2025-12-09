package encoders

import "github.com/splunk/stef/go/pkg"

type BoolEncoder struct {
	buf     pkg.BitsWriter
	limiter *pkg.SizeLimiter
}

func (e *BoolEncoder) Init(limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet) error {
	e.limiter = limiter
	return nil
}

func (e *BoolEncoder) Reset() {
}

func (e *BoolEncoder) Encode(val bool) {
	var v uint
	if val {
		v = 1
	}
	e.buf.WriteBit(v)

	e.limiter.AddFrameBits(1)
}

func (e *BoolEncoder) CollectColumns(columnSet *pkg.WriteColumnSet) {
	columnSet.SetBits(&e.buf)
}

func (e *BoolEncoder) WriteTo(buf *pkg.BytesWriter) {
	buf.WriteBytes(e.buf.Bytes())
}

type BoolDecoder struct {
	buf    pkg.BitsReader
	column *pkg.ReadableColumn
}

func (d *BoolDecoder) Continue() {
	d.buf.Reset(d.column.Data())
}

func (d *BoolDecoder) Decode(dst *bool) error {
	bit := d.buf.ReadBit()
	if bit == 0 {
		*dst = false
	} else {
		*dst = true
	}
	return nil
}

func (d *BoolDecoder) Reset() {
}

func (d *BoolDecoder) Init(columns *pkg.ReadColumnSet) error {
	d.column = columns.Column()
	return nil
}
