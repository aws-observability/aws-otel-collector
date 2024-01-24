// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomeType The JSON:API type for an outcome.
type OutcomeType string

// List of OutcomeType.
const (
	OUTCOMETYPE_OUTCOME OutcomeType = "outcome"
)

var allowedOutcomeTypeEnumValues = []OutcomeType{
	OUTCOMETYPE_OUTCOME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OutcomeType) GetAllowedValues() []OutcomeType {
	return allowedOutcomeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OutcomeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OutcomeType(value)
	return nil
}

// NewOutcomeTypeFromValue returns a pointer to a valid OutcomeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOutcomeTypeFromValue(v string) (*OutcomeType, error) {
	ev := OutcomeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OutcomeType: valid values are %v", v, allowedOutcomeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OutcomeType) IsValid() bool {
	for _, existing := range allowedOutcomeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutcomeType value.
func (v OutcomeType) Ptr() *OutcomeType {
	return &v
}
