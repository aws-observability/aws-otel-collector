// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormDataDefinitionType The root schema type.
type FormDataDefinitionType string

// List of FormDataDefinitionType.
const (
	FORMDATADEFINITIONTYPE_OBJECT FormDataDefinitionType = "object"
)

var allowedFormDataDefinitionTypeEnumValues = []FormDataDefinitionType{
	FORMDATADEFINITIONTYPE_OBJECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormDataDefinitionType) GetAllowedValues() []FormDataDefinitionType {
	return allowedFormDataDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormDataDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormDataDefinitionType(value)
	return nil
}

// NewFormDataDefinitionTypeFromValue returns a pointer to a valid FormDataDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormDataDefinitionTypeFromValue(v string) (*FormDataDefinitionType, error) {
	ev := FormDataDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormDataDefinitionType: valid values are %v", v, allowedFormDataDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormDataDefinitionType) IsValid() bool {
	for _, existing := range allowedFormDataDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormDataDefinitionType value.
func (v FormDataDefinitionType) Ptr() *FormDataDefinitionType {
	return &v
}
