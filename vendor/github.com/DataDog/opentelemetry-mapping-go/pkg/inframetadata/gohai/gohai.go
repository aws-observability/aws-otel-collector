// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Original sources of this file:
//   - https://github.com/DataDog/datadog-agent/blob/dfab82/pkg/metadata/internal/gohai/payload.go
//   - https://github.com/DataDog/datadog-agent/blob/dfab82/pkg/metadata/v5/payload.go
//
// This file defines the 'Gohai' payload. gohai (https://github.com/DataDog/gohai) is a library for fetching system data.
// Its payload definition is used by the Datadog Agent and OpenTelemetry Collector Datadog exporter to export data about a given
// host, including information about its CPU, memory usage, network capabilities and platform.
//
// Most of the details on the payload definition are actually in DataDog/gohai, while here fields are just an empty interface.
// To make matters worse, this needs custom double marshaling because of legacy reasons.
//
// This payload definition is undergoing a refactor; the DataDog/gohai library is being deprecated in favor of a new module
// in DataDog/datadog-agent (see https://github.com/DataDog/gohai/pull/180).

package gohai

import (
	"encoding/json"
)

type Gohai struct {
	CPU        interface{} `json:"cpu"`
	FileSystem interface{} `json:"filesystem"`
	Memory     interface{} `json:"memory"`
	Network    interface{} `json:"network"`
	Platform   interface{} `json:"platform"`
}

// Payload handles the JSON unmarshalling of the metadata payload
// As weird as it sounds, in the v5 payload the value of the "gohai" field
// is a JSON-formatted string. So this struct contains a MarshaledGohaiPayload
// which will be marshaled as a JSON-formatted string.
type Payload struct {
	Gohai gohaiMarshaler `json:"gohai"`
}

// Platform returns a reference to the Gohai payload 'platform' map.
func (p *Payload) Platform() map[string]string {
	return p.Gohai.Gohai.Platform.(map[string]string)
}

// CPU returns a reference to the Gohai payload 'cpu' map.
func (p *Payload) CPU() map[string]string {
	return p.Gohai.Gohai.CPU.(map[string]string)
}

// gohaiSerializer implements json.Marshaler and json.Unmarshaler on top of a gohai payload
type gohaiMarshaler struct {
	Gohai *Gohai
}

// MarshalJSON implements the json.Marshaler interface.
// It marshals the gohai struct twice (to a string) to comply with
// the v5 payload format
func (m gohaiMarshaler) MarshalJSON() ([]byte, error) {
	marshaledPayload, err := json.Marshal(m.Gohai)
	if err != nil {
		return []byte(""), err
	}
	doubleMarshaledPayload, err := json.Marshal(string(marshaledPayload))
	if err != nil {
		return []byte(""), err
	}
	return doubleMarshaledPayload, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Unmarshals the passed bytes twice (first to a string, then to gohai.Gohai)
func (m *gohaiMarshaler) UnmarshalJSON(bytes []byte) error {
	// First, unmarshal to a string
	var out string
	err := json.Unmarshal(bytes, &out)
	if err != nil {
		return err
	}

	// Then, unmarshal the JSON-formatted string into a gohai.Gohai struct.
	return json.Unmarshal([]byte(out), &(m.Gohai))
}

// NewEmpty creates a new empty Gohai payload.
func NewEmpty() Payload {
	return Payload{
		Gohai: gohaiMarshaler{
			Gohai: &Gohai{
				Platform: map[string]string{},
				CPU:      map[string]string{},
			},
		},
	}
}
