// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetLegendLayout Layout of the legend.
type TimeseriesWidgetLegendLayout string

// List of TimeseriesWidgetLegendLayout.
const (
	TIMESERIESWIDGETLEGENDLAYOUT_AUTO       TimeseriesWidgetLegendLayout = "auto"
	TIMESERIESWIDGETLEGENDLAYOUT_HORIZONTAL TimeseriesWidgetLegendLayout = "horizontal"
	TIMESERIESWIDGETLEGENDLAYOUT_VERTICAL   TimeseriesWidgetLegendLayout = "vertical"
)

var allowedTimeseriesWidgetLegendLayoutEnumValues = []TimeseriesWidgetLegendLayout{
	TIMESERIESWIDGETLEGENDLAYOUT_AUTO,
	TIMESERIESWIDGETLEGENDLAYOUT_HORIZONTAL,
	TIMESERIESWIDGETLEGENDLAYOUT_VERTICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetLegendLayout) GetAllowedValues() []TimeseriesWidgetLegendLayout {
	return allowedTimeseriesWidgetLegendLayoutEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetLegendLayout) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetLegendLayout(value)
	return nil
}

// NewTimeseriesWidgetLegendLayoutFromValue returns a pointer to a valid TimeseriesWidgetLegendLayout
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetLegendLayoutFromValue(v string) (*TimeseriesWidgetLegendLayout, error) {
	ev := TimeseriesWidgetLegendLayout(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetLegendLayout: valid values are %v", v, allowedTimeseriesWidgetLegendLayoutEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetLegendLayout) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetLegendLayoutEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetLegendLayout value.
func (v TimeseriesWidgetLegendLayout) Ptr() *TimeseriesWidgetLegendLayout {
	return &v
}
