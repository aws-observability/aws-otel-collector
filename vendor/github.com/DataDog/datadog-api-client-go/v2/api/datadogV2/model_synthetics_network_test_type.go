// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestType Type of the Synthetic test, `network`.
type SyntheticsNetworkTestType string

// List of SyntheticsNetworkTestType.
const (
	SYNTHETICSNETWORKTESTTYPE_NETWORK SyntheticsNetworkTestType = "network"
)

var allowedSyntheticsNetworkTestTypeEnumValues = []SyntheticsNetworkTestType{
	SYNTHETICSNETWORKTESTTYPE_NETWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkTestType) GetAllowedValues() []SyntheticsNetworkTestType {
	return allowedSyntheticsNetworkTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkTestType(value)
	return nil
}

// NewSyntheticsNetworkTestTypeFromValue returns a pointer to a valid SyntheticsNetworkTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkTestTypeFromValue(v string) (*SyntheticsNetworkTestType, error) {
	ev := SyntheticsNetworkTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkTestType: valid values are %v", v, allowedSyntheticsNetworkTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkTestType value.
func (v SyntheticsNetworkTestType) Ptr() *SyntheticsNetworkTestType {
	return &v
}
