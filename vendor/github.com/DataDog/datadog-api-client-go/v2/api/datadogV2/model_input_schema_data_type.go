// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InputSchemaDataType The definition of `InputSchemaDataType` object.
type InputSchemaDataType string

// List of InputSchemaDataType.
const (
	INPUTSCHEMADATATYPE_INPUTSCHEMA InputSchemaDataType = "inputSchema"
)

var allowedInputSchemaDataTypeEnumValues = []InputSchemaDataType{
	INPUTSCHEMADATATYPE_INPUTSCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InputSchemaDataType) GetAllowedValues() []InputSchemaDataType {
	return allowedInputSchemaDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InputSchemaDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InputSchemaDataType(value)
	return nil
}

// NewInputSchemaDataTypeFromValue returns a pointer to a valid InputSchemaDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInputSchemaDataTypeFromValue(v string) (*InputSchemaDataType, error) {
	ev := InputSchemaDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InputSchemaDataType: valid values are %v", v, allowedInputSchemaDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InputSchemaDataType) IsValid() bool {
	for _, existing := range allowedInputSchemaDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InputSchemaDataType value.
func (v InputSchemaDataType) Ptr() *InputSchemaDataType {
	return &v
}
