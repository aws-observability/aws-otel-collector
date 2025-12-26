// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkType Team hierarchy link type
type TeamHierarchyLinkType string

// List of TeamHierarchyLinkType.
const (
	TEAMHIERARCHYLINKTYPE_TEAM_HIERARCHY_LINKS TeamHierarchyLinkType = "team_hierarchy_links"
)

var allowedTeamHierarchyLinkTypeEnumValues = []TeamHierarchyLinkType{
	TEAMHIERARCHYLINKTYPE_TEAM_HIERARCHY_LINKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamHierarchyLinkType) GetAllowedValues() []TeamHierarchyLinkType {
	return allowedTeamHierarchyLinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamHierarchyLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamHierarchyLinkType(value)
	return nil
}

// NewTeamHierarchyLinkTypeFromValue returns a pointer to a valid TeamHierarchyLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamHierarchyLinkTypeFromValue(v string) (*TeamHierarchyLinkType, error) {
	ev := TeamHierarchyLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamHierarchyLinkType: valid values are %v", v, allowedTeamHierarchyLinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamHierarchyLinkType) IsValid() bool {
	for _, existing := range allowedTeamHierarchyLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamHierarchyLinkType value.
func (v TeamHierarchyLinkType) Ptr() *TeamHierarchyLinkType {
	return &v
}
