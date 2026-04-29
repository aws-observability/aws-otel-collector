// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsDataSource A data source for APM metrics queries.
type ApmMetricsDataSource string

// List of ApmMetricsDataSource.
const (
	APMMETRICSDATASOURCE_APM_METRICS ApmMetricsDataSource = "apm_metrics"
)

var allowedApmMetricsDataSourceEnumValues = []ApmMetricsDataSource{
	APMMETRICSDATASOURCE_APM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricsDataSource) GetAllowedValues() []ApmMetricsDataSource {
	return allowedApmMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricsDataSource(value)
	return nil
}

// NewApmMetricsDataSourceFromValue returns a pointer to a valid ApmMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricsDataSourceFromValue(v string) (*ApmMetricsDataSource, error) {
	ev := ApmMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricsDataSource: valid values are %v", v, allowedApmMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricsDataSource) IsValid() bool {
	for _, existing := range allowedApmMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricsDataSource value.
func (v ApmMetricsDataSource) Ptr() *ApmMetricsDataSource {
	return &v
}
