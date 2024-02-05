// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamType Team type
type TeamType string

// List of TeamType.
const (
	TEAMTYPE_TEAM TeamType = "team"
)

var allowedTeamTypeEnumValues = []TeamType{
	TEAMTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamType) GetAllowedValues() []TeamType {
	return allowedTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamType(value)
	return nil
}

// NewTeamTypeFromValue returns a pointer to a valid TeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamTypeFromValue(v string) (*TeamType, error) {
	ev := TeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamType: valid values are %v", v, allowedTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamType) IsValid() bool {
	for _, existing := range allowedTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamType value.
func (v TeamType) Ptr() *TeamType {
	return &v
}
