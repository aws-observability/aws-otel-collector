// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageTimeSeriesType Type of usage data.
type UsageTimeSeriesType string

// List of UsageTimeSeriesType.
const (
	USAGETIMESERIESTYPE_USAGE_TIMESERIES UsageTimeSeriesType = "usage_timeseries"
)

var allowedUsageTimeSeriesTypeEnumValues = []UsageTimeSeriesType{
	USAGETIMESERIESTYPE_USAGE_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageTimeSeriesType) GetAllowedValues() []UsageTimeSeriesType {
	return allowedUsageTimeSeriesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageTimeSeriesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageTimeSeriesType(value)
	return nil
}

// NewUsageTimeSeriesTypeFromValue returns a pointer to a valid UsageTimeSeriesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageTimeSeriesTypeFromValue(v string) (*UsageTimeSeriesType, error) {
	ev := UsageTimeSeriesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageTimeSeriesType: valid values are %v", v, allowedUsageTimeSeriesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageTimeSeriesType) IsValid() bool {
	for _, existing := range allowedUsageTimeSeriesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageTimeSeriesType value.
func (v UsageTimeSeriesType) Ptr() *UsageTimeSeriesType {
	return &v
}
