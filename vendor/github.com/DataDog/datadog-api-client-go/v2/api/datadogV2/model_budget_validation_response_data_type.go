// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetValidationResponseDataType Budget validation resource type.
type BudgetValidationResponseDataType string

// List of BudgetValidationResponseDataType.
const (
	BUDGETVALIDATIONRESPONSEDATATYPE_BUDGET_VALIDATION BudgetValidationResponseDataType = "budget_validation"
)

var allowedBudgetValidationResponseDataTypeEnumValues = []BudgetValidationResponseDataType{
	BUDGETVALIDATIONRESPONSEDATATYPE_BUDGET_VALIDATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BudgetValidationResponseDataType) GetAllowedValues() []BudgetValidationResponseDataType {
	return allowedBudgetValidationResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BudgetValidationResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BudgetValidationResponseDataType(value)
	return nil
}

// NewBudgetValidationResponseDataTypeFromValue returns a pointer to a valid BudgetValidationResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBudgetValidationResponseDataTypeFromValue(v string) (*BudgetValidationResponseDataType, error) {
	ev := BudgetValidationResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BudgetValidationResponseDataType: valid values are %v", v, allowedBudgetValidationResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BudgetValidationResponseDataType) IsValid() bool {
	for _, existing := range allowedBudgetValidationResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BudgetValidationResponseDataType value.
func (v BudgetValidationResponseDataType) Ptr() *BudgetValidationResponseDataType {
	return &v
}
