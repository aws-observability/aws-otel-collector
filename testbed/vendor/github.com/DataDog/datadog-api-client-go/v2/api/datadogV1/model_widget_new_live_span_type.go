// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetNewLiveSpanType Type "live" denotes a live span in the new format.
type WidgetNewLiveSpanType string

// List of WidgetNewLiveSpanType.
const (
	WIDGETNEWLIVESPANTYPE_LIVE WidgetNewLiveSpanType = "live"
)

var allowedWidgetNewLiveSpanTypeEnumValues = []WidgetNewLiveSpanType{
	WIDGETNEWLIVESPANTYPE_LIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetNewLiveSpanType) GetAllowedValues() []WidgetNewLiveSpanType {
	return allowedWidgetNewLiveSpanTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetNewLiveSpanType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetNewLiveSpanType(value)
	return nil
}

// NewWidgetNewLiveSpanTypeFromValue returns a pointer to a valid WidgetNewLiveSpanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetNewLiveSpanTypeFromValue(v string) (*WidgetNewLiveSpanType, error) {
	ev := WidgetNewLiveSpanType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetNewLiveSpanType: valid values are %v", v, allowedWidgetNewLiveSpanTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetNewLiveSpanType) IsValid() bool {
	for _, existing := range allowedWidgetNewLiveSpanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetNewLiveSpanType value.
func (v WidgetNewLiveSpanType) Ptr() *WidgetNewLiveSpanType {
	return &v
}
