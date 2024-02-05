// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetDefinitionType Type of the timeseries widget.
type TimeseriesWidgetDefinitionType string

// List of TimeseriesWidgetDefinitionType.
const (
	TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES TimeseriesWidgetDefinitionType = "timeseries"
)

var allowedTimeseriesWidgetDefinitionTypeEnumValues = []TimeseriesWidgetDefinitionType{
	TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetDefinitionType) GetAllowedValues() []TimeseriesWidgetDefinitionType {
	return allowedTimeseriesWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetDefinitionType(value)
	return nil
}

// NewTimeseriesWidgetDefinitionTypeFromValue returns a pointer to a valid TimeseriesWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetDefinitionTypeFromValue(v string) (*TimeseriesWidgetDefinitionType, error) {
	ev := TimeseriesWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetDefinitionType: valid values are %v", v, allowedTimeseriesWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetDefinitionType value.
func (v TimeseriesWidgetDefinitionType) Ptr() *TimeseriesWidgetDefinitionType {
	return &v
}
