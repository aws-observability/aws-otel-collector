// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsField Supported teams field.
type TeamsField string

// List of TeamsField.
const (
	TEAMSFIELD_ID                    TeamsField = "id"
	TEAMSFIELD_NAME                  TeamsField = "name"
	TEAMSFIELD_HANDLE                TeamsField = "handle"
	TEAMSFIELD_SUMMARY               TeamsField = "summary"
	TEAMSFIELD_DESCRIPTION           TeamsField = "description"
	TEAMSFIELD_AVATAR                TeamsField = "avatar"
	TEAMSFIELD_BANNER                TeamsField = "banner"
	TEAMSFIELD_VISIBLE_MODULES       TeamsField = "visible_modules"
	TEAMSFIELD_HIDDEN_MODULES        TeamsField = "hidden_modules"
	TEAMSFIELD_CREATED_AT            TeamsField = "created_at"
	TEAMSFIELD_MODIFIED_AT           TeamsField = "modified_at"
	TEAMSFIELD_USER_COUNT            TeamsField = "user_count"
	TEAMSFIELD_LINK_COUNT            TeamsField = "link_count"
	TEAMSFIELD_TEAM_LINKS            TeamsField = "team_links"
	TEAMSFIELD_USER_TEAM_PERMISSIONS TeamsField = "user_team_permissions"
)

var allowedTeamsFieldEnumValues = []TeamsField{
	TEAMSFIELD_ID,
	TEAMSFIELD_NAME,
	TEAMSFIELD_HANDLE,
	TEAMSFIELD_SUMMARY,
	TEAMSFIELD_DESCRIPTION,
	TEAMSFIELD_AVATAR,
	TEAMSFIELD_BANNER,
	TEAMSFIELD_VISIBLE_MODULES,
	TEAMSFIELD_HIDDEN_MODULES,
	TEAMSFIELD_CREATED_AT,
	TEAMSFIELD_MODIFIED_AT,
	TEAMSFIELD_USER_COUNT,
	TEAMSFIELD_LINK_COUNT,
	TEAMSFIELD_TEAM_LINKS,
	TEAMSFIELD_USER_TEAM_PERMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamsField) GetAllowedValues() []TeamsField {
	return allowedTeamsFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamsField) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamsField(value)
	return nil
}

// NewTeamsFieldFromValue returns a pointer to a valid TeamsField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamsFieldFromValue(v string) (*TeamsField, error) {
	ev := TeamsField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamsField: valid values are %v", v, allowedTeamsFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamsField) IsValid() bool {
	for _, existing := range allowedTeamsFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamsField value.
func (v TeamsField) Ptr() *TeamsField {
	return &v
}
