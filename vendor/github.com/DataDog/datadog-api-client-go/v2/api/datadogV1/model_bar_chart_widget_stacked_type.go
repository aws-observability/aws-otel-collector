// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetStackedType Bar chart widget stacked display type.
type BarChartWidgetStackedType string

// List of BarChartWidgetStackedType.
const (
	BARCHARTWIDGETSTACKEDTYPE_STACKED BarChartWidgetStackedType = "stacked"
)

var allowedBarChartWidgetStackedTypeEnumValues = []BarChartWidgetStackedType{
	BARCHARTWIDGETSTACKEDTYPE_STACKED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartWidgetStackedType) GetAllowedValues() []BarChartWidgetStackedType {
	return allowedBarChartWidgetStackedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartWidgetStackedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartWidgetStackedType(value)
	return nil
}

// NewBarChartWidgetStackedTypeFromValue returns a pointer to a valid BarChartWidgetStackedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartWidgetStackedTypeFromValue(v string) (*BarChartWidgetStackedType, error) {
	ev := BarChartWidgetStackedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartWidgetStackedType: valid values are %v", v, allowedBarChartWidgetStackedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartWidgetStackedType) IsValid() bool {
	for _, existing := range allowedBarChartWidgetStackedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartWidgetStackedType value.
func (v BarChartWidgetStackedType) Ptr() *BarChartWidgetStackedType {
	return &v
}
