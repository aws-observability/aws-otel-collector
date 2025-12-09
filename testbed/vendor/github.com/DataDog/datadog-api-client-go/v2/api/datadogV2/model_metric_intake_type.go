// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricIntakeType The type of metric. The available types are `0` (unspecified), `1` (count), `2` (rate), and `3` (gauge).
type MetricIntakeType int32

// List of MetricIntakeType.
const (
	METRICINTAKETYPE_UNSPECIFIED MetricIntakeType = 0
	METRICINTAKETYPE_COUNT       MetricIntakeType = 1
	METRICINTAKETYPE_RATE        MetricIntakeType = 2
	METRICINTAKETYPE_GAUGE       MetricIntakeType = 3
)

var allowedMetricIntakeTypeEnumValues = []MetricIntakeType{
	METRICINTAKETYPE_UNSPECIFIED,
	METRICINTAKETYPE_COUNT,
	METRICINTAKETYPE_RATE,
	METRICINTAKETYPE_GAUGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricIntakeType) GetAllowedValues() []MetricIntakeType {
	return allowedMetricIntakeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricIntakeType) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricIntakeType(value)
	return nil
}

// NewMetricIntakeTypeFromValue returns a pointer to a valid MetricIntakeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricIntakeTypeFromValue(v int32) (*MetricIntakeType, error) {
	ev := MetricIntakeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricIntakeType: valid values are %v", v, allowedMetricIntakeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricIntakeType) IsValid() bool {
	for _, existing := range allowedMetricIntakeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricIntakeType value.
func (v MetricIntakeType) Ptr() *MetricIntakeType {
	return &v
}
