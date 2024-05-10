// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMonitorType Monitor resource type.
type MetricMonitorType string

// List of MetricMonitorType.
const (
	METRICMONITORTYPE_MONITORS MetricMonitorType = "monitors"
)

var allowedMetricMonitorTypeEnumValues = []MetricMonitorType{
	METRICMONITORTYPE_MONITORS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricMonitorType) GetAllowedValues() []MetricMonitorType {
	return allowedMetricMonitorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricMonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricMonitorType(value)
	return nil
}

// NewMetricMonitorTypeFromValue returns a pointer to a valid MetricMonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricMonitorTypeFromValue(v string) (*MetricMonitorType, error) {
	ev := MetricMonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricMonitorType: valid values are %v", v, allowedMetricMonitorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricMonitorType) IsValid() bool {
	for _, existing := range allowedMetricMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricMonitorType value.
func (v MetricMonitorType) Ptr() *MetricMonitorType {
	return &v
}
