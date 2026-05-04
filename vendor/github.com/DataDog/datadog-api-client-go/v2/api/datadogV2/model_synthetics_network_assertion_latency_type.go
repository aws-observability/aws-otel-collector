// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionLatencyType Type of the latency assertion.
type SyntheticsNetworkAssertionLatencyType string

// List of SyntheticsNetworkAssertionLatencyType.
const (
	SYNTHETICSNETWORKASSERTIONLATENCYTYPE_LATENCY SyntheticsNetworkAssertionLatencyType = "latency"
)

var allowedSyntheticsNetworkAssertionLatencyTypeEnumValues = []SyntheticsNetworkAssertionLatencyType{
	SYNTHETICSNETWORKASSERTIONLATENCYTYPE_LATENCY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkAssertionLatencyType) GetAllowedValues() []SyntheticsNetworkAssertionLatencyType {
	return allowedSyntheticsNetworkAssertionLatencyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkAssertionLatencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkAssertionLatencyType(value)
	return nil
}

// NewSyntheticsNetworkAssertionLatencyTypeFromValue returns a pointer to a valid SyntheticsNetworkAssertionLatencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkAssertionLatencyTypeFromValue(v string) (*SyntheticsNetworkAssertionLatencyType, error) {
	ev := SyntheticsNetworkAssertionLatencyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkAssertionLatencyType: valid values are %v", v, allowedSyntheticsNetworkAssertionLatencyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkAssertionLatencyType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkAssertionLatencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkAssertionLatencyType value.
func (v SyntheticsNetworkAssertionLatencyType) Ptr() *SyntheticsNetworkAssertionLatencyType {
	return &v
}
