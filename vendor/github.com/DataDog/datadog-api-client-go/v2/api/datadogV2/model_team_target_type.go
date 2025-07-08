// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamTargetType Indicates that the resource is of type `teams`.
type TeamTargetType string

// List of TeamTargetType.
const (
	TEAMTARGETTYPE_TEAMS TeamTargetType = "teams"
)

var allowedTeamTargetTypeEnumValues = []TeamTargetType{
	TEAMTARGETTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamTargetType) GetAllowedValues() []TeamTargetType {
	return allowedTeamTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamTargetType(value)
	return nil
}

// NewTeamTargetTypeFromValue returns a pointer to a valid TeamTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamTargetTypeFromValue(v string) (*TeamTargetType, error) {
	ev := TeamTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamTargetType: valid values are %v", v, allowedTeamTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamTargetType) IsValid() bool {
	for _, existing := range allowedTeamTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamTargetType value.
func (v TeamTargetType) Ptr() *TeamTargetType {
	return &v
}
