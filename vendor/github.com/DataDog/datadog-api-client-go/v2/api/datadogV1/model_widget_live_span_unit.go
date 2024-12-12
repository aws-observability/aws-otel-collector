// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetLiveSpanUnit Unit of the time span.
type WidgetLiveSpanUnit string

// List of WidgetLiveSpanUnit.
const (
	WIDGETLIVESPANUNIT_MINUTE WidgetLiveSpanUnit = "minute"
	WIDGETLIVESPANUNIT_HOUR   WidgetLiveSpanUnit = "hour"
	WIDGETLIVESPANUNIT_DAY    WidgetLiveSpanUnit = "day"
	WIDGETLIVESPANUNIT_WEEK   WidgetLiveSpanUnit = "week"
	WIDGETLIVESPANUNIT_MONTH  WidgetLiveSpanUnit = "month"
	WIDGETLIVESPANUNIT_YEAR   WidgetLiveSpanUnit = "year"
)

var allowedWidgetLiveSpanUnitEnumValues = []WidgetLiveSpanUnit{
	WIDGETLIVESPANUNIT_MINUTE,
	WIDGETLIVESPANUNIT_HOUR,
	WIDGETLIVESPANUNIT_DAY,
	WIDGETLIVESPANUNIT_WEEK,
	WIDGETLIVESPANUNIT_MONTH,
	WIDGETLIVESPANUNIT_YEAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetLiveSpanUnit) GetAllowedValues() []WidgetLiveSpanUnit {
	return allowedWidgetLiveSpanUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetLiveSpanUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetLiveSpanUnit(value)
	return nil
}

// NewWidgetLiveSpanUnitFromValue returns a pointer to a valid WidgetLiveSpanUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetLiveSpanUnitFromValue(v string) (*WidgetLiveSpanUnit, error) {
	ev := WidgetLiveSpanUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetLiveSpanUnit: valid values are %v", v, allowedWidgetLiveSpanUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetLiveSpanUnit) IsValid() bool {
	for _, existing := range allowedWidgetLiveSpanUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetLiveSpanUnit value.
func (v WidgetLiveSpanUnit) Ptr() *WidgetLiveSpanUnit {
	return &v
}
