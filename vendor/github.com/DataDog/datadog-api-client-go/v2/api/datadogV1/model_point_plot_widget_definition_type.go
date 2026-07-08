// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotWidgetDefinitionType Type of the point plot widget.
type PointPlotWidgetDefinitionType string

// List of PointPlotWidgetDefinitionType.
const (
	POINTPLOTWIDGETDEFINITIONTYPE_POINT_PLOT PointPlotWidgetDefinitionType = "point_plot"
)

var allowedPointPlotWidgetDefinitionTypeEnumValues = []PointPlotWidgetDefinitionType{
	POINTPLOTWIDGETDEFINITIONTYPE_POINT_PLOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotWidgetDefinitionType) GetAllowedValues() []PointPlotWidgetDefinitionType {
	return allowedPointPlotWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotWidgetDefinitionType(value)
	return nil
}

// NewPointPlotWidgetDefinitionTypeFromValue returns a pointer to a valid PointPlotWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotWidgetDefinitionTypeFromValue(v string) (*PointPlotWidgetDefinitionType, error) {
	ev := PointPlotWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotWidgetDefinitionType: valid values are %v", v, allowedPointPlotWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedPointPlotWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotWidgetDefinitionType value.
func (v PointPlotWidgetDefinitionType) Ptr() *PointPlotWidgetDefinitionType {
	return &v
}
