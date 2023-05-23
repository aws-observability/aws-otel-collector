// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// TeamPermissionSettingSerializerAction The identifier for the action
type TeamPermissionSettingSerializerAction string

// List of TeamPermissionSettingSerializerAction.
const (
	TEAMPERMISSIONSETTINGSERIALIZERACTION_MANAGE_MEMBERSHIP TeamPermissionSettingSerializerAction = "manage_membership"
	TEAMPERMISSIONSETTINGSERIALIZERACTION_EDIT              TeamPermissionSettingSerializerAction = "edit"
)

var allowedTeamPermissionSettingSerializerActionEnumValues = []TeamPermissionSettingSerializerAction{
	TEAMPERMISSIONSETTINGSERIALIZERACTION_MANAGE_MEMBERSHIP,
	TEAMPERMISSIONSETTINGSERIALIZERACTION_EDIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamPermissionSettingSerializerAction) GetAllowedValues() []TeamPermissionSettingSerializerAction {
	return allowedTeamPermissionSettingSerializerActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamPermissionSettingSerializerAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamPermissionSettingSerializerAction(value)
	return nil
}

// NewTeamPermissionSettingSerializerActionFromValue returns a pointer to a valid TeamPermissionSettingSerializerAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamPermissionSettingSerializerActionFromValue(v string) (*TeamPermissionSettingSerializerAction, error) {
	ev := TeamPermissionSettingSerializerAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamPermissionSettingSerializerAction: valid values are %v", v, allowedTeamPermissionSettingSerializerActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamPermissionSettingSerializerAction) IsValid() bool {
	for _, existing := range allowedTeamPermissionSettingSerializerActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamPermissionSettingSerializerAction value.
func (v TeamPermissionSettingSerializerAction) Ptr() *TeamPermissionSettingSerializerAction {
	return &v
}

// NullableTeamPermissionSettingSerializerAction handles when a null is used for TeamPermissionSettingSerializerAction.
type NullableTeamPermissionSettingSerializerAction struct {
	value *TeamPermissionSettingSerializerAction
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamPermissionSettingSerializerAction) Get() *TeamPermissionSettingSerializerAction {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamPermissionSettingSerializerAction) Set(val *TeamPermissionSettingSerializerAction) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamPermissionSettingSerializerAction) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableTeamPermissionSettingSerializerAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamPermissionSettingSerializerAction initializes the struct as if Set has been called.
func NewNullableTeamPermissionSettingSerializerAction(val *TeamPermissionSettingSerializerAction) *NullableTeamPermissionSettingSerializerAction {
	return &NullableTeamPermissionSettingSerializerAction{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamPermissionSettingSerializerAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamPermissionSettingSerializerAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
