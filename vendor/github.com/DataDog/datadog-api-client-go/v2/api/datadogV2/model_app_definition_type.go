// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AppDefinitionType The app definition type.
type AppDefinitionType string

// List of AppDefinitionType.
const (
	APPDEFINITIONTYPE_APPDEFINITIONS AppDefinitionType = "appDefinitions"
)

var allowedAppDefinitionTypeEnumValues = []AppDefinitionType{
	APPDEFINITIONTYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AppDefinitionType) GetAllowedValues() []AppDefinitionType {
	return allowedAppDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AppDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AppDefinitionType(value)
	return nil
}

// NewAppDefinitionTypeFromValue returns a pointer to a valid AppDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAppDefinitionTypeFromValue(v string) (*AppDefinitionType, error) {
	ev := AppDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AppDefinitionType: valid values are %v", v, allowedAppDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AppDefinitionType) IsValid() bool {
	for _, existing := range allowedAppDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppDefinitionType value.
func (v AppDefinitionType) Ptr() *AppDefinitionType {
	return &v
}
