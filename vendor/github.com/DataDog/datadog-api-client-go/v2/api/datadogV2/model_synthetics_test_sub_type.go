// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestSubType Subtype of the Synthetic test that produced this result.
type SyntheticsTestSubType string

// List of SyntheticsTestSubType.
const (
	SYNTHETICSTESTSUBTYPE_DNS       SyntheticsTestSubType = "dns"
	SYNTHETICSTESTSUBTYPE_GRPC      SyntheticsTestSubType = "grpc"
	SYNTHETICSTESTSUBTYPE_HTTP      SyntheticsTestSubType = "http"
	SYNTHETICSTESTSUBTYPE_ICMP      SyntheticsTestSubType = "icmp"
	SYNTHETICSTESTSUBTYPE_MCP       SyntheticsTestSubType = "mcp"
	SYNTHETICSTESTSUBTYPE_MULTI     SyntheticsTestSubType = "multi"
	SYNTHETICSTESTSUBTYPE_SSL       SyntheticsTestSubType = "ssl"
	SYNTHETICSTESTSUBTYPE_TCP       SyntheticsTestSubType = "tcp"
	SYNTHETICSTESTSUBTYPE_UDP       SyntheticsTestSubType = "udp"
	SYNTHETICSTESTSUBTYPE_WEBSOCKET SyntheticsTestSubType = "websocket"
)

var allowedSyntheticsTestSubTypeEnumValues = []SyntheticsTestSubType{
	SYNTHETICSTESTSUBTYPE_DNS,
	SYNTHETICSTESTSUBTYPE_GRPC,
	SYNTHETICSTESTSUBTYPE_HTTP,
	SYNTHETICSTESTSUBTYPE_ICMP,
	SYNTHETICSTESTSUBTYPE_MCP,
	SYNTHETICSTESTSUBTYPE_MULTI,
	SYNTHETICSTESTSUBTYPE_SSL,
	SYNTHETICSTESTSUBTYPE_TCP,
	SYNTHETICSTESTSUBTYPE_UDP,
	SYNTHETICSTESTSUBTYPE_WEBSOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestSubType) GetAllowedValues() []SyntheticsTestSubType {
	return allowedSyntheticsTestSubTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestSubType(value)
	return nil
}

// NewSyntheticsTestSubTypeFromValue returns a pointer to a valid SyntheticsTestSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestSubTypeFromValue(v string) (*SyntheticsTestSubType, error) {
	ev := SyntheticsTestSubType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestSubType: valid values are %v", v, allowedSyntheticsTestSubTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestSubType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestSubType value.
func (v SyntheticsTestSubType) Ptr() *SyntheticsTestSubType {
	return &v
}
