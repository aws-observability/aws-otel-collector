// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// IncidentTodoType Todo resource type.
type IncidentTodoType string

// List of IncidentTodoType.
const (
	INCIDENTTODOTYPE_INCIDENT_TODOS IncidentTodoType = "incident_todos"
)

var allowedIncidentTodoTypeEnumValues = []IncidentTodoType{
	INCIDENTTODOTYPE_INCIDENT_TODOS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTodoType) GetAllowedValues() []IncidentTodoType {
	return allowedIncidentTodoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTodoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTodoType(value)
	return nil
}

// NewIncidentTodoTypeFromValue returns a pointer to a valid IncidentTodoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTodoTypeFromValue(v string) (*IncidentTodoType, error) {
	ev := IncidentTodoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTodoType: valid values are %v", v, allowedIncidentTodoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTodoType) IsValid() bool {
	for _, existing := range allowedIncidentTodoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTodoType value.
func (v IncidentTodoType) Ptr() *IncidentTodoType {
	return &v
}
<<<<<<< HEAD

// NullableIncidentTodoType handles when a null is used for IncidentTodoType.
type NullableIncidentTodoType struct {
	value *IncidentTodoType
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentTodoType) Get() *IncidentTodoType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentTodoType) Set(val *IncidentTodoType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentTodoType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableIncidentTodoType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentTodoType initializes the struct as if Set has been called.
func NewNullableIncidentTodoType(val *IncidentTodoType) *NullableIncidentTodoType {
	return &NullableIncidentTodoType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentTodoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentTodoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
