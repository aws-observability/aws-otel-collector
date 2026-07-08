// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotWidgetLegendType Type of legend to show for the point plot widget.
type PointPlotWidgetLegendType string

// List of PointPlotWidgetLegendType.
const (
	POINTPLOTWIDGETLEGENDTYPE_AUTOMATIC PointPlotWidgetLegendType = "automatic"
	POINTPLOTWIDGETLEGENDTYPE_NONE      PointPlotWidgetLegendType = "none"
)

var allowedPointPlotWidgetLegendTypeEnumValues = []PointPlotWidgetLegendType{
	POINTPLOTWIDGETLEGENDTYPE_AUTOMATIC,
	POINTPLOTWIDGETLEGENDTYPE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotWidgetLegendType) GetAllowedValues() []PointPlotWidgetLegendType {
	return allowedPointPlotWidgetLegendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotWidgetLegendType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotWidgetLegendType(value)
	return nil
}

// NewPointPlotWidgetLegendTypeFromValue returns a pointer to a valid PointPlotWidgetLegendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotWidgetLegendTypeFromValue(v string) (*PointPlotWidgetLegendType, error) {
	ev := PointPlotWidgetLegendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotWidgetLegendType: valid values are %v", v, allowedPointPlotWidgetLegendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotWidgetLegendType) IsValid() bool {
	for _, existing := range allowedPointPlotWidgetLegendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotWidgetLegendType value.
func (v PointPlotWidgetLegendType) Ptr() *PointPlotWidgetLegendType {
	return &v
}
