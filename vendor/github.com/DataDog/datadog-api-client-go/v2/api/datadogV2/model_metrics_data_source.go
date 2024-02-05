// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricsDataSource A data source that is powered by the Metrics platform.
type MetricsDataSource string

// List of MetricsDataSource.
const (
	METRICSDATASOURCE_METRICS    MetricsDataSource = "metrics"
	METRICSDATASOURCE_CLOUD_COST MetricsDataSource = "cloud_cost"
)

var allowedMetricsDataSourceEnumValues = []MetricsDataSource{
	METRICSDATASOURCE_METRICS,
	METRICSDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricsDataSource) GetAllowedValues() []MetricsDataSource {
	return allowedMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricsDataSource(value)
	return nil
}

// NewMetricsDataSourceFromValue returns a pointer to a valid MetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricsDataSourceFromValue(v string) (*MetricsDataSource, error) {
	ev := MetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricsDataSource: valid values are %v", v, allowedMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricsDataSource) IsValid() bool {
	for _, existing := range allowedMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsDataSource value.
func (v MetricsDataSource) Ptr() *MetricsDataSource {
	return &v
}
