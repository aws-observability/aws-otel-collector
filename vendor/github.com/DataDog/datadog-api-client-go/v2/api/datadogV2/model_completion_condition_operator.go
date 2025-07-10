// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompletionConditionOperator The definition of `CompletionConditionOperator` object.
type CompletionConditionOperator string

// List of CompletionConditionOperator.
const (
	COMPLETIONCONDITIONOPERATOR_OPERATOR_EQUAL                    CompletionConditionOperator = "OPERATOR_EQUAL"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_NOT_EQUAL                CompletionConditionOperator = "OPERATOR_NOT_EQUAL"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_GREATER_THAN             CompletionConditionOperator = "OPERATOR_GREATER_THAN"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_LESS_THAN                CompletionConditionOperator = "OPERATOR_LESS_THAN"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_GREATER_THAN_OR_EQUAL_TO CompletionConditionOperator = "OPERATOR_GREATER_THAN_OR_EQUAL_TO"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_LESS_THAN_OR_EQUAL_TO    CompletionConditionOperator = "OPERATOR_LESS_THAN_OR_EQUAL_TO"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_CONTAINS                 CompletionConditionOperator = "OPERATOR_CONTAINS"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_DOES_NOT_CONTAIN         CompletionConditionOperator = "OPERATOR_DOES_NOT_CONTAIN"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NULL                  CompletionConditionOperator = "OPERATOR_IS_NULL"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NOT_NULL              CompletionConditionOperator = "OPERATOR_IS_NOT_NULL"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_EMPTY                 CompletionConditionOperator = "OPERATOR_IS_EMPTY"
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NOT_EMPTY             CompletionConditionOperator = "OPERATOR_IS_NOT_EMPTY"
)

var allowedCompletionConditionOperatorEnumValues = []CompletionConditionOperator{
	COMPLETIONCONDITIONOPERATOR_OPERATOR_EQUAL,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_NOT_EQUAL,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_GREATER_THAN,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_LESS_THAN,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_GREATER_THAN_OR_EQUAL_TO,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_LESS_THAN_OR_EQUAL_TO,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_CONTAINS,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_DOES_NOT_CONTAIN,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NULL,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NOT_NULL,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_EMPTY,
	COMPLETIONCONDITIONOPERATOR_OPERATOR_IS_NOT_EMPTY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompletionConditionOperator) GetAllowedValues() []CompletionConditionOperator {
	return allowedCompletionConditionOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompletionConditionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompletionConditionOperator(value)
	return nil
}

// NewCompletionConditionOperatorFromValue returns a pointer to a valid CompletionConditionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompletionConditionOperatorFromValue(v string) (*CompletionConditionOperator, error) {
	ev := CompletionConditionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompletionConditionOperator: valid values are %v", v, allowedCompletionConditionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompletionConditionOperator) IsValid() bool {
	for _, existing := range allowedCompletionConditionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompletionConditionOperator value.
func (v CompletionConditionOperator) Ptr() *CompletionConditionOperator {
	return &v
}
