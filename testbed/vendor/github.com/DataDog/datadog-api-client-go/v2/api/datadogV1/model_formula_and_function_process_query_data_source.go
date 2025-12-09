// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionProcessQueryDataSource Data sources that rely on the process backend.
type FormulaAndFunctionProcessQueryDataSource string

// List of FormulaAndFunctionProcessQueryDataSource.
const (
	FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_PROCESS   FormulaAndFunctionProcessQueryDataSource = "process"
	FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_CONTAINER FormulaAndFunctionProcessQueryDataSource = "container"
)

var allowedFormulaAndFunctionProcessQueryDataSourceEnumValues = []FormulaAndFunctionProcessQueryDataSource{
	FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_PROCESS,
	FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_CONTAINER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionProcessQueryDataSource) GetAllowedValues() []FormulaAndFunctionProcessQueryDataSource {
	return allowedFormulaAndFunctionProcessQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionProcessQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionProcessQueryDataSource(value)
	return nil
}

// NewFormulaAndFunctionProcessQueryDataSourceFromValue returns a pointer to a valid FormulaAndFunctionProcessQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionProcessQueryDataSourceFromValue(v string) (*FormulaAndFunctionProcessQueryDataSource, error) {
	ev := FormulaAndFunctionProcessQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionProcessQueryDataSource: valid values are %v", v, allowedFormulaAndFunctionProcessQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionProcessQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionProcessQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionProcessQueryDataSource value.
func (v FormulaAndFunctionProcessQueryDataSource) Ptr() *FormulaAndFunctionProcessQueryDataSource {
	return &v
}
