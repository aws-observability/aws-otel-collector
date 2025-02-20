// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreeTextWidgetDefinitionType Type of the free text widget.
type FreeTextWidgetDefinitionType string

// List of FreeTextWidgetDefinitionType.
const (
	FREETEXTWIDGETDEFINITIONTYPE_FREE_TEXT FreeTextWidgetDefinitionType = "free_text"
)

var allowedFreeTextWidgetDefinitionTypeEnumValues = []FreeTextWidgetDefinitionType{
	FREETEXTWIDGETDEFINITIONTYPE_FREE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FreeTextWidgetDefinitionType) GetAllowedValues() []FreeTextWidgetDefinitionType {
	return allowedFreeTextWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FreeTextWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FreeTextWidgetDefinitionType(value)
	return nil
}

// NewFreeTextWidgetDefinitionTypeFromValue returns a pointer to a valid FreeTextWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFreeTextWidgetDefinitionTypeFromValue(v string) (*FreeTextWidgetDefinitionType, error) {
	ev := FreeTextWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FreeTextWidgetDefinitionType: valid values are %v", v, allowedFreeTextWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FreeTextWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedFreeTextWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FreeTextWidgetDefinitionType value.
func (v FreeTextWidgetDefinitionType) Ptr() *FreeTextWidgetDefinitionType {
	return &v
}
