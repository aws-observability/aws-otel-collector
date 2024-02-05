// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesFormulaResponseType The type of the resource. The value should always be timeseries_response.
type TimeseriesFormulaResponseType string

// List of TimeseriesFormulaResponseType.
const (
	TIMESERIESFORMULARESPONSETYPE_TIMESERIES_RESPONSE TimeseriesFormulaResponseType = "timeseries_response"
)

var allowedTimeseriesFormulaResponseTypeEnumValues = []TimeseriesFormulaResponseType{
	TIMESERIESFORMULARESPONSETYPE_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesFormulaResponseType) GetAllowedValues() []TimeseriesFormulaResponseType {
	return allowedTimeseriesFormulaResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesFormulaResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesFormulaResponseType(value)
	return nil
}

// NewTimeseriesFormulaResponseTypeFromValue returns a pointer to a valid TimeseriesFormulaResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesFormulaResponseTypeFromValue(v string) (*TimeseriesFormulaResponseType, error) {
	ev := TimeseriesFormulaResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesFormulaResponseType: valid values are %v", v, allowedTimeseriesFormulaResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesFormulaResponseType) IsValid() bool {
	for _, existing := range allowedTimeseriesFormulaResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesFormulaResponseType value.
func (v TimeseriesFormulaResponseType) Ptr() *TimeseriesFormulaResponseType {
	return &v
}
