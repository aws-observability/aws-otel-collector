// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetaPageType Type of metric pagination.
type MetricMetaPageType string

// List of MetricMetaPageType.
const (
	METRICMETAPAGETYPE_CURSOR_LIMIT MetricMetaPageType = "cursor_limit"
)

var allowedMetricMetaPageTypeEnumValues = []MetricMetaPageType{
	METRICMETAPAGETYPE_CURSOR_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricMetaPageType) GetAllowedValues() []MetricMetaPageType {
	return allowedMetricMetaPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricMetaPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricMetaPageType(value)
	return nil
}

// NewMetricMetaPageTypeFromValue returns a pointer to a valid MetricMetaPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricMetaPageTypeFromValue(v string) (*MetricMetaPageType, error) {
	ev := MetricMetaPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricMetaPageType: valid values are %v", v, allowedMetricMetaPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricMetaPageType) IsValid() bool {
	for _, existing := range allowedMetricMetaPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricMetaPageType value.
func (v MetricMetaPageType) Ptr() *MetricMetaPageType {
	return &v
}
