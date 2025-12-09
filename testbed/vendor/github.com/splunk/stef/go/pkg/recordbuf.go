package pkg

import (
	"encoding/binary"
	"io"
)

type WriteColumnSet struct {
	data       []byte
	subColumns []*WriteColumnSet
}

func (s *WriteColumnSet) TotalCount() uint {
	c := uint(0)
	for _, column := range s.subColumns {
		c += column.TotalCount()
	}
	return c + 1
}

func (s *WriteColumnSet) AddSubColumn() *WriteColumnSet {
	s.subColumns = append(s.subColumns, &WriteColumnSet{})
	return s.subColumns[len(s.subColumns)-1]
}

func (s *WriteColumnSet) SetBits(b *BitsWriter) {
	// Close BitWriter to make sure all pending bits are written.
	b.Close()
	s.data = b.Bytes()

	// We got the data. Prepare the buffer for next writing cycle.
	b.Reset()
}

func (s *WriteColumnSet) SetBytes(b *BytesWriter) {
	s.data = b.Bytes()

	// We got the data. Prepare the buffer for next writing cycle.
	b.Reset()
}

func (s *WriteColumnSet) WriteSizesTo(buf *BitsWriter) {
	// Write data size
	buf.WriteUvarintCompact(uint64(len(s.data)))

	if len(s.data) == 0 {
		// If column is empty, subcolumns must be empty too.
		return
	}

	// Recursively write subcolumns
	for _, subColumn := range s.subColumns {
		subColumn.WriteSizesTo(buf)
	}
}

func (s *WriteColumnSet) WriteDataTo(buf io.Writer) error {
	// Write data
	if _, err := buf.Write(s.data); err != nil {
		return err
	}

	if len(s.data) == 0 {
		// If column is empty, subcolumns must be empty too.
		return nil
	}

	// Recursively write subcolumn data
	for _, subColumn := range s.subColumns {
		if err := subColumn.WriteDataTo(buf); err != nil {
			return err
		}
	}

	return nil
}

func (s *WriteColumnSet) PrintSchema(indent int) {
	//fmt.Printf("%s%d\n", strings.Repeat("-", indent), len(s.subColumns))
	//for _, subColumn := range s.subColumns {
	//	subColumn.PrintSchema(indent + 1)
	//}
}

// At returns the subcolumn at specified index.
func (s *WriteColumnSet) At(i int) *WriteColumnSet {
	return s.subColumns[i]
}

type WriteBufs struct {
	Columns WriteColumnSet
	tempBuf BitsWriter
	bytes   []byte
}

func (w *WriteBufs) WriteTo(buf io.Writer) error {
	w.Columns.PrintSchema(0)

	// Collect all column data sizes and counts into tempBuf
	w.tempBuf.Reset()

	w.Columns.WriteSizesTo(&w.tempBuf)
	w.tempBuf.Close()

	// Write the size of tempBuf to buf
	bufSize := uint64(len(w.tempBuf.Bytes()))
	w.bytes = w.bytes[:0]
	w.bytes = binary.AppendUvarint(w.bytes, bufSize)

	if _, err := buf.Write(w.bytes); err != nil {
		return err
	}

	// Write tempBuf content to buf
	if _, err := buf.Write(w.tempBuf.Bytes()); err != nil {
		return err
	}

	// Write column data to buf
	return w.Columns.WriteDataTo(buf)
}

type ReadableColumn struct {
	data []byte
}

func (c *ReadableColumn) Data() []byte {
	return c.data
}

// BorrowData returns the data and sets the internal data to nil.
// This allows to avoid copying the data if the caller wants to take
// exclusive ownership.
func (c *ReadableColumn) BorrowData() []byte {
	d := c.data
	c.data = nil
	return d
}

type ReadColumnSet struct {
	column     ReadableColumn
	subColumns []*ReadColumnSet
}

func (s *ReadColumnSet) Column() *ReadableColumn {
	return &s.column
}

func (s *ReadColumnSet) AddSubColumn() *ReadColumnSet {
	s.subColumns = append(s.subColumns, &ReadColumnSet{})
	return s.subColumns[len(s.subColumns)-1]
}

func (s *ReadColumnSet) SubColumnLen() int {
	return len(s.subColumns)
}

func (s *ReadColumnSet) ReadSizesFrom(buf *BitsReader) error {
	// Read data size
	dataSize := buf.ReadUvarintCompact()
	s.column.data = EnsureLen(s.column.data, int(dataSize))

	if dataSize == 0 {
		// If column is empty, subcolumns must be empty too.
		for i := 0; i < len(s.subColumns); i++ {
			s.subColumns[i].ResetData()
		}
		return nil
	}

	// Recursively read subcolumns
	for i := 0; i < len(s.subColumns); i++ {
		if err := s.subColumns[i].ReadSizesFrom(buf); err != nil {
			return err
		}
	}

	return nil
}

func (s *ReadColumnSet) ReadDataFrom(buf ByteAndBlockReader) error {
	// Read data
	if _, err := io.ReadFull(buf, s.column.data); err != nil {
		return err
	}

	// Recursively read subcolumn data
	for _, subColumn := range s.subColumns {
		if err := subColumn.ReadDataFrom(buf); err != nil {
			return err
		}
	}

	//s.readIndex = 0

	return nil
}

func (s *ReadColumnSet) PrintSchema(indent int) {
	//fmt.Printf("%s%d\n", strings.Repeat("-", indent), len(s.subColumns))
	//for _, subColumn := range s.subColumns {
	//	subColumn.PrintSchema(indent + 1)
	//}
}

func (s *ReadColumnSet) ResetData() {
	s.column.data = nil
	for i := range s.subColumns {
		s.subColumns[i].ResetData()
	}
}

type ReadBufs struct {
	Columns      ReadColumnSet
	tempBuf      BitsReader
	tempBufBytes []byte
}

func (s *ReadBufs) ReadFrom(buf ByteAndBlockReader) error {
	bufSize, err := binary.ReadUvarint(buf)
	if err != nil {
		return err
	}
	s.tempBufBytes = EnsureLen(s.tempBufBytes, int(bufSize))
	if _, err := io.ReadFull(buf, s.tempBufBytes); err != nil {
		return err
	}
	s.tempBuf.Reset(s.tempBufBytes)

	if err := s.Columns.ReadSizesFrom(&s.tempBuf); err != nil {
		return err
	}

	return s.Columns.ReadDataFrom(buf)
}
