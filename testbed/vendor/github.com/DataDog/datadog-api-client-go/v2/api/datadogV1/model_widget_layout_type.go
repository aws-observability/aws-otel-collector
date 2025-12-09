// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetLayoutType Layout type of the group.
type WidgetLayoutType string

// List of WidgetLayoutType.
const (
	WIDGETLAYOUTTYPE_ORDERED WidgetLayoutType = "ordered"
)

var allowedWidgetLayoutTypeEnumValues = []WidgetLayoutType{
	WIDGETLAYOUTTYPE_ORDERED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetLayoutType) GetAllowedValues() []WidgetLayoutType {
	return allowedWidgetLayoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLayoutType(value)
	return nil
}

// NewWidgetLayoutTypeFromValue returns a pointer to a valid WidgetLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetLayoutTypeFromValue(v string) (*WidgetLayoutType, error) {
	ev := WidgetLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetLayoutType: valid values are %v", v, allowedWidgetLayoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetLayoutType) IsValid() bool {
	for _, existing := range allowedWidgetLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLayoutType value.
func (v WidgetLayoutType) Ptr() *WidgetLayoutType {
	return &v
}
