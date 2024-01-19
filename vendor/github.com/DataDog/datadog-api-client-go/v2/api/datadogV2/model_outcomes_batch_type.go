// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesBatchType The JSON:API type for scorecard outcomes.
type OutcomesBatchType string

// List of OutcomesBatchType.
const (
	OUTCOMESBATCHTYPE_BATCHED_OUTCOME OutcomesBatchType = "batched-outcome"
)

var allowedOutcomesBatchTypeEnumValues = []OutcomesBatchType{
	OUTCOMESBATCHTYPE_BATCHED_OUTCOME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OutcomesBatchType) GetAllowedValues() []OutcomesBatchType {
	return allowedOutcomesBatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OutcomesBatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OutcomesBatchType(value)
	return nil
}

// NewOutcomesBatchTypeFromValue returns a pointer to a valid OutcomesBatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOutcomesBatchTypeFromValue(v string) (*OutcomesBatchType, error) {
	ev := OutcomesBatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OutcomesBatchType: valid values are %v", v, allowedOutcomesBatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OutcomesBatchType) IsValid() bool {
	for _, existing := range allowedOutcomesBatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutcomesBatchType value.
func (v OutcomesBatchType) Ptr() *OutcomesBatchType {
	return &v
}
