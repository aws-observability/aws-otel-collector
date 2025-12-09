// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackWidgetDefinitionType Type of the powerpack widget.
type PowerpackWidgetDefinitionType string

// List of PowerpackWidgetDefinitionType.
const (
	POWERPACKWIDGETDEFINITIONTYPE_POWERPACK PowerpackWidgetDefinitionType = "powerpack"
)

var allowedPowerpackWidgetDefinitionTypeEnumValues = []PowerpackWidgetDefinitionType{
	POWERPACKWIDGETDEFINITIONTYPE_POWERPACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PowerpackWidgetDefinitionType) GetAllowedValues() []PowerpackWidgetDefinitionType {
	return allowedPowerpackWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PowerpackWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PowerpackWidgetDefinitionType(value)
	return nil
}

// NewPowerpackWidgetDefinitionTypeFromValue returns a pointer to a valid PowerpackWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPowerpackWidgetDefinitionTypeFromValue(v string) (*PowerpackWidgetDefinitionType, error) {
	ev := PowerpackWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PowerpackWidgetDefinitionType: valid values are %v", v, allowedPowerpackWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PowerpackWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedPowerpackWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerpackWidgetDefinitionType value.
func (v PowerpackWidgetDefinitionType) Ptr() *PowerpackWidgetDefinitionType {
	return &v
}
