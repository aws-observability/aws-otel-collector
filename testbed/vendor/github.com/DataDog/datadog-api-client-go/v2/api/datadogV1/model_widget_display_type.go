// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetDisplayType Type of display to use for the request.
type WidgetDisplayType string

// List of WidgetDisplayType.
const (
	WIDGETDISPLAYTYPE_AREA    WidgetDisplayType = "area"
	WIDGETDISPLAYTYPE_BARS    WidgetDisplayType = "bars"
	WIDGETDISPLAYTYPE_LINE    WidgetDisplayType = "line"
	WIDGETDISPLAYTYPE_OVERLAY WidgetDisplayType = "overlay"
)

var allowedWidgetDisplayTypeEnumValues = []WidgetDisplayType{
	WIDGETDISPLAYTYPE_AREA,
	WIDGETDISPLAYTYPE_BARS,
	WIDGETDISPLAYTYPE_LINE,
	WIDGETDISPLAYTYPE_OVERLAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetDisplayType) GetAllowedValues() []WidgetDisplayType {
	return allowedWidgetDisplayTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetDisplayType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetDisplayType(value)
	return nil
}

// NewWidgetDisplayTypeFromValue returns a pointer to a valid WidgetDisplayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetDisplayTypeFromValue(v string) (*WidgetDisplayType, error) {
	ev := WidgetDisplayType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetDisplayType: valid values are %v", v, allowedWidgetDisplayTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetDisplayType) IsValid() bool {
	for _, existing := range allowedWidgetDisplayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetDisplayType value.
func (v WidgetDisplayType) Ptr() *WidgetDisplayType {
	return &v
}
