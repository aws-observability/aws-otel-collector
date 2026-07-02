// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotDimension Dimension of the point plot.
type PointPlotDimension string

// List of PointPlotDimension.
const (
	POINTPLOTDIMENSION_GROUP  PointPlotDimension = "group"
	POINTPLOTDIMENSION_TIME   PointPlotDimension = "time"
	POINTPLOTDIMENSION_Y      PointPlotDimension = "y"
	POINTPLOTDIMENSION_RADIUS PointPlotDimension = "radius"
)

var allowedPointPlotDimensionEnumValues = []PointPlotDimension{
	POINTPLOTDIMENSION_GROUP,
	POINTPLOTDIMENSION_TIME,
	POINTPLOTDIMENSION_Y,
	POINTPLOTDIMENSION_RADIUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotDimension) GetAllowedValues() []PointPlotDimension {
	return allowedPointPlotDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotDimension(value)
	return nil
}

// NewPointPlotDimensionFromValue returns a pointer to a valid PointPlotDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotDimensionFromValue(v string) (*PointPlotDimension, error) {
	ev := PointPlotDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotDimension: valid values are %v", v, allowedPointPlotDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotDimension) IsValid() bool {
	for _, existing := range allowedPointPlotDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotDimension value.
func (v PointPlotDimension) Ptr() *PointPlotDimension {
	return &v
}
