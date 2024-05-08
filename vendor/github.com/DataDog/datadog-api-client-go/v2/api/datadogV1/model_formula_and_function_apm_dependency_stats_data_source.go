// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionApmDependencyStatsDataSource Data source for APM dependency stats queries.
type FormulaAndFunctionApmDependencyStatsDataSource string

// List of FormulaAndFunctionApmDependencyStatsDataSource.
const (
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS FormulaAndFunctionApmDependencyStatsDataSource = "apm_dependency_stats"
)

var allowedFormulaAndFunctionApmDependencyStatsDataSourceEnumValues = []FormulaAndFunctionApmDependencyStatsDataSource{
	FORMULAANDFUNCTIONAPMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionApmDependencyStatsDataSource) GetAllowedValues() []FormulaAndFunctionApmDependencyStatsDataSource {
	return allowedFormulaAndFunctionApmDependencyStatsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionApmDependencyStatsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionApmDependencyStatsDataSource(value)
	return nil
}

// NewFormulaAndFunctionApmDependencyStatsDataSourceFromValue returns a pointer to a valid FormulaAndFunctionApmDependencyStatsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionApmDependencyStatsDataSourceFromValue(v string) (*FormulaAndFunctionApmDependencyStatsDataSource, error) {
	ev := FormulaAndFunctionApmDependencyStatsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionApmDependencyStatsDataSource: valid values are %v", v, allowedFormulaAndFunctionApmDependencyStatsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionApmDependencyStatsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionApmDependencyStatsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionApmDependencyStatsDataSource value.
func (v FormulaAndFunctionApmDependencyStatsDataSource) Ptr() *FormulaAndFunctionApmDependencyStatsDataSource {
	return &v
}
