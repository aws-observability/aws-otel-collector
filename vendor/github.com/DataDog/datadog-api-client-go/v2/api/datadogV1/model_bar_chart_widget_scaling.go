// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetScaling Bar chart widget scaling definition.
type BarChartWidgetScaling string

// List of BarChartWidgetScaling.
const (
	BARCHARTWIDGETSCALING_ABSOLUTE BarChartWidgetScaling = "absolute"
	BARCHARTWIDGETSCALING_RELATIVE BarChartWidgetScaling = "relative"
)

var allowedBarChartWidgetScalingEnumValues = []BarChartWidgetScaling{
	BARCHARTWIDGETSCALING_ABSOLUTE,
	BARCHARTWIDGETSCALING_RELATIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartWidgetScaling) GetAllowedValues() []BarChartWidgetScaling {
	return allowedBarChartWidgetScalingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartWidgetScaling) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartWidgetScaling(value)
	return nil
}

// NewBarChartWidgetScalingFromValue returns a pointer to a valid BarChartWidgetScaling
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartWidgetScalingFromValue(v string) (*BarChartWidgetScaling, error) {
	ev := BarChartWidgetScaling(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartWidgetScaling: valid values are %v", v, allowedBarChartWidgetScalingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartWidgetScaling) IsValid() bool {
	for _, existing := range allowedBarChartWidgetScalingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartWidgetScaling value.
func (v BarChartWidgetScaling) Ptr() *BarChartWidgetScaling {
	return &v
}
