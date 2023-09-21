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

// UserTeamType Team membership type
type UserTeamType string

// List of UserTeamType.
const (
	USERTEAMTYPE_TEAM_MEMBERSHIPS UserTeamType = "team_memberships"
)

var allowedUserTeamTypeEnumValues = []UserTeamType{
	USERTEAMTYPE_TEAM_MEMBERSHIPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTeamType) GetAllowedValues() []UserTeamType {
	return allowedUserTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTeamType(value)
	return nil
}

// NewUserTeamTypeFromValue returns a pointer to a valid UserTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTeamTypeFromValue(v string) (*UserTeamType, error) {
	ev := UserTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTeamType: valid values are %v", v, allowedUserTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTeamType) IsValid() bool {
	for _, existing := range allowedUserTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTeamType value.
func (v UserTeamType) Ptr() *UserTeamType {
	return &v
}
<<<<<<< HEAD

// NullableUserTeamType handles when a null is used for UserTeamType.
type NullableUserTeamType struct {
	value *UserTeamType
	isSet bool
}

// Get returns the associated value.
func (v NullableUserTeamType) Get() *UserTeamType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUserTeamType) Set(val *UserTeamType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUserTeamType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUserTeamType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUserTeamType initializes the struct as if Set has been called.
func NewNullableUserTeamType(val *UserTeamType) *NullableUserTeamType {
	return &NullableUserTeamType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUserTeamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUserTeamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
