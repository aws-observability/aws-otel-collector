// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetLegend Bar chart widget stacked legend behavior.
type BarChartWidgetLegend string

// List of BarChartWidgetLegend.
const (
	BARCHARTWIDGETLEGEND_AUTOMATIC BarChartWidgetLegend = "automatic"
	BARCHARTWIDGETLEGEND_INLINE    BarChartWidgetLegend = "inline"
	BARCHARTWIDGETLEGEND_NONE      BarChartWidgetLegend = "none"
)

var allowedBarChartWidgetLegendEnumValues = []BarChartWidgetLegend{
	BARCHARTWIDGETLEGEND_AUTOMATIC,
	BARCHARTWIDGETLEGEND_INLINE,
	BARCHARTWIDGETLEGEND_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartWidgetLegend) GetAllowedValues() []BarChartWidgetLegend {
	return allowedBarChartWidgetLegendEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartWidgetLegend) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartWidgetLegend(value)
	return nil
}

// NewBarChartWidgetLegendFromValue returns a pointer to a valid BarChartWidgetLegend
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartWidgetLegendFromValue(v string) (*BarChartWidgetLegend, error) {
	ev := BarChartWidgetLegend(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartWidgetLegend: valid values are %v", v, allowedBarChartWidgetLegendEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartWidgetLegend) IsValid() bool {
	for _, existing := range allowedBarChartWidgetLegendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartWidgetLegend value.
func (v BarChartWidgetLegend) Ptr() *BarChartWidgetLegend {
	return &v
}
