// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetVerticalAlign Vertical alignment.
type WidgetVerticalAlign string

// List of WidgetVerticalAlign.
const (
	WIDGETVERTICALALIGN_CENTER WidgetVerticalAlign = "center"
	WIDGETVERTICALALIGN_TOP    WidgetVerticalAlign = "top"
	WIDGETVERTICALALIGN_BOTTOM WidgetVerticalAlign = "bottom"
)

var allowedWidgetVerticalAlignEnumValues = []WidgetVerticalAlign{
	WIDGETVERTICALALIGN_CENTER,
	WIDGETVERTICALALIGN_TOP,
	WIDGETVERTICALALIGN_BOTTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetVerticalAlign) GetAllowedValues() []WidgetVerticalAlign {
	return allowedWidgetVerticalAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetVerticalAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetVerticalAlign(value)
	return nil
}

// NewWidgetVerticalAlignFromValue returns a pointer to a valid WidgetVerticalAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetVerticalAlignFromValue(v string) (*WidgetVerticalAlign, error) {
	ev := WidgetVerticalAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetVerticalAlign: valid values are %v", v, allowedWidgetVerticalAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetVerticalAlign) IsValid() bool {
	for _, existing := range allowedWidgetVerticalAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetVerticalAlign value.
func (v WidgetVerticalAlign) Ptr() *WidgetVerticalAlign {
	return &v
}
