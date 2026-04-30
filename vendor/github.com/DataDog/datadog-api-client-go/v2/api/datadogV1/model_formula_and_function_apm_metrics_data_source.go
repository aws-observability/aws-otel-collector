// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionApmMetricsDataSource Data source for APM metrics queries.
type FormulaAndFunctionApmMetricsDataSource string

// List of FormulaAndFunctionApmMetricsDataSource.
const (
	FORMULAANDFUNCTIONAPMMETRICSDATASOURCE_APM_METRICS FormulaAndFunctionApmMetricsDataSource = "apm_metrics"
)

var allowedFormulaAndFunctionApmMetricsDataSourceEnumValues = []FormulaAndFunctionApmMetricsDataSource{
	FORMULAANDFUNCTIONAPMMETRICSDATASOURCE_APM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionApmMetricsDataSource) GetAllowedValues() []FormulaAndFunctionApmMetricsDataSource {
	return allowedFormulaAndFunctionApmMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionApmMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionApmMetricsDataSource(value)
	return nil
}

// NewFormulaAndFunctionApmMetricsDataSourceFromValue returns a pointer to a valid FormulaAndFunctionApmMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionApmMetricsDataSourceFromValue(v string) (*FormulaAndFunctionApmMetricsDataSource, error) {
	ev := FormulaAndFunctionApmMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionApmMetricsDataSource: valid values are %v", v, allowedFormulaAndFunctionApmMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionApmMetricsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionApmMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionApmMetricsDataSource value.
func (v FormulaAndFunctionApmMetricsDataSource) Ptr() *FormulaAndFunctionApmMetricsDataSource {
	return &v
}
