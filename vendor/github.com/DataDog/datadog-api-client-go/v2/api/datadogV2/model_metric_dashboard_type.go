// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDashboardType Dashboard resource type.
type MetricDashboardType string

// List of MetricDashboardType.
const (
	METRICDASHBOARDTYPE_DASHBOARDS MetricDashboardType = "dashboards"
)

var allowedMetricDashboardTypeEnumValues = []MetricDashboardType{
	METRICDASHBOARDTYPE_DASHBOARDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricDashboardType) GetAllowedValues() []MetricDashboardType {
	return allowedMetricDashboardTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDashboardType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDashboardType(value)
	return nil
}

// NewMetricDashboardTypeFromValue returns a pointer to a valid MetricDashboardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDashboardTypeFromValue(v string) (*MetricDashboardType, error) {
	ev := MetricDashboardType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDashboardType: valid values are %v", v, allowedMetricDashboardTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDashboardType) IsValid() bool {
	for _, existing := range allowedMetricDashboardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDashboardType value.
func (v MetricDashboardType) Ptr() *MetricDashboardType {
	return &v
}
