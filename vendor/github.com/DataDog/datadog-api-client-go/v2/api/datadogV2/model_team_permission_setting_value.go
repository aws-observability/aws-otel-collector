// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamPermissionSettingValue What type of user is allowed to perform the specified action
type TeamPermissionSettingValue string

// List of TeamPermissionSettingValue.
const (
	TEAMPERMISSIONSETTINGVALUE_ADMINS             TeamPermissionSettingValue = "admins"
	TEAMPERMISSIONSETTINGVALUE_MEMBERS            TeamPermissionSettingValue = "members"
	TEAMPERMISSIONSETTINGVALUE_ORGANIZATION       TeamPermissionSettingValue = "organization"
	TEAMPERMISSIONSETTINGVALUE_USER_ACCESS_MANAGE TeamPermissionSettingValue = "user_access_manage"
	TEAMPERMISSIONSETTINGVALUE_TEAMS_MANAGE       TeamPermissionSettingValue = "teams_manage"
)

var allowedTeamPermissionSettingValueEnumValues = []TeamPermissionSettingValue{
	TEAMPERMISSIONSETTINGVALUE_ADMINS,
	TEAMPERMISSIONSETTINGVALUE_MEMBERS,
	TEAMPERMISSIONSETTINGVALUE_ORGANIZATION,
	TEAMPERMISSIONSETTINGVALUE_USER_ACCESS_MANAGE,
	TEAMPERMISSIONSETTINGVALUE_TEAMS_MANAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamPermissionSettingValue) GetAllowedValues() []TeamPermissionSettingValue {
	return allowedTeamPermissionSettingValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamPermissionSettingValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamPermissionSettingValue(value)
	return nil
}

// NewTeamPermissionSettingValueFromValue returns a pointer to a valid TeamPermissionSettingValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamPermissionSettingValueFromValue(v string) (*TeamPermissionSettingValue, error) {
	ev := TeamPermissionSettingValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamPermissionSettingValue: valid values are %v", v, allowedTeamPermissionSettingValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamPermissionSettingValue) IsValid() bool {
	for _, existing := range allowedTeamPermissionSettingValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamPermissionSettingValue value.
func (v TeamPermissionSettingValue) Ptr() *TeamPermissionSettingValue {
	return &v
}
