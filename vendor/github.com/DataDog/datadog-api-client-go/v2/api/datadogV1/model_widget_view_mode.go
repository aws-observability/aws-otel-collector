// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetViewMode Define how you want the SLO to be displayed.
type WidgetViewMode string

// List of WidgetViewMode.
const (
	WIDGETVIEWMODE_OVERALL   WidgetViewMode = "overall"
	WIDGETVIEWMODE_COMPONENT WidgetViewMode = "component"
	WIDGETVIEWMODE_BOTH      WidgetViewMode = "both"
)

var allowedWidgetViewModeEnumValues = []WidgetViewMode{
	WIDGETVIEWMODE_OVERALL,
	WIDGETVIEWMODE_COMPONENT,
	WIDGETVIEWMODE_BOTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetViewMode) GetAllowedValues() []WidgetViewMode {
	return allowedWidgetViewModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetViewMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetViewMode(value)
	return nil
}

// NewWidgetViewModeFromValue returns a pointer to a valid WidgetViewMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetViewModeFromValue(v string) (*WidgetViewMode, error) {
	ev := WidgetViewMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetViewMode: valid values are %v", v, allowedWidgetViewModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetViewMode) IsValid() bool {
	for _, existing := range allowedWidgetViewModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetViewMode value.
func (v WidgetViewMode) Ptr() *WidgetViewMode {
	return &v
}
