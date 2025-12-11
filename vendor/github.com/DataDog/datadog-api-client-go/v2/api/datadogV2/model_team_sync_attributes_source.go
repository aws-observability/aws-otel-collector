// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncAttributesSource The external source platform for team synchronization. Only "github" is supported.
type TeamSyncAttributesSource string

// List of TeamSyncAttributesSource.
const (
	TEAMSYNCATTRIBUTESSOURCE_GITHUB TeamSyncAttributesSource = "github"
)

var allowedTeamSyncAttributesSourceEnumValues = []TeamSyncAttributesSource{
	TEAMSYNCATTRIBUTESSOURCE_GITHUB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamSyncAttributesSource) GetAllowedValues() []TeamSyncAttributesSource {
	return allowedTeamSyncAttributesSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamSyncAttributesSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamSyncAttributesSource(value)
	return nil
}

// NewTeamSyncAttributesSourceFromValue returns a pointer to a valid TeamSyncAttributesSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamSyncAttributesSourceFromValue(v string) (*TeamSyncAttributesSource, error) {
	ev := TeamSyncAttributesSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamSyncAttributesSource: valid values are %v", v, allowedTeamSyncAttributesSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamSyncAttributesSource) IsValid() bool {
	for _, existing := range allowedTeamSyncAttributesSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamSyncAttributesSource value.
func (v TeamSyncAttributesSource) Ptr() *TeamSyncAttributesSource {
	return &v
}
