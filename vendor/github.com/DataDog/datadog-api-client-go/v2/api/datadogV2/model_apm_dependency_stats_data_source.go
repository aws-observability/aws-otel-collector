// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmDependencyStatsDataSource A data source for APM dependency statistics queries.
type ApmDependencyStatsDataSource string

// List of ApmDependencyStatsDataSource.
const (
	APMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS ApmDependencyStatsDataSource = "apm_dependency_stats"
)

var allowedApmDependencyStatsDataSourceEnumValues = []ApmDependencyStatsDataSource{
	APMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmDependencyStatsDataSource) GetAllowedValues() []ApmDependencyStatsDataSource {
	return allowedApmDependencyStatsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmDependencyStatsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmDependencyStatsDataSource(value)
	return nil
}

// NewApmDependencyStatsDataSourceFromValue returns a pointer to a valid ApmDependencyStatsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmDependencyStatsDataSourceFromValue(v string) (*ApmDependencyStatsDataSource, error) {
	ev := ApmDependencyStatsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmDependencyStatsDataSource: valid values are %v", v, allowedApmDependencyStatsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmDependencyStatsDataSource) IsValid() bool {
	for _, existing := range allowedApmDependencyStatsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmDependencyStatsDataSource value.
func (v ApmDependencyStatsDataSource) Ptr() *ApmDependencyStatsDataSource {
	return &v
}
