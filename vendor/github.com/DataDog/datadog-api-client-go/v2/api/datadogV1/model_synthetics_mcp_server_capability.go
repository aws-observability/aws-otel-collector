// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMCPServerCapability A capability advertised by an MCP server.
type SyntheticsMCPServerCapability string

// List of SyntheticsMCPServerCapability.
const (
	SYNTHETICSMCPSERVERCAPABILITY_COMPLETIONS  SyntheticsMCPServerCapability = "completions"
	SYNTHETICSMCPSERVERCAPABILITY_EXPERIMENTAL SyntheticsMCPServerCapability = "experimental"
	SYNTHETICSMCPSERVERCAPABILITY_LOGGING      SyntheticsMCPServerCapability = "logging"
	SYNTHETICSMCPSERVERCAPABILITY_PROMPTS      SyntheticsMCPServerCapability = "prompts"
	SYNTHETICSMCPSERVERCAPABILITY_RESOURCES    SyntheticsMCPServerCapability = "resources"
	SYNTHETICSMCPSERVERCAPABILITY_TOOLS        SyntheticsMCPServerCapability = "tools"
)

var allowedSyntheticsMCPServerCapabilityEnumValues = []SyntheticsMCPServerCapability{
	SYNTHETICSMCPSERVERCAPABILITY_COMPLETIONS,
	SYNTHETICSMCPSERVERCAPABILITY_EXPERIMENTAL,
	SYNTHETICSMCPSERVERCAPABILITY_LOGGING,
	SYNTHETICSMCPSERVERCAPABILITY_PROMPTS,
	SYNTHETICSMCPSERVERCAPABILITY_RESOURCES,
	SYNTHETICSMCPSERVERCAPABILITY_TOOLS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMCPServerCapability) GetAllowedValues() []SyntheticsMCPServerCapability {
	return allowedSyntheticsMCPServerCapabilityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMCPServerCapability) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMCPServerCapability(value)
	return nil
}

// NewSyntheticsMCPServerCapabilityFromValue returns a pointer to a valid SyntheticsMCPServerCapability
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMCPServerCapabilityFromValue(v string) (*SyntheticsMCPServerCapability, error) {
	ev := SyntheticsMCPServerCapability(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMCPServerCapability: valid values are %v", v, allowedSyntheticsMCPServerCapabilityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMCPServerCapability) IsValid() bool {
	for _, existing := range allowedSyntheticsMCPServerCapabilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMCPServerCapability value.
func (v SyntheticsMCPServerCapability) Ptr() *SyntheticsMCPServerCapability {
	return &v
}
