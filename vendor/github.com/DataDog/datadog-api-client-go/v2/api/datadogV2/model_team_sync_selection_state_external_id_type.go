// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncSelectionStateExternalIdType The type of external identifier for the selection state item.
// For GitHub synchronization, the allowed values are `team` and
// `organization`.
type TeamSyncSelectionStateExternalIdType string

// List of TeamSyncSelectionStateExternalIdType.
const (
	TEAMSYNCSELECTIONSTATEEXTERNALIDTYPE_TEAM         TeamSyncSelectionStateExternalIdType = "team"
	TEAMSYNCSELECTIONSTATEEXTERNALIDTYPE_ORGANIZATION TeamSyncSelectionStateExternalIdType = "organization"
)

var allowedTeamSyncSelectionStateExternalIdTypeEnumValues = []TeamSyncSelectionStateExternalIdType{
	TEAMSYNCSELECTIONSTATEEXTERNALIDTYPE_TEAM,
	TEAMSYNCSELECTIONSTATEEXTERNALIDTYPE_ORGANIZATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncSelectionStateExternalIdType) GetAllowedValues() []TeamSyncSelectionStateExternalIdType {
	return allowedTeamSyncSelectionStateExternalIdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncSelectionStateExternalIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncSelectionStateExternalIdType(value)
	return nil
}

// NewTeamSyncSelectionStateExternalIdTypeFromValue returns a pointer to a valid TeamSyncSelectionStateExternalIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncSelectionStateExternalIdTypeFromValue(v string) (*TeamSyncSelectionStateExternalIdType, error) {
	ev := TeamSyncSelectionStateExternalIdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncSelectionStateExternalIdType: valid values are %v", v, allowedTeamSyncSelectionStateExternalIdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncSelectionStateExternalIdType) IsValid() bool {
	for _, existing := range allowedTeamSyncSelectionStateExternalIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncSelectionStateExternalIdType value.
func (v TeamSyncSelectionStateExternalIdType) Ptr() *TeamSyncSelectionStateExternalIdType {
	return &v
}
