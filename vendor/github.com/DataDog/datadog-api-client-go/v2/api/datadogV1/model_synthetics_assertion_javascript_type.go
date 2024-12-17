// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJavascriptType Type of the assertion.
type SyntheticsAssertionJavascriptType string

// List of SyntheticsAssertionJavascriptType.
const (
	SYNTHETICSASSERTIONJAVASCRIPTTYPE_JAVASCRIPT SyntheticsAssertionJavascriptType = "javascript"
)

var allowedSyntheticsAssertionJavascriptTypeEnumValues = []SyntheticsAssertionJavascriptType{
	SYNTHETICSASSERTIONJAVASCRIPTTYPE_JAVASCRIPT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionJavascriptType) GetAllowedValues() []SyntheticsAssertionJavascriptType {
	return allowedSyntheticsAssertionJavascriptTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionJavascriptType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionJavascriptType(value)
	return nil
}

// NewSyntheticsAssertionJavascriptTypeFromValue returns a pointer to a valid SyntheticsAssertionJavascriptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionJavascriptTypeFromValue(v string) (*SyntheticsAssertionJavascriptType, error) {
	ev := SyntheticsAssertionJavascriptType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionJavascriptType: valid values are %v", v, allowedSyntheticsAssertionJavascriptTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionJavascriptType) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionJavascriptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionJavascriptType value.
func (v SyntheticsAssertionJavascriptType) Ptr() *SyntheticsAssertionJavascriptType {
	return &v
}
