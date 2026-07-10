// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionMCPServerCapabilitiesType Type of the assertion.
type SyntheticsAssertionMCPServerCapabilitiesType string

// List of SyntheticsAssertionMCPServerCapabilitiesType.
const (
	SYNTHETICSASSERTIONMCPSERVERCAPABILITIESTYPE_MCP_SERVER_CAPABILITIES SyntheticsAssertionMCPServerCapabilitiesType = "mcpServerCapabilities"
)

var allowedSyntheticsAssertionMCPServerCapabilitiesTypeEnumValues = []SyntheticsAssertionMCPServerCapabilitiesType{
	SYNTHETICSASSERTIONMCPSERVERCAPABILITIESTYPE_MCP_SERVER_CAPABILITIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionMCPServerCapabilitiesType) GetAllowedValues() []SyntheticsAssertionMCPServerCapabilitiesType {
	return allowedSyntheticsAssertionMCPServerCapabilitiesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionMCPServerCapabilitiesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionMCPServerCapabilitiesType(value)
	return nil
}

// NewSyntheticsAssertionMCPServerCapabilitiesTypeFromValue returns a pointer to a valid SyntheticsAssertionMCPServerCapabilitiesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionMCPServerCapabilitiesTypeFromValue(v string) (*SyntheticsAssertionMCPServerCapabilitiesType, error) {
	ev := SyntheticsAssertionMCPServerCapabilitiesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionMCPServerCapabilitiesType: valid values are %v", v, allowedSyntheticsAssertionMCPServerCapabilitiesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionMCPServerCapabilitiesType) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionMCPServerCapabilitiesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionMCPServerCapabilitiesType value.
func (v SyntheticsAssertionMCPServerCapabilitiesType) Ptr() *SyntheticsAssertionMCPServerCapabilitiesType {
	return &v
}
