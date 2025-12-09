package encoders

import (
	"github.com/splunk/stef/go/pkg"
)

// StringEncoder implements a basic string encoder.
// It encodes strings by writing their length followed by their bytes.
// For dictionary-based encoding, see StringDictEncoder.
type StringEncoder struct {
	buf     pkg.BytesWriter
	limiter *pkg.SizeLimiter
}

func (e *StringEncoder) Init(limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet) error {
	e.limiter = limiter
	return nil
}

func (e *StringEncoder) Encode(val string) {
	oldLen := len(e.buf.Bytes())
	strLen := len(val)
	e.buf.WriteVarint(int64(strLen))
	e.buf.WriteStringBytes(val)
	newLen := len(e.buf.Bytes())
	e.limiter.AddFrameBytes(uint(newLen - oldLen))
}

func (e *StringEncoder) CollectColumns(columnSet *pkg.WriteColumnSet) {
	columnSet.SetBytes(&e.buf)
}

func (e *StringEncoder) WriteTo(buf *pkg.BytesWriter) {
	buf.WriteBytes(e.buf.Bytes())
}

func (e *StringEncoder) Reset() {
}

// StringDecoder implements a basic string decoder.
type StringDecoder struct {
	buf    pkg.BytesReader
	column *pkg.ReadableColumn
}

func (d *StringDecoder) Continue() {
	// Borrow data from the column for exclusive ownership.
	// This allows using ReadStringMapped to avoid copying
	// during decoding.
	d.buf.Reset(d.column.BorrowData())
}

func (d *StringDecoder) Reset() {}

func (d *StringDecoder) Decode(dst *string) error {
	varint, err := d.buf.ReadVarint()
	if err != nil {
		return err
	}

	if varint >= 0 {
		strLen := int(varint)
		// Use ReadStringMapped to avoid copying.
		*dst, err = d.buf.ReadStringMapped(strLen)
		if err != nil {
			return err
		}
		return nil
	}
	return ErrInvalidRefNum
}

func (d *StringDecoder) Init(columns *pkg.ReadColumnSet) error {
	d.column = columns.Column()
	return nil
}
