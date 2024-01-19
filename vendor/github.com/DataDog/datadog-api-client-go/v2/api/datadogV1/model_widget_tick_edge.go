// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTickEdge Define how you want to align the text on the widget.
type WidgetTickEdge string

// List of WidgetTickEdge.
const (
	WIDGETTICKEDGE_BOTTOM WidgetTickEdge = "bottom"
	WIDGETTICKEDGE_LEFT   WidgetTickEdge = "left"
	WIDGETTICKEDGE_RIGHT  WidgetTickEdge = "right"
	WIDGETTICKEDGE_TOP    WidgetTickEdge = "top"
)

var allowedWidgetTickEdgeEnumValues = []WidgetTickEdge{
	WIDGETTICKEDGE_BOTTOM,
	WIDGETTICKEDGE_LEFT,
	WIDGETTICKEDGE_RIGHT,
	WIDGETTICKEDGE_TOP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetTickEdge) GetAllowedValues() []WidgetTickEdge {
	return allowedWidgetTickEdgeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetTickEdge) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetTickEdge(value)
	return nil
}

// NewWidgetTickEdgeFromValue returns a pointer to a valid WidgetTickEdge
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetTickEdgeFromValue(v string) (*WidgetTickEdge, error) {
	ev := WidgetTickEdge(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetTickEdge: valid values are %v", v, allowedWidgetTickEdgeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetTickEdge) IsValid() bool {
	for _, existing := range allowedWidgetTickEdgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetTickEdge value.
func (v WidgetTickEdge) Ptr() *WidgetTickEdge {
	return &v
}
