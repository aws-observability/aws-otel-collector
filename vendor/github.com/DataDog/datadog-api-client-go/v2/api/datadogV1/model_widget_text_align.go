// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTextAlign How to align the text on the widget.
type WidgetTextAlign string

// List of WidgetTextAlign.
const (
	WIDGETTEXTALIGN_CENTER WidgetTextAlign = "center"
	WIDGETTEXTALIGN_LEFT   WidgetTextAlign = "left"
	WIDGETTEXTALIGN_RIGHT  WidgetTextAlign = "right"
)

var allowedWidgetTextAlignEnumValues = []WidgetTextAlign{
	WIDGETTEXTALIGN_CENTER,
	WIDGETTEXTALIGN_LEFT,
	WIDGETTEXTALIGN_RIGHT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetTextAlign) GetAllowedValues() []WidgetTextAlign {
	return allowedWidgetTextAlignEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetTextAlign) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetTextAlign(value)
	return nil
}

// NewWidgetTextAlignFromValue returns a pointer to a valid WidgetTextAlign
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetTextAlignFromValue(v string) (*WidgetTextAlign, error) {
	ev := WidgetTextAlign(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetTextAlign: valid values are %v", v, allowedWidgetTextAlignEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetTextAlign) IsValid() bool {
	for _, existing := range allowedWidgetTextAlignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetTextAlign value.
func (v WidgetTextAlign) Ptr() *WidgetTextAlign {
	return &v
}
