// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApiMultistepSubtestType Type of the subtest resource.
type SyntheticsApiMultistepSubtestType string

// List of SyntheticsApiMultistepSubtestType.
const (
	SYNTHETICSAPIMULTISTEPSUBTESTTYPE_SUBTEST SyntheticsApiMultistepSubtestType = "subtest"
)

var allowedSyntheticsApiMultistepSubtestTypeEnumValues = []SyntheticsApiMultistepSubtestType{
	SYNTHETICSAPIMULTISTEPSUBTESTTYPE_SUBTEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsApiMultistepSubtestType) GetAllowedValues() []SyntheticsApiMultistepSubtestType {
	return allowedSyntheticsApiMultistepSubtestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsApiMultistepSubtestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsApiMultistepSubtestType(value)
	return nil
}

// NewSyntheticsApiMultistepSubtestTypeFromValue returns a pointer to a valid SyntheticsApiMultistepSubtestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsApiMultistepSubtestTypeFromValue(v string) (*SyntheticsApiMultistepSubtestType, error) {
	ev := SyntheticsApiMultistepSubtestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsApiMultistepSubtestType: valid values are %v", v, allowedSyntheticsApiMultistepSubtestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsApiMultistepSubtestType) IsValid() bool {
	for _, existing := range allowedSyntheticsApiMultistepSubtestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsApiMultistepSubtestType value.
func (v SyntheticsApiMultistepSubtestType) Ptr() *SyntheticsApiMultistepSubtestType {
	return &v
}
