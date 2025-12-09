package internal

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func WriteString(str string, dst *bytes.Buffer) error {
	var b []byte
	b = binary.AppendUvarint(b, uint64(len(str)))
	b = append(b, str...)
	_, err := dst.Write(b)
	return err
}

// Maximum length of a string representing schema symbols such as struct name or user data key/value.
const maxStringLen = 256

var errStringTooLong = errors.New("string too long")

func ReadString(src *bytes.Buffer) (string, error) {
	l, err := binary.ReadUvarint(src)
	if err != nil {
		return "", err
	}
	if l > maxStringLen {
		return "", errStringTooLong
	}

	b := make([]byte, l)
	if _, err = src.Read(b); err != nil {
		return "", err
	}
	return string(b), nil
}

func WriteUvarint(v uint64, dst *bytes.Buffer) error {
	var b []byte
	b = binary.AppendUvarint(b, v)
	_, err := dst.Write(b)
	return err
}
