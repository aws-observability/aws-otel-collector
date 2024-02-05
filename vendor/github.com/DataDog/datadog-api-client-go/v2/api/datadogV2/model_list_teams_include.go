// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListTeamsInclude Included related resources optionally requested.
type ListTeamsInclude string

// List of ListTeamsInclude.
const (
	LISTTEAMSINCLUDE_TEAM_LINKS            ListTeamsInclude = "team_links"
	LISTTEAMSINCLUDE_USER_TEAM_PERMISSIONS ListTeamsInclude = "user_team_permissions"
)

var allowedListTeamsIncludeEnumValues = []ListTeamsInclude{
	LISTTEAMSINCLUDE_TEAM_LINKS,
	LISTTEAMSINCLUDE_USER_TEAM_PERMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListTeamsInclude) GetAllowedValues() []ListTeamsInclude {
	return allowedListTeamsIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListTeamsInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListTeamsInclude(value)
	return nil
}

// NewListTeamsIncludeFromValue returns a pointer to a valid ListTeamsInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListTeamsIncludeFromValue(v string) (*ListTeamsInclude, error) {
	ev := ListTeamsInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListTeamsInclude: valid values are %v", v, allowedListTeamsIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListTeamsInclude) IsValid() bool {
	for _, existing := range allowedListTeamsIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListTeamsInclude value.
func (v ListTeamsInclude) Ptr() *ListTeamsInclude {
	return &v
}
