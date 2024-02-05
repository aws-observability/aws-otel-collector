// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetLineType Type of lines displayed.
type WidgetLineType string

// List of WidgetLineType.
const (
	WIDGETLINETYPE_DASHED WidgetLineType = "dashed"
	WIDGETLINETYPE_DOTTED WidgetLineType = "dotted"
	WIDGETLINETYPE_SOLID  WidgetLineType = "solid"
)

var allowedWidgetLineTypeEnumValues = []WidgetLineType{
	WIDGETLINETYPE_DASHED,
	WIDGETLINETYPE_DOTTED,
	WIDGETLINETYPE_SOLID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetLineType) GetAllowedValues() []WidgetLineType {
	return allowedWidgetLineTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetLineType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLineType(value)
	return nil
}

// NewWidgetLineTypeFromValue returns a pointer to a valid WidgetLineType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetLineTypeFromValue(v string) (*WidgetLineType, error) {
	ev := WidgetLineType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetLineType: valid values are %v", v, allowedWidgetLineTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetLineType) IsValid() bool {
	for _, existing := range allowedWidgetLineTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLineType value.
func (v WidgetLineType) Ptr() *WidgetLineType {
	return &v
}
