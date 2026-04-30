// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestSubType Subtype of the Synthetic test that produced this result.
type SyntheticsFastTestSubType string

// List of SyntheticsFastTestSubType.
const (
	SYNTHETICSFASTTESTSUBTYPE_DNS       SyntheticsFastTestSubType = "dns"
	SYNTHETICSFASTTESTSUBTYPE_GRPC      SyntheticsFastTestSubType = "grpc"
	SYNTHETICSFASTTESTSUBTYPE_HTTP      SyntheticsFastTestSubType = "http"
	SYNTHETICSFASTTESTSUBTYPE_ICMP      SyntheticsFastTestSubType = "icmp"
	SYNTHETICSFASTTESTSUBTYPE_MCP       SyntheticsFastTestSubType = "mcp"
	SYNTHETICSFASTTESTSUBTYPE_MULTI     SyntheticsFastTestSubType = "multi"
	SYNTHETICSFASTTESTSUBTYPE_SSL       SyntheticsFastTestSubType = "ssl"
	SYNTHETICSFASTTESTSUBTYPE_TCP       SyntheticsFastTestSubType = "tcp"
	SYNTHETICSFASTTESTSUBTYPE_UDP       SyntheticsFastTestSubType = "udp"
	SYNTHETICSFASTTESTSUBTYPE_WEBSOCKET SyntheticsFastTestSubType = "websocket"
)

var allowedSyntheticsFastTestSubTypeEnumValues = []SyntheticsFastTestSubType{
	SYNTHETICSFASTTESTSUBTYPE_DNS,
	SYNTHETICSFASTTESTSUBTYPE_GRPC,
	SYNTHETICSFASTTESTSUBTYPE_HTTP,
	SYNTHETICSFASTTESTSUBTYPE_ICMP,
	SYNTHETICSFASTTESTSUBTYPE_MCP,
	SYNTHETICSFASTTESTSUBTYPE_MULTI,
	SYNTHETICSFASTTESTSUBTYPE_SSL,
	SYNTHETICSFASTTESTSUBTYPE_TCP,
	SYNTHETICSFASTTESTSUBTYPE_UDP,
	SYNTHETICSFASTTESTSUBTYPE_WEBSOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsFastTestSubType) GetAllowedValues() []SyntheticsFastTestSubType {
	return allowedSyntheticsFastTestSubTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsFastTestSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsFastTestSubType(value)
	return nil
}

// NewSyntheticsFastTestSubTypeFromValue returns a pointer to a valid SyntheticsFastTestSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsFastTestSubTypeFromValue(v string) (*SyntheticsFastTestSubType, error) {
	ev := SyntheticsFastTestSubType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsFastTestSubType: valid values are %v", v, allowedSyntheticsFastTestSubTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsFastTestSubType) IsValid() bool {
	for _, existing := range allowedSyntheticsFastTestSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsFastTestSubType value.
func (v SyntheticsFastTestSubType) Ptr() *SyntheticsFastTestSubType {
	return &v
}
