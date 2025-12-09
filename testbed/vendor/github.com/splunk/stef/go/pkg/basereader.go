package pkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/splunk/stef/go/pkg/schema"
)

type BaseReader struct {
	Source *bufio.Reader

	FixedHeader FixedHeader
	VarHeader   VarHeader
	Schema      *schema.WireSchema

	ReadBufs ReadBufs

	FrameDecoder     FrameDecoder
	FrameRecordCount uint64
	RecordCount      uint64
}

func (r *BaseReader) Init(source *bufio.Reader) error {
	r.Source = source

	if err := r.ReadFixedHeader(); err != nil {
		return err
	}

	if err := r.FrameDecoder.Init(r.Source, r.FixedHeader.Compression); err != nil {
		return err
	}

	return nil
}

func (r *BaseReader) ReadFixedHeader() error {
	hdrSignature := make([]byte, len(HdrSignature))
	_, err := r.Source.Read(hdrSignature)
	if err != nil {
		return err
	}
	if string(hdrSignature) != HdrSignature {
		return ErrInvalidHeaderSignature
	}

	contentSize, err := binary.ReadUvarint(r.Source)
	if err != nil {
		return err
	}

	if contentSize < 2 || contentSize > HdrContentSizeLimit {
		return ErrInvalidHeader
	}

	hdrContent := make([]byte, contentSize)
	_, err = r.Source.Read(hdrContent)
	if err != nil {
		return err
	}

	versionAndType := hdrContent[0]
	version := versionAndType & HdrFormatVersionMask
	if version != HdrFormatVersion {
		return ErrInvalidFormatVersion
	}

	flags := hdrContent[1]

	r.FixedHeader.Compression = Compression(flags & HdrFlagsCompressionMethod)
	switch r.FixedHeader.Compression {
	case CompressionNone, CompressionZstd:
	default:
		return ErrInvalidCompression
	}

	return nil
}

func (r *BaseReader) ReadVarHeader(ownSchema schema.WireSchema) error {
	if _, err := r.FrameDecoder.Next(); err != nil {
		return err
	}

	hdrBytes := make([]byte, r.FrameDecoder.RemainingSize())
	n, err := r.FrameDecoder.Read(hdrBytes)
	if err != nil {
		return err
	}
	hdrBytes = hdrBytes[:n]

	buf := bytes.NewBuffer(hdrBytes)
	err = r.VarHeader.Deserialize(buf)
	if err != nil {
		return err
	}

	if len(r.VarHeader.SchemaWireBytes) != 0 {
		buf = bytes.NewBuffer(r.VarHeader.SchemaWireBytes)
		r.Schema = &schema.WireSchema{}
		err = r.Schema.Deserialize(buf)
		if err != nil {
			return err
		}
		if _, err := ownSchema.Compatible(r.Schema); err != nil {
			return fmt.Errorf("schema is not compatible with BaseReader: %w", err)
		}
	}

	return nil
}

func (r *BaseReader) NextFrame() (FrameFlags, error) {
	frameFlags, err := r.FrameDecoder.Next()
	if err != nil {
		return 0, err
	}

	r.FrameRecordCount, err = binary.ReadUvarint(&r.FrameDecoder)
	if err != nil {
		return 0, err
	}

	if err := r.ReadBufs.ReadFrom(&r.FrameDecoder); err != nil {
		return 0, err
	}

	return frameFlags, nil
}
