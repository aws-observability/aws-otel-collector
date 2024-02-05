// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionMetricDataSource Data source for metrics queries.
type FormulaAndFunctionMetricDataSource string

// List of FormulaAndFunctionMetricDataSource.
const (
	FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS FormulaAndFunctionMetricDataSource = "metrics"
)

var allowedFormulaAndFunctionMetricDataSourceEnumValues = []FormulaAndFunctionMetricDataSource{
	FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionMetricDataSource) GetAllowedValues() []FormulaAndFunctionMetricDataSource {
	return allowedFormulaAndFunctionMetricDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionMetricDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionMetricDataSource(value)
	return nil
}

// NewFormulaAndFunctionMetricDataSourceFromValue returns a pointer to a valid FormulaAndFunctionMetricDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionMetricDataSourceFromValue(v string) (*FormulaAndFunctionMetricDataSource, error) {
	ev := FormulaAndFunctionMetricDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionMetricDataSource: valid values are %v", v, allowedFormulaAndFunctionMetricDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionMetricDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionMetricDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionMetricDataSource value.
func (v FormulaAndFunctionMetricDataSource) Ptr() *FormulaAndFunctionMetricDataSource {
	return &v
}
