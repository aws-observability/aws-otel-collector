// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

//go:build !goccy_gojson

package datadog

import (
	stdbytes "bytes"
	"encoding/json"
	"io"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// UnmarshalUseNumber decodes JSON with UseNumber enabled so that JSON numbers
// landing in interface{} targets (e.g. map[string]interface{} from
// additionalProperties) are kept as json.Number rather than coerced to float64.
// This preserves precision for integers above 2^53.
func UnmarshalUseNumber(data []byte, v interface{}) error {
	dec := json.NewDecoder(stdbytes.NewReader(data))
	dec.UseNumber()
	return dec.Decode(v)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
