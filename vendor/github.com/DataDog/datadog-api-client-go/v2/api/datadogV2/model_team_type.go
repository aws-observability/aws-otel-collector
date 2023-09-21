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

// TeamType Team type
type TeamType string

// List of TeamType.
const (
	TEAMTYPE_TEAM TeamType = "team"
)

var allowedTeamTypeEnumValues = []TeamType{
	TEAMTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamType) GetAllowedValues() []TeamType {
	return allowedTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamType(value)
	return nil
}

// NewTeamTypeFromValue returns a pointer to a valid TeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamTypeFromValue(v string) (*TeamType, error) {
	ev := TeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamType: valid values are %v", v, allowedTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamType) IsValid() bool {
	for _, existing := range allowedTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamType value.
func (v TeamType) Ptr() *TeamType {
	return &v
}
<<<<<<< HEAD

// NullableTeamType handles when a null is used for TeamType.
type NullableTeamType struct {
	value *TeamType
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamType) Get() *TeamType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamType) Set(val *TeamType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableTeamType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamType initializes the struct as if Set has been called.
func NewNullableTeamType(val *TeamType) *NullableTeamType {
	return &NullableTeamType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
