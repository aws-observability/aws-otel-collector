// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotProjectionType Type of the projection.
type PointPlotProjectionType string

// List of PointPlotProjectionType.
const (
	POINTPLOTPROJECTIONTYPE_POINT_PLOT PointPlotProjectionType = "point_plot"
)

var allowedPointPlotProjectionTypeEnumValues = []PointPlotProjectionType{
	POINTPLOTPROJECTIONTYPE_POINT_PLOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotProjectionType) GetAllowedValues() []PointPlotProjectionType {
	return allowedPointPlotProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotProjectionType(value)
	return nil
}

// NewPointPlotProjectionTypeFromValue returns a pointer to a valid PointPlotProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotProjectionTypeFromValue(v string) (*PointPlotProjectionType, error) {
	ev := PointPlotProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotProjectionType: valid values are %v", v, allowedPointPlotProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotProjectionType) IsValid() bool {
	for _, existing := range allowedPointPlotProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotProjectionType value.
func (v PointPlotProjectionType) Ptr() *PointPlotProjectionType {
	return &v
}
