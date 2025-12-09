// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamConnectionType Team connection resource type.
type TeamConnectionType string

// List of TeamConnectionType.
const (
	TEAMCONNECTIONTYPE_TEAM_CONNECTION TeamConnectionType = "team_connection"
)

var allowedTeamConnectionTypeEnumValues = []TeamConnectionType{
	TEAMCONNECTIONTYPE_TEAM_CONNECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamConnectionType) GetAllowedValues() []TeamConnectionType {
	return allowedTeamConnectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamConnectionType(value)
	return nil
}

// NewTeamConnectionTypeFromValue returns a pointer to a valid TeamConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamConnectionTypeFromValue(v string) (*TeamConnectionType, error) {
	ev := TeamConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamConnectionType: valid values are %v", v, allowedTeamConnectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamConnectionType) IsValid() bool {
	for _, existing := range allowedTeamConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamConnectionType value.
func (v TeamConnectionType) Ptr() *TeamConnectionType {
	return &v
}
