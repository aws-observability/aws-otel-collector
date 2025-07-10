// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InputSchemaParametersType The definition of `InputSchemaParametersType` object.
type InputSchemaParametersType string

// List of InputSchemaParametersType.
const (
	INPUTSCHEMAPARAMETERSTYPE_STRING        InputSchemaParametersType = "STRING"
	INPUTSCHEMAPARAMETERSTYPE_NUMBER        InputSchemaParametersType = "NUMBER"
	INPUTSCHEMAPARAMETERSTYPE_BOOLEAN       InputSchemaParametersType = "BOOLEAN"
	INPUTSCHEMAPARAMETERSTYPE_OBJECT        InputSchemaParametersType = "OBJECT"
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_STRING  InputSchemaParametersType = "ARRAY_STRING"
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_NUMBER  InputSchemaParametersType = "ARRAY_NUMBER"
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_BOOLEAN InputSchemaParametersType = "ARRAY_BOOLEAN"
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_OBJECT  InputSchemaParametersType = "ARRAY_OBJECT"
)

var allowedInputSchemaParametersTypeEnumValues = []InputSchemaParametersType{
	INPUTSCHEMAPARAMETERSTYPE_STRING,
	INPUTSCHEMAPARAMETERSTYPE_NUMBER,
	INPUTSCHEMAPARAMETERSTYPE_BOOLEAN,
	INPUTSCHEMAPARAMETERSTYPE_OBJECT,
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_STRING,
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_NUMBER,
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_BOOLEAN,
	INPUTSCHEMAPARAMETERSTYPE_ARRAY_OBJECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *InputSchemaParametersType) GetAllowedValues() []InputSchemaParametersType {
	return allowedInputSchemaParametersTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InputSchemaParametersType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InputSchemaParametersType(value)
	return nil
}

// NewInputSchemaParametersTypeFromValue returns a pointer to a valid InputSchemaParametersType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInputSchemaParametersTypeFromValue(v string) (*InputSchemaParametersType, error) {
	ev := InputSchemaParametersType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InputSchemaParametersType: valid values are %v", v, allowedInputSchemaParametersTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InputSchemaParametersType) IsValid() bool {
	for _, existing := range allowedInputSchemaParametersTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InputSchemaParametersType value.
func (v InputSchemaParametersType) Ptr() *InputSchemaParametersType {
	return &v
}
