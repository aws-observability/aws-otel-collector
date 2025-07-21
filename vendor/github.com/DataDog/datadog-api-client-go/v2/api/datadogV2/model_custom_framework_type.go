// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomFrameworkType The type of the resource. The value must be `custom_framework`.
type CustomFrameworkType string

// List of CustomFrameworkType.
const (
	CUSTOMFRAMEWORKTYPE_CUSTOM_FRAMEWORK CustomFrameworkType = "custom_framework"
)

var allowedCustomFrameworkTypeEnumValues = []CustomFrameworkType{
	CUSTOMFRAMEWORKTYPE_CUSTOM_FRAMEWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomFrameworkType) GetAllowedValues() []CustomFrameworkType {
	return allowedCustomFrameworkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomFrameworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomFrameworkType(value)
	return nil
}

// NewCustomFrameworkTypeFromValue returns a pointer to a valid CustomFrameworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomFrameworkTypeFromValue(v string) (*CustomFrameworkType, error) {
	ev := CustomFrameworkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomFrameworkType: valid values are %v", v, allowedCustomFrameworkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomFrameworkType) IsValid() bool {
	for _, existing := range allowedCustomFrameworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomFrameworkType value.
func (v CustomFrameworkType) Ptr() *CustomFrameworkType {
	return &v
}
