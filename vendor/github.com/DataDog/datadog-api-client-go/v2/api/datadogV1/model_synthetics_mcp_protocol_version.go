// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMCPProtocolVersion The MCP protocol version used by the step. See https://modelcontextprotocol.io/specification.
type SyntheticsMCPProtocolVersion string

// List of SyntheticsMCPProtocolVersion.
const (
	SYNTHETICSMCPPROTOCOLVERSION_VERSION_2025_06_18 SyntheticsMCPProtocolVersion = "2025-06-18"
)

var allowedSyntheticsMCPProtocolVersionEnumValues = []SyntheticsMCPProtocolVersion{
	SYNTHETICSMCPPROTOCOLVERSION_VERSION_2025_06_18,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMCPProtocolVersion) GetAllowedValues() []SyntheticsMCPProtocolVersion {
	return allowedSyntheticsMCPProtocolVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMCPProtocolVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMCPProtocolVersion(value)
	return nil
}

// NewSyntheticsMCPProtocolVersionFromValue returns a pointer to a valid SyntheticsMCPProtocolVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMCPProtocolVersionFromValue(v string) (*SyntheticsMCPProtocolVersion, error) {
	ev := SyntheticsMCPProtocolVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMCPProtocolVersion: valid values are %v", v, allowedSyntheticsMCPProtocolVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMCPProtocolVersion) IsValid() bool {
	for _, existing := range allowedSyntheticsMCPProtocolVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMCPProtocolVersion value.
func (v SyntheticsMCPProtocolVersion) Ptr() *SyntheticsMCPProtocolVersion {
	return &v
}
