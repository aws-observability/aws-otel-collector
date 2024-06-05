// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamPermissionSettingType Team permission setting type
type TeamPermissionSettingType string

// List of TeamPermissionSettingType.
const (
	TEAMPERMISSIONSETTINGTYPE_TEAM_PERMISSION_SETTINGS TeamPermissionSettingType = "team_permission_settings"
)

var allowedTeamPermissionSettingTypeEnumValues = []TeamPermissionSettingType{
	TEAMPERMISSIONSETTINGTYPE_TEAM_PERMISSION_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamPermissionSettingType) GetAllowedValues() []TeamPermissionSettingType {
	return allowedTeamPermissionSettingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamPermissionSettingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamPermissionSettingType(value)
	return nil
}

// NewTeamPermissionSettingTypeFromValue returns a pointer to a valid TeamPermissionSettingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamPermissionSettingTypeFromValue(v string) (*TeamPermissionSettingType, error) {
	ev := TeamPermissionSettingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamPermissionSettingType: valid values are %v", v, allowedTeamPermissionSettingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamPermissionSettingType) IsValid() bool {
	for _, existing := range allowedTeamPermissionSettingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamPermissionSettingType value.
func (v TeamPermissionSettingType) Ptr() *TeamPermissionSettingType {
	return &v
}
