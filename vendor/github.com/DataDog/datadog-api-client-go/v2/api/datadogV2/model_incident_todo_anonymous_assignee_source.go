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

// IncidentTodoAnonymousAssigneeSource The source of the anonymous assignee.
type IncidentTodoAnonymousAssigneeSource string

// List of IncidentTodoAnonymousAssigneeSource.
const (
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_SLACK           IncidentTodoAnonymousAssigneeSource = "slack"
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_MICROSOFT_TEAMS IncidentTodoAnonymousAssigneeSource = "microsoft_teams"
)

var allowedIncidentTodoAnonymousAssigneeSourceEnumValues = []IncidentTodoAnonymousAssigneeSource{
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_SLACK,
	INCIDENTTODOANONYMOUSASSIGNEESOURCE_MICROSOFT_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTodoAnonymousAssigneeSource) GetAllowedValues() []IncidentTodoAnonymousAssigneeSource {
	return allowedIncidentTodoAnonymousAssigneeSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTodoAnonymousAssigneeSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTodoAnonymousAssigneeSource(value)
	return nil
}

// NewIncidentTodoAnonymousAssigneeSourceFromValue returns a pointer to a valid IncidentTodoAnonymousAssigneeSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTodoAnonymousAssigneeSourceFromValue(v string) (*IncidentTodoAnonymousAssigneeSource, error) {
	ev := IncidentTodoAnonymousAssigneeSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTodoAnonymousAssigneeSource: valid values are %v", v, allowedIncidentTodoAnonymousAssigneeSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTodoAnonymousAssigneeSource) IsValid() bool {
	for _, existing := range allowedIncidentTodoAnonymousAssigneeSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTodoAnonymousAssigneeSource value.
func (v IncidentTodoAnonymousAssigneeSource) Ptr() *IncidentTodoAnonymousAssigneeSource {
	return &v
}
<<<<<<< HEAD

// NullableIncidentTodoAnonymousAssigneeSource handles when a null is used for IncidentTodoAnonymousAssigneeSource.
type NullableIncidentTodoAnonymousAssigneeSource struct {
	value *IncidentTodoAnonymousAssigneeSource
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentTodoAnonymousAssigneeSource) Get() *IncidentTodoAnonymousAssigneeSource {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentTodoAnonymousAssigneeSource) Set(val *IncidentTodoAnonymousAssigneeSource) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentTodoAnonymousAssigneeSource) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableIncidentTodoAnonymousAssigneeSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentTodoAnonymousAssigneeSource initializes the struct as if Set has been called.
func NewNullableIncidentTodoAnonymousAssigneeSource(val *IncidentTodoAnonymousAssigneeSource) *NullableIncidentTodoAnonymousAssigneeSource {
	return &NullableIncidentTodoAnonymousAssigneeSource{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentTodoAnonymousAssigneeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentTodoAnonymousAssigneeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
