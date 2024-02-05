// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetColorPreference Which color to use on the widget.
type WidgetColorPreference string

// List of WidgetColorPreference.
const (
	WIDGETCOLORPREFERENCE_BACKGROUND WidgetColorPreference = "background"
	WIDGETCOLORPREFERENCE_TEXT       WidgetColorPreference = "text"
)

var allowedWidgetColorPreferenceEnumValues = []WidgetColorPreference{
	WIDGETCOLORPREFERENCE_BACKGROUND,
	WIDGETCOLORPREFERENCE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetColorPreference) GetAllowedValues() []WidgetColorPreference {
	return allowedWidgetColorPreferenceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetColorPreference) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetColorPreference(value)
	return nil
}

// NewWidgetColorPreferenceFromValue returns a pointer to a valid WidgetColorPreference
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetColorPreferenceFromValue(v string) (*WidgetColorPreference, error) {
	ev := WidgetColorPreference(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetColorPreference: valid values are %v", v, allowedWidgetColorPreferenceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetColorPreference) IsValid() bool {
	for _, existing := range allowedWidgetColorPreferenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetColorPreference value.
func (v WidgetColorPreference) Ptr() *WidgetColorPreference {
	return &v
}
