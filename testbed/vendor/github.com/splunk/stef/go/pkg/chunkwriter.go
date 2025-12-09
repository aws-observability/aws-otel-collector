package pkg

import (
	"bytes"
	"io"
)

// ChunkWriter allows to write a chunk: either a header or a frame of STEF byte stream.
type ChunkWriter interface {
	// WriteChunk writes a chunk of bytes, first the header, then the content.
	WriteChunk(header []byte, content []byte) error
}

// MemChunkWriter is a ChunkWriter that accumulates chunks in a memory buffer.
type MemChunkWriter struct {
	buf bytes.Buffer
}

func (m *MemChunkWriter) WriteChunk(header []byte, content []byte) error {
	_, err := m.buf.Write(header)
	if err != nil {
		return err
	}
	_, err = m.buf.Write(content)
	return err
}

func (m *MemChunkWriter) Bytes() []byte {
	return m.buf.Bytes()
}

// WrapChunkWriter is a ChunkWriter that writes to a wrapped io.Writer.
type WrapChunkWriter struct {
	w io.Writer
}

func NewWrapChunkWriter(w io.Writer) *WrapChunkWriter {
	return &WrapChunkWriter{w}
}

func (m *WrapChunkWriter) WriteChunk(header []byte, content []byte) error {
	_, err := m.w.Write(header)
	if err != nil {
		return err
	}
	_, err = m.w.Write(content)
	return err
}
