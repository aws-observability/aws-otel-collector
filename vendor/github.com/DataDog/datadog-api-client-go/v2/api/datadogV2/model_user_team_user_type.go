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

// UserTeamUserType User team user type
type UserTeamUserType string

// List of UserTeamUserType.
const (
	USERTEAMUSERTYPE_USERS UserTeamUserType = "users"
)

var allowedUserTeamUserTypeEnumValues = []UserTeamUserType{
	USERTEAMUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTeamUserType) GetAllowedValues() []UserTeamUserType {
	return allowedUserTeamUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTeamUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTeamUserType(value)
	return nil
}

// NewUserTeamUserTypeFromValue returns a pointer to a valid UserTeamUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTeamUserTypeFromValue(v string) (*UserTeamUserType, error) {
	ev := UserTeamUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTeamUserType: valid values are %v", v, allowedUserTeamUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTeamUserType) IsValid() bool {
	for _, existing := range allowedUserTeamUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTeamUserType value.
func (v UserTeamUserType) Ptr() *UserTeamUserType {
	return &v
}
<<<<<<< HEAD

// NullableUserTeamUserType handles when a null is used for UserTeamUserType.
type NullableUserTeamUserType struct {
	value *UserTeamUserType
	isSet bool
}

// Get returns the associated value.
func (v NullableUserTeamUserType) Get() *UserTeamUserType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUserTeamUserType) Set(val *UserTeamUserType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUserTeamUserType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUserTeamUserType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUserTeamUserType initializes the struct as if Set has been called.
func NewNullableUserTeamUserType(val *UserTeamUserType) *NullableUserTeamUserType {
	return &NullableUserTeamUserType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUserTeamUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUserTeamUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
