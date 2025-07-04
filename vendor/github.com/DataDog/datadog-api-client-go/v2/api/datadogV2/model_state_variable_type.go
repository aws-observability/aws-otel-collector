// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StateVariableType The state variable type.
type StateVariableType string

// List of StateVariableType.
const (
	STATEVARIABLETYPE_STATEVARIABLE StateVariableType = "stateVariable"
)

var allowedStateVariableTypeEnumValues = []StateVariableType{
	STATEVARIABLETYPE_STATEVARIABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StateVariableType) GetAllowedValues() []StateVariableType {
	return allowedStateVariableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StateVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StateVariableType(value)
	return nil
}

// NewStateVariableTypeFromValue returns a pointer to a valid StateVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStateVariableTypeFromValue(v string) (*StateVariableType, error) {
	ev := StateVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StateVariableType: valid values are %v", v, allowedStateVariableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StateVariableType) IsValid() bool {
	for _, existing := range allowedStateVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateVariableType value.
func (v StateVariableType) Ptr() *StateVariableType {
	return &v
}
