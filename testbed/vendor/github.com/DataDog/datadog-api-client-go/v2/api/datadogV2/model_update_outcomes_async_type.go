// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateOutcomesAsyncType The JSON:API type for scorecard outcomes.
type UpdateOutcomesAsyncType string

// List of UpdateOutcomesAsyncType.
const (
	UPDATEOUTCOMESASYNCTYPE_BATCHED_OUTCOME UpdateOutcomesAsyncType = "batched-outcome"
)

var allowedUpdateOutcomesAsyncTypeEnumValues = []UpdateOutcomesAsyncType{
	UPDATEOUTCOMESASYNCTYPE_BATCHED_OUTCOME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateOutcomesAsyncType) GetAllowedValues() []UpdateOutcomesAsyncType {
	return allowedUpdateOutcomesAsyncTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateOutcomesAsyncType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateOutcomesAsyncType(value)
	return nil
}

// NewUpdateOutcomesAsyncTypeFromValue returns a pointer to a valid UpdateOutcomesAsyncType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateOutcomesAsyncTypeFromValue(v string) (*UpdateOutcomesAsyncType, error) {
	ev := UpdateOutcomesAsyncType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateOutcomesAsyncType: valid values are %v", v, allowedUpdateOutcomesAsyncTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateOutcomesAsyncType) IsValid() bool {
	for _, existing := range allowedUpdateOutcomesAsyncTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateOutcomesAsyncType value.
func (v UpdateOutcomesAsyncType) Ptr() *UpdateOutcomesAsyncType {
	return &v
}
