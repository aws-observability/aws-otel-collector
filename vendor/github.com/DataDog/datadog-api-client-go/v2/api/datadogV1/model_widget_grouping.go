// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetGrouping The kind of grouping to use.
type WidgetGrouping string

// List of WidgetGrouping.
const (
	WIDGETGROUPING_CHECK   WidgetGrouping = "check"
	WIDGETGROUPING_CLUSTER WidgetGrouping = "cluster"
)

var allowedWidgetGroupingEnumValues = []WidgetGrouping{
	WIDGETGROUPING_CHECK,
	WIDGETGROUPING_CLUSTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetGrouping) GetAllowedValues() []WidgetGrouping {
	return allowedWidgetGroupingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetGrouping) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetGrouping(value)
	return nil
}

// NewWidgetGroupingFromValue returns a pointer to a valid WidgetGrouping
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetGroupingFromValue(v string) (*WidgetGrouping, error) {
	ev := WidgetGrouping(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetGrouping: valid values are %v", v, allowedWidgetGroupingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetGrouping) IsValid() bool {
	for _, existing := range allowedWidgetGroupingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetGrouping value.
func (v WidgetGrouping) Ptr() *WidgetGrouping {
	return &v
}
