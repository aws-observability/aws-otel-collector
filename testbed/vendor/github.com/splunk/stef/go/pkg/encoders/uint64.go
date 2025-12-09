package encoders

import "github.com/splunk/stef/go/pkg"

type Uint64Encoder struct {
	buf       pkg.BytesWriter
	limiter   *pkg.SizeLimiter
	lastVal   uint64
	lastDelta uint64
}

func (e *Uint64Encoder) Init(limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet) error {
	e.limiter = limiter
	return nil
}

func (e *Uint64Encoder) Reset() {
	e.lastVal = 0
	e.lastDelta = 0
}

func (e *Uint64Encoder) IsEqual(val uint64) bool {
	return e.lastVal == val
}

func (e *Uint64Encoder) Encode(val uint64) {
	delta := val - e.lastVal
	e.lastVal = val

	deltaOfDelta := int64(delta - e.lastDelta)
	e.lastDelta = delta

	oldLen := len(e.buf.Bytes())
	e.buf.WriteVarint(deltaOfDelta)

	newLen := len(e.buf.Bytes())
	e.limiter.AddFrameBytes(uint(newLen - oldLen))
}

func (e *Uint64Encoder) CollectColumns(columnSet *pkg.WriteColumnSet) {
	columnSet.SetBytes(&e.buf)
}

func (e *Uint64Encoder) WriteTo(buf *pkg.BytesWriter) {
	buf.WriteBytes(e.buf.Bytes())
}

type Uint64Decoder struct {
	buf       pkg.BytesReader
	column    *pkg.ReadableColumn
	lastVal   uint64
	lastDelta uint64
}

func (d *Uint64Decoder) Continue() {
	d.buf.Reset(d.column.Data())
}

func (d *Uint64Decoder) Decode(dst *uint64) error {
	tsDeltaOfDelta, err := d.buf.ReadVarint()
	if err != nil {
		return err
	}

	tsDelta := d.lastDelta + uint64(tsDeltaOfDelta)
	d.lastDelta = tsDelta

	d.lastVal += tsDelta

	*dst = d.lastVal
	return nil
}

func (d *Uint64Decoder) Reset() {
	d.lastVal = 0
	d.lastDelta = 0
}

func (d *Uint64Decoder) Init(columns *pkg.ReadColumnSet) error {
	d.column = columns.Column()
	return nil
}
