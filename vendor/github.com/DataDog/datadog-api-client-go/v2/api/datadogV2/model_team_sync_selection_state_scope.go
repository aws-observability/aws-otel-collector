// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncSelectionStateScope The scope of the selection. When set to `subtree`,
// synchronization includes the referenced team or
// organization and everything nested under it.
type TeamSyncSelectionStateScope string

// List of TeamSyncSelectionStateScope.
const (
	TEAMSYNCSELECTIONSTATESCOPE_SUBTREE TeamSyncSelectionStateScope = "subtree"
)

var allowedTeamSyncSelectionStateScopeEnumValues = []TeamSyncSelectionStateScope{
	TEAMSYNCSELECTIONSTATESCOPE_SUBTREE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncSelectionStateScope) GetAllowedValues() []TeamSyncSelectionStateScope {
	return allowedTeamSyncSelectionStateScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncSelectionStateScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncSelectionStateScope(value)
	return nil
}

// NewTeamSyncSelectionStateScopeFromValue returns a pointer to a valid TeamSyncSelectionStateScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncSelectionStateScopeFromValue(v string) (*TeamSyncSelectionStateScope, error) {
	ev := TeamSyncSelectionStateScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncSelectionStateScope: valid values are %v", v, allowedTeamSyncSelectionStateScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncSelectionStateScope) IsValid() bool {
	for _, existing := range allowedTeamSyncSelectionStateScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncSelectionStateScope value.
func (v TeamSyncSelectionStateScope) Ptr() *TeamSyncSelectionStateScope {
	return &v
}
