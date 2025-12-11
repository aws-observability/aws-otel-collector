// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncBulkType Team sync bulk type.
type TeamSyncBulkType string

// List of TeamSyncBulkType.
const (
	TEAMSYNCBULKTYPE_TEAM_SYNC_BULK TeamSyncBulkType = "team_sync_bulk"
)

var allowedTeamSyncBulkTypeEnumValues = []TeamSyncBulkType{
	TEAMSYNCBULKTYPE_TEAM_SYNC_BULK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncBulkType) GetAllowedValues() []TeamSyncBulkType {
	return allowedTeamSyncBulkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncBulkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncBulkType(value)
	return nil
}

// NewTeamSyncBulkTypeFromValue returns a pointer to a valid TeamSyncBulkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncBulkTypeFromValue(v string) (*TeamSyncBulkType, error) {
	ev := TeamSyncBulkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncBulkType: valid values are %v", v, allowedTeamSyncBulkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncBulkType) IsValid() bool {
	for _, existing := range allowedTeamSyncBulkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncBulkType value.
func (v TeamSyncBulkType) Ptr() *TeamSyncBulkType {
	return &v
}
