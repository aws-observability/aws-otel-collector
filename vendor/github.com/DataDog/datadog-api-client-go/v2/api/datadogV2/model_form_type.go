// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormType The resource type for a form.
type FormType string

// List of FormType.
const (
	FORMTYPE_FORMS FormType = "forms"
)

var allowedFormTypeEnumValues = []FormType{
	FORMTYPE_FORMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormType) GetAllowedValues() []FormType {
	return allowedFormTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormType(value)
	return nil
}

// NewFormTypeFromValue returns a pointer to a valid FormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormTypeFromValue(v string) (*FormType, error) {
	ev := FormType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormType: valid values are %v", v, allowedFormTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormType) IsValid() bool {
	for _, existing := range allowedFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormType value.
func (v FormType) Ptr() *FormType {
	return &v
}
