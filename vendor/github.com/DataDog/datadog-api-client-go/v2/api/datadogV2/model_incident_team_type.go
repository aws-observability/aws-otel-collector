// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTeamType Incident Team resource type.
type IncidentTeamType string

// List of IncidentTeamType.
const (
	INCIDENTTEAMTYPE_TEAMS IncidentTeamType = "teams"
)

var allowedIncidentTeamTypeEnumValues = []IncidentTeamType{
	INCIDENTTEAMTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTeamType) GetAllowedValues() []IncidentTeamType {
	return allowedIncidentTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTeamType(value)
	return nil
}

// NewIncidentTeamTypeFromValue returns a pointer to a valid IncidentTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTeamTypeFromValue(v string) (*IncidentTeamType, error) {
	ev := IncidentTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTeamType: valid values are %v", v, allowedIncidentTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTeamType) IsValid() bool {
	for _, existing := range allowedIncidentTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTeamType value.
func (v IncidentTeamType) Ptr() *IncidentTeamType {
	return &v
}
