// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersionType The resource type for a form version.
type FormVersionType string

// List of FormVersionType.
const (
	FORMVERSIONTYPE_FORM_VERSIONS FormVersionType = "form_versions"
)

var allowedFormVersionTypeEnumValues = []FormVersionType{
	FORMVERSIONTYPE_FORM_VERSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormVersionType) GetAllowedValues() []FormVersionType {
	return allowedFormVersionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormVersionType(value)
	return nil
}

// NewFormVersionTypeFromValue returns a pointer to a valid FormVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormVersionTypeFromValue(v string) (*FormVersionType, error) {
	ev := FormVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormVersionType: valid values are %v", v, allowedFormVersionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormVersionType) IsValid() bool {
	for _, existing := range allowedFormVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormVersionType value.
func (v FormVersionType) Ptr() *FormVersionType {
	return &v
}
