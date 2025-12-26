// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncAttributesFrequency How often the sync process should be run. Defaults to `once` when not provided.
type TeamSyncAttributesFrequency string

// List of TeamSyncAttributesFrequency.
const (
	TEAMSYNCATTRIBUTESFREQUENCY_ONCE         TeamSyncAttributesFrequency = "once"
	TEAMSYNCATTRIBUTESFREQUENCY_CONTINUOUSLY TeamSyncAttributesFrequency = "continuously"
	TEAMSYNCATTRIBUTESFREQUENCY_PAUSED       TeamSyncAttributesFrequency = "paused"
)

var allowedTeamSyncAttributesFrequencyEnumValues = []TeamSyncAttributesFrequency{
	TEAMSYNCATTRIBUTESFREQUENCY_ONCE,
	TEAMSYNCATTRIBUTESFREQUENCY_CONTINUOUSLY,
	TEAMSYNCATTRIBUTESFREQUENCY_PAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncAttributesFrequency) GetAllowedValues() []TeamSyncAttributesFrequency {
	return allowedTeamSyncAttributesFrequencyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncAttributesFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncAttributesFrequency(value)
	return nil
}

// NewTeamSyncAttributesFrequencyFromValue returns a pointer to a valid TeamSyncAttributesFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncAttributesFrequencyFromValue(v string) (*TeamSyncAttributesFrequency, error) {
	ev := TeamSyncAttributesFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncAttributesFrequency: valid values are %v", v, allowedTeamSyncAttributesFrequencyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncAttributesFrequency) IsValid() bool {
	for _, existing := range allowedTeamSyncAttributesFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncAttributesFrequency value.
func (v TeamSyncAttributesFrequency) Ptr() *TeamSyncAttributesFrequency {
	return &v
}
