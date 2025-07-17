// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutputSchemaParametersType The definition of `OutputSchemaParametersType` object.
type OutputSchemaParametersType string

// List of OutputSchemaParametersType.
const (
	OUTPUTSCHEMAPARAMETERSTYPE_STRING        OutputSchemaParametersType = "STRING"
	OUTPUTSCHEMAPARAMETERSTYPE_NUMBER        OutputSchemaParametersType = "NUMBER"
	OUTPUTSCHEMAPARAMETERSTYPE_BOOLEAN       OutputSchemaParametersType = "BOOLEAN"
	OUTPUTSCHEMAPARAMETERSTYPE_OBJECT        OutputSchemaParametersType = "OBJECT"
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_STRING  OutputSchemaParametersType = "ARRAY_STRING"
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_NUMBER  OutputSchemaParametersType = "ARRAY_NUMBER"
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_BOOLEAN OutputSchemaParametersType = "ARRAY_BOOLEAN"
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_OBJECT  OutputSchemaParametersType = "ARRAY_OBJECT"
)

var allowedOutputSchemaParametersTypeEnumValues = []OutputSchemaParametersType{
	OUTPUTSCHEMAPARAMETERSTYPE_STRING,
	OUTPUTSCHEMAPARAMETERSTYPE_NUMBER,
	OUTPUTSCHEMAPARAMETERSTYPE_BOOLEAN,
	OUTPUTSCHEMAPARAMETERSTYPE_OBJECT,
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_STRING,
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_NUMBER,
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_BOOLEAN,
	OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_OBJECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OutputSchemaParametersType) GetAllowedValues() []OutputSchemaParametersType {
	return allowedOutputSchemaParametersTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OutputSchemaParametersType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OutputSchemaParametersType(value)
	return nil
}

// NewOutputSchemaParametersTypeFromValue returns a pointer to a valid OutputSchemaParametersType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOutputSchemaParametersTypeFromValue(v string) (*OutputSchemaParametersType, error) {
	ev := OutputSchemaParametersType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OutputSchemaParametersType: valid values are %v", v, allowedOutputSchemaParametersTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OutputSchemaParametersType) IsValid() bool {
	for _, existing := range allowedOutputSchemaParametersTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutputSchemaParametersType value.
func (v OutputSchemaParametersType) Ptr() *OutputSchemaParametersType {
	return &v
}
