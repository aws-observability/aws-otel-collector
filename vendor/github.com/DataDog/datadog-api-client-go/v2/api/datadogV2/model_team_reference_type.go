// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamReferenceType Teams resource type.
type TeamReferenceType string

// List of TeamReferenceType.
const (
	TEAMREFERENCETYPE_TEAMS TeamReferenceType = "teams"
)

var allowedTeamReferenceTypeEnumValues = []TeamReferenceType{
	TEAMREFERENCETYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamReferenceType) GetAllowedValues() []TeamReferenceType {
	return allowedTeamReferenceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamReferenceType(value)
	return nil
}

// NewTeamReferenceTypeFromValue returns a pointer to a valid TeamReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamReferenceTypeFromValue(v string) (*TeamReferenceType, error) {
	ev := TeamReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamReferenceType: valid values are %v", v, allowedTeamReferenceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamReferenceType) IsValid() bool {
	for _, existing := range allowedTeamReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamReferenceType value.
func (v TeamReferenceType) Ptr() *TeamReferenceType {
	return &v
}
