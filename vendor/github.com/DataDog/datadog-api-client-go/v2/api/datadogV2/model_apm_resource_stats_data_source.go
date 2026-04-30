// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmResourceStatsDataSource A data source for APM resource statistics queries.
type ApmResourceStatsDataSource string

// List of ApmResourceStatsDataSource.
const (
	APMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS ApmResourceStatsDataSource = "apm_resource_stats"
)

var allowedApmResourceStatsDataSourceEnumValues = []ApmResourceStatsDataSource{
	APMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmResourceStatsDataSource) GetAllowedValues() []ApmResourceStatsDataSource {
	return allowedApmResourceStatsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmResourceStatsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmResourceStatsDataSource(value)
	return nil
}

// NewApmResourceStatsDataSourceFromValue returns a pointer to a valid ApmResourceStatsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmResourceStatsDataSourceFromValue(v string) (*ApmResourceStatsDataSource, error) {
	ev := ApmResourceStatsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmResourceStatsDataSource: valid values are %v", v, allowedApmResourceStatsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmResourceStatsDataSource) IsValid() bool {
	for _, existing := range allowedApmResourceStatsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmResourceStatsDataSource value.
func (v ApmResourceStatsDataSource) Ptr() *ApmResourceStatsDataSource {
	return &v
}
