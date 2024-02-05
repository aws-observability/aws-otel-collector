// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesFormulaRequestType The type of the resource. The value should always be timeseries_request.
type TimeseriesFormulaRequestType string

// List of TimeseriesFormulaRequestType.
const (
	TIMESERIESFORMULAREQUESTTYPE_TIMESERIES_REQUEST TimeseriesFormulaRequestType = "timeseries_request"
)

var allowedTimeseriesFormulaRequestTypeEnumValues = []TimeseriesFormulaRequestType{
	TIMESERIESFORMULAREQUESTTYPE_TIMESERIES_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesFormulaRequestType) GetAllowedValues() []TimeseriesFormulaRequestType {
	return allowedTimeseriesFormulaRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesFormulaRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesFormulaRequestType(value)
	return nil
}

// NewTimeseriesFormulaRequestTypeFromValue returns a pointer to a valid TimeseriesFormulaRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesFormulaRequestTypeFromValue(v string) (*TimeseriesFormulaRequestType, error) {
	ev := TimeseriesFormulaRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesFormulaRequestType: valid values are %v", v, allowedTimeseriesFormulaRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesFormulaRequestType) IsValid() bool {
	for _, existing := range allowedTimeseriesFormulaRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesFormulaRequestType value.
func (v TimeseriesFormulaRequestType) Ptr() *TimeseriesFormulaRequestType {
	return &v
}
