// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionCloudCostDataSource Data source for Cloud Cost queries.
type FormulaAndFunctionCloudCostDataSource string

// List of FormulaAndFunctionCloudCostDataSource.
const (
	FORMULAANDFUNCTIONCLOUDCOSTDATASOURCE_CLOUD_COST FormulaAndFunctionCloudCostDataSource = "cloud_cost"
)

var allowedFormulaAndFunctionCloudCostDataSourceEnumValues = []FormulaAndFunctionCloudCostDataSource{
	FORMULAANDFUNCTIONCLOUDCOSTDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionCloudCostDataSource) GetAllowedValues() []FormulaAndFunctionCloudCostDataSource {
	return allowedFormulaAndFunctionCloudCostDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionCloudCostDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionCloudCostDataSource(value)
	return nil
}

// NewFormulaAndFunctionCloudCostDataSourceFromValue returns a pointer to a valid FormulaAndFunctionCloudCostDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionCloudCostDataSourceFromValue(v string) (*FormulaAndFunctionCloudCostDataSource, error) {
	ev := FormulaAndFunctionCloudCostDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionCloudCostDataSource: valid values are %v", v, allowedFormulaAndFunctionCloudCostDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionCloudCostDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionCloudCostDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionCloudCostDataSource value.
func (v FormulaAndFunctionCloudCostDataSource) Ptr() *FormulaAndFunctionCloudCostDataSource {
	return &v
}
