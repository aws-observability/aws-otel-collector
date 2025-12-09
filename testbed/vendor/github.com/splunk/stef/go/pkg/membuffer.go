package pkg

import (
	"encoding/binary"
	"io"
	"unsafe"
)

type BytesReader struct {
	buf       []byte
	byteIndex int
}

func (r *BytesReader) Reset(buf []byte) {
	r.buf = buf
	r.byteIndex = 0
}

func (r *BytesReader) ReadByte() (byte, error) {
	if r.byteIndex >= len(r.buf) {
		return 0, io.EOF
	}
	b := r.buf[r.byteIndex]
	r.byteIndex++
	return b, nil
}

func (r *BytesReader) ReadUvarint() (value uint64, err error) {
	val, n := binary.Uvarint(r.buf[r.byteIndex:])
	if n <= 0 {
		return 0, io.EOF
	}
	r.byteIndex += n
	return val, nil
}

func (r *BytesReader) ReadVarint() (value int64, err error) {
	x, n := binary.Uvarint(r.buf[r.byteIndex:])
	if n <= 0 {
		return 0, io.EOF
	}
	r.byteIndex += n
	return int64((x >> 1) ^ (-(x & 1))), err
}

func (r *BytesReader) ReadStringBytes(byteSize int) (string, error) {
	if len(r.buf[r.byteIndex:]) < byteSize {
		return "", io.EOF
	}
	str := string(r.buf[r.byteIndex : r.byteIndex+byteSize])
	r.byteIndex += byteSize
	return str, nil
}

func (r *BytesReader) ReadBytesMapped(byteSize int) ([]byte, error) {
	if len(r.buf[r.byteIndex:]) < byteSize {
		return nil, io.EOF
	}

	// Map instead of copying.
	mappedBuf := r.buf[r.byteIndex : r.byteIndex+byteSize]
	r.byteIndex += byteSize

	return mappedBuf, nil
}

func (r *BytesReader) ReadStringMapped(byteSize int) (string, error) {
	if byteSize == 0 {
		return "", nil
	}

	if r.byteIndex+byteSize > len(r.buf) {
		return "", io.EOF
	}

	// Map instead of copying.
	str := unsafe.String(&r.buf[r.byteIndex], byteSize)
	r.byteIndex += byteSize

	return str, nil
}

func (r *BytesReader) MapBytesFromMemBuf(src *BytesReader, byteSize int) error {
	buf, err := src.ReadBytesMapped(byteSize)
	if err != nil {
		return err
	}

	r.buf = buf
	r.byteIndex = 0
	return nil
}

type BytesWriter struct {
	buf       []byte
	byteIndex int
}

func NewBytesWriter(cap int) BytesWriter {
	return BytesWriter{buf: make([]byte, 0, cap)}
}

func (w *BytesWriter) WriteByte(b byte) {
	w.buf = append(w.buf, b)
}

func (w *BytesWriter) WriteBytes(bytes []byte) {
	w.buf = append(w.buf, bytes...)
}

func (w *BytesWriter) WriteStringBytes(val string) {
	w.buf = append(w.buf, val...)
}

func (w *BytesWriter) WriteUvarint(value uint64) {
	w.buf = binary.AppendUvarint(w.buf, value)
}

func (w *BytesWriter) WriteVarint(x int64) {
	ux := uint64((x >> 63) ^ (x << 1))
	w.WriteUvarint(ux)
}

func (w *BytesWriter) Reset() {
	w.buf = w.buf[:0]
	w.byteIndex = 0
}

func (w *BytesWriter) ResetAndReserve(len int) {
	needCap := len + 8
	if cap(w.buf) < needCap {
		w.buf = append(w.buf[0:cap(w.buf)], make([]byte, int(needCap)-cap(w.buf))...)
	}
	w.buf = w.buf[0:int(len)]
	w.byteIndex = 0
}

func (w *BytesWriter) MapBytesToBitsReader(dest *BitsReader, byteSize int) error {
	if len(w.buf[w.byteIndex:]) < byteSize {
		return io.EOF
	}

	// Map instead of copying.
	dest.buf = w.buf[w.byteIndex : w.byteIndex+byteSize]
	dest.byteIndex = 0
	dest.availBitCount = 0

	w.byteIndex += byteSize
	return nil
}

func (w *BytesWriter) Bytes() []byte {
	return w.buf
}
