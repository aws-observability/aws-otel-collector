//go:build cgo

package process

import (
	"bytes"
	"fmt"

	"github.com/DataDog/zstd"
	"github.com/DataDog/zstd_0"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	nocgozstd "github.com/klauspost/compress/zstd"
)

func unmarshal(enc MessageEncoding, body []byte, m proto.Message) error {
	switch enc {
	case MessageEncodingProtobuf:
		return proto.Unmarshal(body, m)
	case MessageEncodingJSON:
		return jsonpb.Unmarshal(bytes.NewReader(body), m)
	case MessageEncodingZstdPB, MessageEncodingZstd1xPB, MessageEncodingZstdPBxNoCgo:
		var d []byte
		var err error
		if enc == MessageEncodingZstd1xPB {
			d, err = zstd.Decompress(nil, body)
		} else if enc == MessageEncodingZstdPBxNoCgo {
			var decoder *nocgozstd.Decoder
			decoder, err = nocgozstd.NewReader(nil)
			if err != nil {
				return err
			}
			d, err = decoder.DecodeAll(body, nil)
		} else {
			d, err = zstd_0.Decompress(nil, body)
		}
		if err != nil {
			return err
		}
		return proto.Unmarshal(d, m)
	}
	return fmt.Errorf("unknown message encoding: %d", enc)
}

// EncodeMessage encodes a message object into bytes with protobuf. A type
// header is added for ease of decoding.
func EncodeMessage(m Message) ([]byte, error) {
	hb, err := encodeHeader(m.Header)
	if err != nil {
		return nil, fmt.Errorf("could not encode header: %s", err)
	}

	b := new(bytes.Buffer)
	if _, err := b.Write(hb); err != nil {
		return nil, err
	}

	var p []byte
	switch m.Header.Encoding {
	case MessageEncodingProtobuf:
		p, err = proto.Marshal(m.Body)
		if err != nil {
			return nil, err
		}
	case MessageEncodingJSON:
		marshaler := jsonpb.Marshaler{EmitDefaults: true}
		s, err := marshaler.MarshalToString(m.Body)
		if err != nil {
			return nil, err
		}
		p = []byte(s)
	case MessageEncodingZstdPB, MessageEncodingZstd1xPB, MessageEncodingZstdPBxNoCgo:
		pb, err := proto.Marshal(m.Body)
		if err != nil {
			return nil, err
		}

		if m.Header.Encoding == MessageEncodingZstd1xPB {
			p, err = zstd.Compress(nil, pb)
			if err != nil {
				return nil, err
			}
		} else if m.Header.Encoding == MessageEncodingZstdPBxNoCgo {
			var encoder *nocgozstd.Encoder
			encoder, err = nocgozstd.NewWriter(nil, nocgozstd.WithEncoderLevel(nocgozstd.SpeedDefault))
			if err != nil {
				return nil, err
			}
			p = encoder.EncodeAll(pb, nil)
			err = encoder.Close()
			if err != nil {
				return nil, err
			}
		} else {
			p, err = zstd_0.Compress(nil, pb)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unknown message encoding: %d", m.Header.Encoding)
	}
	_, err = b.Write(p)
	return b.Bytes(), err
}
