// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConditionOperator The operator used in a targeting condition.
type ConditionOperator string

// List of ConditionOperator.
const (
	CONDITIONOPERATOR_LT          ConditionOperator = "LT"
	CONDITIONOPERATOR_LTE         ConditionOperator = "LTE"
	CONDITIONOPERATOR_GT          ConditionOperator = "GT"
	CONDITIONOPERATOR_GTE         ConditionOperator = "GTE"
	CONDITIONOPERATOR_MATCHES     ConditionOperator = "MATCHES"
	CONDITIONOPERATOR_NOT_MATCHES ConditionOperator = "NOT_MATCHES"
	CONDITIONOPERATOR_ONE_OF      ConditionOperator = "ONE_OF"
	CONDITIONOPERATOR_NOT_ONE_OF  ConditionOperator = "NOT_ONE_OF"
	CONDITIONOPERATOR_IS_NULL     ConditionOperator = "IS_NULL"
	CONDITIONOPERATOR_EQUALS      ConditionOperator = "EQUALS"
)

var allowedConditionOperatorEnumValues = []ConditionOperator{
	CONDITIONOPERATOR_LT,
	CONDITIONOPERATOR_LTE,
	CONDITIONOPERATOR_GT,
	CONDITIONOPERATOR_GTE,
	CONDITIONOPERATOR_MATCHES,
	CONDITIONOPERATOR_NOT_MATCHES,
	CONDITIONOPERATOR_ONE_OF,
	CONDITIONOPERATOR_NOT_ONE_OF,
	CONDITIONOPERATOR_IS_NULL,
	CONDITIONOPERATOR_EQUALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ConditionOperator) GetAllowedValues() []ConditionOperator {
	return allowedConditionOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConditionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConditionOperator(value)
	return nil
}

// NewConditionOperatorFromValue returns a pointer to a valid ConditionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConditionOperatorFromValue(v string) (*ConditionOperator, error) {
	ev := ConditionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConditionOperator: valid values are %v", v, allowedConditionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConditionOperator) IsValid() bool {
	for _, existing := range allowedConditionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConditionOperator value.
func (v ConditionOperator) Ptr() *ConditionOperator {
	return &v
}
