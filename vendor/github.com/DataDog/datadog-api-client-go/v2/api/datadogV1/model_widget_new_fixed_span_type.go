// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetNewFixedSpanType Type "fixed" denotes a fixed span.
type WidgetNewFixedSpanType string

// List of WidgetNewFixedSpanType.
const (
	WIDGETNEWFIXEDSPANTYPE_FIXED WidgetNewFixedSpanType = "fixed"
)

var allowedWidgetNewFixedSpanTypeEnumValues = []WidgetNewFixedSpanType{
	WIDGETNEWFIXEDSPANTYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetNewFixedSpanType) GetAllowedValues() []WidgetNewFixedSpanType {
	return allowedWidgetNewFixedSpanTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetNewFixedSpanType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetNewFixedSpanType(value)
	return nil
}

// NewWidgetNewFixedSpanTypeFromValue returns a pointer to a valid WidgetNewFixedSpanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetNewFixedSpanTypeFromValue(v string) (*WidgetNewFixedSpanType, error) {
	ev := WidgetNewFixedSpanType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetNewFixedSpanType: valid values are %v", v, allowedWidgetNewFixedSpanTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetNewFixedSpanType) IsValid() bool {
	for _, existing := range allowedWidgetNewFixedSpanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetNewFixedSpanType value.
func (v WidgetNewFixedSpanType) Ptr() *WidgetNewFixedSpanType {
	return &v
}
