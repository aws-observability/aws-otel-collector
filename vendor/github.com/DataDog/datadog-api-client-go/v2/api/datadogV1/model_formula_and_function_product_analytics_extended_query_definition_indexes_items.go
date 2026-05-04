// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems Use `"*"` to query all indexes.
type FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems string

// List of FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems.
const (
	FORMULAANDFUNCTIONPRODUCTANALYTICSEXTENDEDQUERYDEFINITIONINDEXESITEMS_ALL FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems = "*"
)

var allowedFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsEnumValues = []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems{
	FORMULAANDFUNCTIONPRODUCTANALYTICSEXTENDEDQUERYDEFINITIONINDEXESITEMS_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems) GetAllowedValues() []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems {
	return allowedFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems(value)
	return nil
}

// NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsFromValue returns a pointer to a valid FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsFromValue(v string) (*FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems, error) {
	ev := FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems: valid values are %v", v, allowedFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItemsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems value.
func (v FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems) Ptr() *FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems {
	return &v
}
