// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionProductAnalyticsExtendedDataSource Data source for Product Analytics Extended queries.
type FormulaAndFunctionProductAnalyticsExtendedDataSource string

// List of FormulaAndFunctionProductAnalyticsExtendedDataSource.
const (
	FORMULAANDFUNCTIONPRODUCTANALYTICSEXTENDEDDATASOURCE_PRODUCT_ANALYTICS_EXTENDED FormulaAndFunctionProductAnalyticsExtendedDataSource = "product_analytics_extended"
)

var allowedFormulaAndFunctionProductAnalyticsExtendedDataSourceEnumValues = []FormulaAndFunctionProductAnalyticsExtendedDataSource{
	FORMULAANDFUNCTIONPRODUCTANALYTICSEXTENDEDDATASOURCE_PRODUCT_ANALYTICS_EXTENDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionProductAnalyticsExtendedDataSource) GetAllowedValues() []FormulaAndFunctionProductAnalyticsExtendedDataSource {
	return allowedFormulaAndFunctionProductAnalyticsExtendedDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionProductAnalyticsExtendedDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionProductAnalyticsExtendedDataSource(value)
	return nil
}

// NewFormulaAndFunctionProductAnalyticsExtendedDataSourceFromValue returns a pointer to a valid FormulaAndFunctionProductAnalyticsExtendedDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionProductAnalyticsExtendedDataSourceFromValue(v string) (*FormulaAndFunctionProductAnalyticsExtendedDataSource, error) {
	ev := FormulaAndFunctionProductAnalyticsExtendedDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionProductAnalyticsExtendedDataSource: valid values are %v", v, allowedFormulaAndFunctionProductAnalyticsExtendedDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionProductAnalyticsExtendedDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionProductAnalyticsExtendedDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionProductAnalyticsExtendedDataSource value.
func (v FormulaAndFunctionProductAnalyticsExtendedDataSource) Ptr() *FormulaAndFunctionProductAnalyticsExtendedDataSource {
	return &v
}
