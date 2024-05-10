// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricSLOType SLO resource type.
type MetricSLOType string

// List of MetricSLOType.
const (
	METRICSLOTYPE_SLOS MetricSLOType = "slos"
)

var allowedMetricSLOTypeEnumValues = []MetricSLOType{
	METRICSLOTYPE_SLOS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricSLOType) GetAllowedValues() []MetricSLOType {
	return allowedMetricSLOTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricSLOType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricSLOType(value)
	return nil
}

// NewMetricSLOTypeFromValue returns a pointer to a valid MetricSLOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricSLOTypeFromValue(v string) (*MetricSLOType, error) {
	ev := MetricSLOType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricSLOType: valid values are %v", v, allowedMetricSLOTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricSLOType) IsValid() bool {
	for _, existing := range allowedMetricSLOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricSLOType value.
func (v MetricSLOType) Ptr() *MetricSLOType {
	return &v
}
