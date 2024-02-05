// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetCompareTo Timeframe used for the change comparison.
type WidgetCompareTo string

// List of WidgetCompareTo.
const (
	WIDGETCOMPARETO_HOUR_BEFORE  WidgetCompareTo = "hour_before"
	WIDGETCOMPARETO_DAY_BEFORE   WidgetCompareTo = "day_before"
	WIDGETCOMPARETO_WEEK_BEFORE  WidgetCompareTo = "week_before"
	WIDGETCOMPARETO_MONTH_BEFORE WidgetCompareTo = "month_before"
)

var allowedWidgetCompareToEnumValues = []WidgetCompareTo{
	WIDGETCOMPARETO_HOUR_BEFORE,
	WIDGETCOMPARETO_DAY_BEFORE,
	WIDGETCOMPARETO_WEEK_BEFORE,
	WIDGETCOMPARETO_MONTH_BEFORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetCompareTo) GetAllowedValues() []WidgetCompareTo {
	return allowedWidgetCompareToEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetCompareTo) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetCompareTo(value)
	return nil
}

// NewWidgetCompareToFromValue returns a pointer to a valid WidgetCompareTo
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetCompareToFromValue(v string) (*WidgetCompareTo, error) {
	ev := WidgetCompareTo(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetCompareTo: valid values are %v", v, allowedWidgetCompareToEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetCompareTo) IsValid() bool {
	for _, existing := range allowedWidgetCompareToEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetCompareTo value.
func (v WidgetCompareTo) Ptr() *WidgetCompareTo {
	return &v
}
