// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionJitterType Type of the jitter assertion.
type SyntheticsNetworkAssertionJitterType string

// List of SyntheticsNetworkAssertionJitterType.
const (
	SYNTHETICSNETWORKASSERTIONJITTERTYPE_JITTER SyntheticsNetworkAssertionJitterType = "jitter"
)

var allowedSyntheticsNetworkAssertionJitterTypeEnumValues = []SyntheticsNetworkAssertionJitterType{
	SYNTHETICSNETWORKASSERTIONJITTERTYPE_JITTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkAssertionJitterType) GetAllowedValues() []SyntheticsNetworkAssertionJitterType {
	return allowedSyntheticsNetworkAssertionJitterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkAssertionJitterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkAssertionJitterType(value)
	return nil
}

// NewSyntheticsNetworkAssertionJitterTypeFromValue returns a pointer to a valid SyntheticsNetworkAssertionJitterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkAssertionJitterTypeFromValue(v string) (*SyntheticsNetworkAssertionJitterType, error) {
	ev := SyntheticsNetworkAssertionJitterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkAssertionJitterType: valid values are %v", v, allowedSyntheticsNetworkAssertionJitterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkAssertionJitterType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkAssertionJitterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkAssertionJitterType value.
func (v SyntheticsNetworkAssertionJitterType) Ptr() *SyntheticsNetworkAssertionJitterType {
	return &v
}
