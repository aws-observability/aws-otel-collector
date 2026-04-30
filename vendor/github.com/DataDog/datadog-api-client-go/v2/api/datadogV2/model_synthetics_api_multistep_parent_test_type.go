// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsApiMultistepParentTestType Type of the parent test resource.
type SyntheticsApiMultistepParentTestType string

// List of SyntheticsApiMultistepParentTestType.
const (
	SYNTHETICSAPIMULTISTEPPARENTTESTTYPE_PARENT_TEST SyntheticsApiMultistepParentTestType = "parent_test"
)

var allowedSyntheticsApiMultistepParentTestTypeEnumValues = []SyntheticsApiMultistepParentTestType{
	SYNTHETICSAPIMULTISTEPPARENTTESTTYPE_PARENT_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsApiMultistepParentTestType) GetAllowedValues() []SyntheticsApiMultistepParentTestType {
	return allowedSyntheticsApiMultistepParentTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsApiMultistepParentTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsApiMultistepParentTestType(value)
	return nil
}

// NewSyntheticsApiMultistepParentTestTypeFromValue returns a pointer to a valid SyntheticsApiMultistepParentTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsApiMultistepParentTestTypeFromValue(v string) (*SyntheticsApiMultistepParentTestType, error) {
	ev := SyntheticsApiMultistepParentTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsApiMultistepParentTestType: valid values are %v", v, allowedSyntheticsApiMultistepParentTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsApiMultistepParentTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsApiMultistepParentTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsApiMultistepParentTestType value.
func (v SyntheticsApiMultistepParentTestType) Ptr() *SyntheticsApiMultistepParentTestType {
	return &v
}
