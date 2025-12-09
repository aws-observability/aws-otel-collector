package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/splunk/stef/go/pkg/internal"
)

type FixedHeader struct {
	Compression Compression
}

type VarHeader struct {
	SchemaWireBytes []byte
	UserData        map[string]string
}

func (h *VarHeader) Serialize(dst *bytes.Buffer) error {
	if err := internal.WriteUvarint(uint64(len(h.SchemaWireBytes)), dst); err != nil {
		return err
	}

	if _, err := dst.Write(h.SchemaWireBytes); err != nil {
		return err
	}

	if err := internal.WriteUvarint(uint64(len(h.UserData)), dst); err != nil {
		return err
	}

	for k, v := range h.UserData {
		if err := internal.WriteString(k, dst); err != nil {
			return err
		}
		if err := internal.WriteString(v, dst); err != nil {
			return err
		}
	}

	return nil
}

const maxSchemaWireBytes = 1024 * 1024
const maxUserDataValues = 1024

func (h *VarHeader) Deserialize(src *bytes.Buffer) error {
	slen, err := binary.ReadUvarint(src)
	if err != nil {
		return err
	}
	if slen > maxSchemaWireBytes {
		return fmt.Errorf("schema too large: %d > %d", slen, maxSchemaWireBytes)
	}

	h.SchemaWireBytes = make([]byte, slen)
	if _, err = src.Read(h.SchemaWireBytes); err != nil {
		return err
	}

	count, err := binary.ReadUvarint(src)
	if err != nil {
		return err
	}
	if count > maxUserDataValues {
		return fmt.Errorf("too many user data values: %d > %d", count, maxUserDataValues)
	}

	h.UserData = make(map[string]string, count)
	for i := 0; i < int(count); i++ {
		k, err := internal.ReadString(src)
		if err != nil {
			return err
		}
		v, err := internal.ReadString(src)
		if err != nil {
			return err
		}
		h.UserData[k] = v
	}

	return nil
}
