// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormulaCellDisplayModeOptionsTrendType Trend type for the cell display mode options.
type WidgetFormulaCellDisplayModeOptionsTrendType string

// List of WidgetFormulaCellDisplayModeOptionsTrendType.
const (
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_AREA WidgetFormulaCellDisplayModeOptionsTrendType = "area"
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_LINE WidgetFormulaCellDisplayModeOptionsTrendType = "line"
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_BARS WidgetFormulaCellDisplayModeOptionsTrendType = "bars"
)

var allowedWidgetFormulaCellDisplayModeOptionsTrendTypeEnumValues = []WidgetFormulaCellDisplayModeOptionsTrendType{
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_AREA,
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_LINE,
	WIDGETFORMULACELLDISPLAYMODEOPTIONSTRENDTYPE_BARS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetFormulaCellDisplayModeOptionsTrendType) GetAllowedValues() []WidgetFormulaCellDisplayModeOptionsTrendType {
	return allowedWidgetFormulaCellDisplayModeOptionsTrendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetFormulaCellDisplayModeOptionsTrendType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetFormulaCellDisplayModeOptionsTrendType(value)
	return nil
}

// NewWidgetFormulaCellDisplayModeOptionsTrendTypeFromValue returns a pointer to a valid WidgetFormulaCellDisplayModeOptionsTrendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetFormulaCellDisplayModeOptionsTrendTypeFromValue(v string) (*WidgetFormulaCellDisplayModeOptionsTrendType, error) {
	ev := WidgetFormulaCellDisplayModeOptionsTrendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetFormulaCellDisplayModeOptionsTrendType: valid values are %v", v, allowedWidgetFormulaCellDisplayModeOptionsTrendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetFormulaCellDisplayModeOptionsTrendType) IsValid() bool {
	for _, existing := range allowedWidgetFormulaCellDisplayModeOptionsTrendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetFormulaCellDisplayModeOptionsTrendType value.
func (v WidgetFormulaCellDisplayModeOptionsTrendType) Ptr() *WidgetFormulaCellDisplayModeOptionsTrendType {
	return &v
}
