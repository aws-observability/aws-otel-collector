package metadata

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"io"
)

var gzipMagic = []byte{0x1F, 0x8B, 0x08}

// GetUserData returns the user data for the current instance.
// NOTE: The result of this endpoint is automatically decoded from base64 and un-gzipped if needed.
func (c *Client) GetUserData(ctx context.Context) (string, error) {
	req := c.R(ctx)

	resp, err := coupleAPIErrors(req.Get("user-data"))
	if err != nil {
		return "", err
	}

	// user-data is returned as a raw string
	decodedBytes, err := base64.StdEncoding.DecodeString(resp.String())
	if err != nil {
		return "", fmt.Errorf("failed to decode user-data: %w", err)
	}
	rawBytes, err := ungzipIfNeeded(decodedBytes)
	if err != nil {
		return "", fmt.Errorf("failed to ungzip user-data: %w", err)
	}
	return string(rawBytes), nil
}

// hasGzipMagicNumber checks for the gzipMagic bytes at the beginning of the source
func hasGzipMagicNumber(source []byte) bool {
	return bytes.HasPrefix(source, gzipMagic)
}

// ungzipIfNeeded checks for the gzip magic number and unzips the bytes if necessary,
// otherwise it returns the original raw bytes
func ungzipIfNeeded(raw []byte) ([]byte, error) {
	if !hasGzipMagicNumber(raw) {
		return raw, nil
	}
	reader, err := gzip.NewReader(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}
