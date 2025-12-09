// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsConfigVariableType Type of the configuration variable.
type SyntheticsConfigVariableType string

// List of SyntheticsConfigVariableType.
const (
	SYNTHETICSCONFIGVARIABLETYPE_GLOBAL SyntheticsConfigVariableType = "global"
	SYNTHETICSCONFIGVARIABLETYPE_TEXT   SyntheticsConfigVariableType = "text"
	SYNTHETICSCONFIGVARIABLETYPE_EMAIL  SyntheticsConfigVariableType = "email"
)

var allowedSyntheticsConfigVariableTypeEnumValues = []SyntheticsConfigVariableType{
	SYNTHETICSCONFIGVARIABLETYPE_GLOBAL,
	SYNTHETICSCONFIGVARIABLETYPE_TEXT,
	SYNTHETICSCONFIGVARIABLETYPE_EMAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsConfigVariableType) GetAllowedValues() []SyntheticsConfigVariableType {
	return allowedSyntheticsConfigVariableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsConfigVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsConfigVariableType(value)
	return nil
}

// NewSyntheticsConfigVariableTypeFromValue returns a pointer to a valid SyntheticsConfigVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsConfigVariableTypeFromValue(v string) (*SyntheticsConfigVariableType, error) {
	ev := SyntheticsConfigVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsConfigVariableType: valid values are %v", v, allowedSyntheticsConfigVariableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsConfigVariableType) IsValid() bool {
	for _, existing := range allowedSyntheticsConfigVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsConfigVariableType value.
func (v SyntheticsConfigVariableType) Ptr() *SyntheticsConfigVariableType {
	return &v
}
