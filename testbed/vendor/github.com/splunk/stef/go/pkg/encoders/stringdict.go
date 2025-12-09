package encoders

import (
	"errors"
	"unsafe"

	"github.com/splunk/stef/go/pkg"
)

// StringDictEncoder implements a dictionary-based string encoder.
// It maintains a dictionary of previously seen strings and encodes
// strings as either a reference to the dictionary or as a new string
// if it hasn't been seen before.
type StringDictEncoder struct {
	buf     pkg.BytesWriter
	dict    *StringDictEncoderDict
	limiter *pkg.SizeLimiter
}

// StringDictEncoderDict maintains the dictionary for StringDictEncoder.
type StringDictEncoderDict struct {
	m       map[string]int
	limiter *pkg.SizeLimiter
}

func (e *StringDictEncoderDict) Init(limiter *pkg.SizeLimiter) {
	e.m = make(map[string]int)
	e.limiter = limiter
}

func (e *StringDictEncoderDict) Reset() {
	e.m = make(map[string]int)
}

func (e *StringDictEncoder) Init(
	dict *StringDictEncoderDict, limiter *pkg.SizeLimiter, columns *pkg.WriteColumnSet,
) error {
	e.dict = dict
	e.limiter = limiter
	return nil
}

func (e *StringDictEncoder) Encode(val string) {
	oldLen := len(e.buf.Bytes())
	if refNum, exists := e.dict.m[val]; exists {
		// Encode as a reference to the dictionary.
		// Use negative values to distinguish from new strings.
		// RefNum is offset by -1 to ensure 0 is not used.
		e.buf.WriteVarint(int64(-refNum - 1))
		newLen := len(e.buf.Bytes())
		e.dict.limiter.AddFrameBytes(uint(newLen - oldLen))
		return
	}

	// New string, encode its length and content.
	strLen := len(val)
	if strLen > 1 {
		// Add to dictionary if length > 1 to save space.
		// Single character strings are not added to the dictionary.
		refNum := len(e.dict.m)
		e.dict.m[val] = refNum
		e.dict.limiter.AddDictElemSize(uint(strLen) + uint(unsafe.Sizeof(val)))
	}
	e.buf.WriteVarint(int64(strLen))
	e.buf.WriteStringBytes(val)
	newLen := len(e.buf.Bytes())
	e.limiter.AddFrameBytes(uint(newLen - oldLen))
}

func (e *StringDictEncoder) CollectColumns(columnSet *pkg.WriteColumnSet) {
	columnSet.SetBytes(&e.buf)
}

func (e *StringDictEncoder) WriteTo(buf *pkg.BytesWriter) {
	buf.WriteBytes(e.buf.Bytes())
}

func (e *StringDictEncoder) Reset() {
}

// StringDictDecoder implements a dictionary-based string decoder.
// It decodes strings that were encoded using StringDictEncoder.
type StringDictDecoder struct {
	buf    pkg.BytesReader
	column *pkg.ReadableColumn
	dict   *StringDictDecoderDict
}

// StringDictDecoderDict maintains the dictionary for StringDictDecoder.
type StringDictDecoderDict struct {
	dict []string
}

func (d *StringDictDecoderDict) Init() {
}

func (d *StringDictDecoderDict) Reset() {
	d.dict = d.dict[:0]
}

var ErrInvalidRefNum = errors.New("invalid RefNum, out of dictionary range")

func (d *StringDictDecoder) Continue() {
	// Borrow data from the column for exclusive ownership.
	// This allows using ReadStringMapped to avoid copying
	// during decoding.
	d.buf.Reset(d.column.BorrowData())
}

func (d *StringDictDecoder) Reset() {}

func (d *StringDictDecoder) Decode(dst *string) error {
	varint, err := d.buf.ReadVarint()
	if err != nil {
		return err
	}

	if varint >= 0 {
		// New string, read its length.
		strLen := int(varint)
		// Use ReadStringMapped to avoid copying.
		*dst, err = d.buf.ReadStringMapped(strLen)
		if err != nil {
			return err
		}
		if strLen > 1 {
			// Add to dictionary if length > 1. Single character
			// strings are not added to the dictionary.
			d.dict.dict = append(d.dict.dict, *dst)
		}
		return nil
	}

	// Reference to a previously seen string.
	refNum := int(-varint - 1)
	if refNum >= len(d.dict.dict) {
		return ErrInvalidRefNum
	}
	*dst = d.dict.dict[refNum]
	return nil
}

func (d *StringDictDecoder) Init(dict *StringDictDecoderDict, columns *pkg.ReadColumnSet) error {
	d.dict = dict
	d.column = columns.Column()
	return nil
}
