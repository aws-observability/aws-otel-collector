// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormulaCellDisplayModeOptionsYScale Y scale for the cell display mode options.
type WidgetFormulaCellDisplayModeOptionsYScale string

// List of WidgetFormulaCellDisplayModeOptionsYScale.
const (
	WIDGETFORMULACELLDISPLAYMODEOPTIONSYSCALE_SHARED      WidgetFormulaCellDisplayModeOptionsYScale = "shared"
	WIDGETFORMULACELLDISPLAYMODEOPTIONSYSCALE_INDEPENDENT WidgetFormulaCellDisplayModeOptionsYScale = "independent"
)

var allowedWidgetFormulaCellDisplayModeOptionsYScaleEnumValues = []WidgetFormulaCellDisplayModeOptionsYScale{
	WIDGETFORMULACELLDISPLAYMODEOPTIONSYSCALE_SHARED,
	WIDGETFORMULACELLDISPLAYMODEOPTIONSYSCALE_INDEPENDENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetFormulaCellDisplayModeOptionsYScale) GetAllowedValues() []WidgetFormulaCellDisplayModeOptionsYScale {
	return allowedWidgetFormulaCellDisplayModeOptionsYScaleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetFormulaCellDisplayModeOptionsYScale) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetFormulaCellDisplayModeOptionsYScale(value)
	return nil
}

// NewWidgetFormulaCellDisplayModeOptionsYScaleFromValue returns a pointer to a valid WidgetFormulaCellDisplayModeOptionsYScale
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetFormulaCellDisplayModeOptionsYScaleFromValue(v string) (*WidgetFormulaCellDisplayModeOptionsYScale, error) {
	ev := WidgetFormulaCellDisplayModeOptionsYScale(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetFormulaCellDisplayModeOptionsYScale: valid values are %v", v, allowedWidgetFormulaCellDisplayModeOptionsYScaleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetFormulaCellDisplayModeOptionsYScale) IsValid() bool {
	for _, existing := range allowedWidgetFormulaCellDisplayModeOptionsYScaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetFormulaCellDisplayModeOptionsYScale value.
func (v WidgetFormulaCellDisplayModeOptionsYScale) Ptr() *WidgetFormulaCellDisplayModeOptionsYScale {
	return &v
}
