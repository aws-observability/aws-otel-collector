package pkg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/klauspost/compress/zstd"
)

type FrameEncoder struct {
	dest             ChunkWriter
	frameContent     io.Writer
	uncompressedSize int
	compressedBuf    bytes.Buffer
	compression      Compression
	compressor       *zstd.Encoder
	hdrByte          FrameFlags
}

func (e *FrameEncoder) Init(dest ChunkWriter, compr Compression) error {
	e.compression = compr
	e.dest = dest
	switch e.compression {
	case CompressionNone:
		e.frameContent = &e.compressedBuf

	case CompressionZstd:
		var err error
		e.compressor, err = zstd.NewWriter(&e.compressedBuf)
		if err != nil {
			return err
		}
		e.frameContent = e.compressor

	default:
		return fmt.Errorf("unknown compression: %d", e.compression)
	}
	return nil
}

func (e *FrameEncoder) OpenFrame(resetFlags FrameFlags) {
	e.hdrByte = resetFlags
	if resetFlags&RestartCompression != 0 {
		e.compressor.Reset(&e.compressedBuf)
	}
}

func (e *FrameEncoder) CloseFrame() error {
	var frameHdrArr [1 + binary.MaxVarintLen64*2]byte
	frameHdr := frameHdrArr[:0]
	frameHdr = append(frameHdr, byte(e.hdrByte))

	frameHdr = binary.AppendUvarint(frameHdr, uint64(e.uncompressedSize))

	var compressedSize int
	if e.compression == CompressionZstd {
		err := e.compressor.Flush()
		if err != nil {
			return err
		}
		compressedSize = e.compressedBuf.Len()
		frameHdr = binary.AppendUvarint(frameHdr, uint64(compressedSize))
	}
	if err := e.dest.WriteChunk(frameHdr, e.compressedBuf.Bytes()); err != nil {
		return err
	}
	e.compressedBuf.Reset()
	e.uncompressedSize = 0
	return nil
}

func (e *FrameEncoder) Write(p []byte) (n int, err error) {
	n, err = e.frameContent.Write(p)
	e.uncompressedSize += n
	return n, err
}

func (e *FrameEncoder) UncompressedSize() int {
	return e.uncompressedSize
}

type ByteAndBlockReader interface {
	ReadByte() (byte, error)
	Read(p []byte) (n int, err error)
}

type FrameDecoder struct {
	src                       ByteAndBlockReader
	compression               Compression
	uncompressedSize          uint64
	ofs                       int
	frameContentSrc           ByteAndBlockReader
	decompressedContentReader *bufio.Reader
	decompressor              *zstd.Decoder
	chunkReader               chunkedReader
	flags                     FrameFlags
	frameLoaded               bool
	notFirstFrame             bool
}

type chunkedReader struct {
	src       ByteAndBlockReader
	limit     int64
	nextChunk func() error
}

func (r *chunkedReader) readByte() (byte, error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}
	b, err := r.src.ReadByte()
	r.limit -= 1
	return b, err
}

func (r *chunkedReader) ReadByte() (byte, error) {
loop:
	for {
		b, err := r.readByte()
		if err == nil {
			return b, err
		}
		if err == io.EOF {
			err = r.nextChunk()
			if err != nil {
				return 0, err
			}
			goto loop
		}
		return 0, err
	}
}

func (r *chunkedReader) readBlock(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}
	n, err = r.src.Read(p)
	r.limit -= int64(n)
	return
}

func (r *chunkedReader) Read(p []byte) (n int, err error) {
loop:
	for {
		n, err := r.readBlock(p)
		if err == nil {
			return n, err
		}
		if err == io.EOF {
			err = r.nextChunk()
			if err != nil {
				return 0, err
			}
			goto loop
		}
		return 0, err
	}
}

func (r *chunkedReader) Init(src ByteAndBlockReader) {
	r.src = src
}

var EndOfFrame = errors.New("end of frame")

const readBufSize = 64 * 1024

func (d *FrameDecoder) Init(src ByteAndBlockReader, compression Compression) error {
	d.src = src
	d.compression = compression
	d.chunkReader.Init(src)
	d.chunkReader.nextChunk = d.nextFrame

	switch d.compression {
	case CompressionNone:
		d.frameContentSrc = &d.chunkReader

	case CompressionZstd:
		var err error
		d.decompressor, err = zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
		if err != nil {
			return err
		}
		d.decompressedContentReader = bufio.NewReaderSize(d.decompressor, readBufSize)
		d.frameContentSrc = d.decompressedContentReader

	default:
		return fmt.Errorf("unknown compression: %d", d.compression)
	}
	return nil
}

func (d *FrameDecoder) nextFrame() error {
	hdrByte, err := d.src.ReadByte()
	if err != nil {
		return err
	}

	d.flags = FrameFlags(hdrByte)
	if d.flags|FrameFlagsMask != FrameFlagsMask {
		return errors.New("invalid frame flags")
	}

	uncompressedSize, err := binary.ReadUvarint(d.src)
	if err != nil {
		return err
	}

	d.uncompressedSize = uncompressedSize

	var compressedSize uint64
	if d.compression != CompressionNone {
		compressedSize, err = binary.ReadUvarint(d.src)
		if err != nil {
			return err
		}
		d.chunkReader.limit = int64(compressedSize)

		if !d.notFirstFrame || d.flags&RestartCompression != 0 {
			d.notFirstFrame = true
			if err := d.decompressor.Reset(&d.chunkReader); err != nil {
				return err
			}
		}

		d.decompressedContentReader.Reset(d.decompressor)
	} else {
		compressedSize = uncompressedSize
		d.chunkReader.limit = int64(uncompressedSize)
	}

	d.frameLoaded = true
	d.ofs = 0
	return nil

}

func (d *FrameDecoder) Next() (FrameFlags, error) {
	for d.uncompressedSize > 0 {
		var tmp [4096]byte
		readSize := int(d.uncompressedSize)
		if readSize > len(tmp) {
			readSize = len(tmp)
		}
		n, err := d.frameContentSrc.Read(tmp[:readSize])
		if err != nil {
			return 0, err
		}
		d.uncompressedSize -= uint64(n)
		d.ofs += n
	}
	if err := d.nextFrame(); err != nil {
		return FrameFlags(0), err
	}

	return d.flags, nil
}

func (d *FrameDecoder) RemainingSize() uint64 {
	return d.uncompressedSize
}

func (d *FrameDecoder) Read(p []byte) (n int, err error) {
	if d.uncompressedSize == 0 {
		d.frameLoaded = false
		return 0, EndOfFrame
	}

	if d.uncompressedSize < uint64(len(p)) {
		p = p[:d.uncompressedSize]
	}

	n, err = d.frameContentSrc.Read(p)
	d.uncompressedSize -= uint64(n)
	d.ofs += n

	return n, err
}

func (d *FrameDecoder) ReadByte() (byte, error) {
	if d.uncompressedSize == 0 {
		d.frameLoaded = false
		return 0, EndOfFrame
	}
	d.uncompressedSize--
	d.ofs++
	return d.frameContentSrc.ReadByte()
}
