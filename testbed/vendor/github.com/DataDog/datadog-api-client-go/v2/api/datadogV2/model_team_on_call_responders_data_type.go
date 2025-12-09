// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersDataType Represents the resource type for a group of users assigned to handle on-call duties within a team.
type TeamOnCallRespondersDataType string

// List of TeamOnCallRespondersDataType.
const (
	TEAMONCALLRESPONDERSDATATYPE_TEAM_ONCALL_RESPONDERS TeamOnCallRespondersDataType = "team_oncall_responders"
)

var allowedTeamOnCallRespondersDataTypeEnumValues = []TeamOnCallRespondersDataType{
	TEAMONCALLRESPONDERSDATATYPE_TEAM_ONCALL_RESPONDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamOnCallRespondersDataType) GetAllowedValues() []TeamOnCallRespondersDataType {
	return allowedTeamOnCallRespondersDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamOnCallRespondersDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamOnCallRespondersDataType(value)
	return nil
}

// NewTeamOnCallRespondersDataTypeFromValue returns a pointer to a valid TeamOnCallRespondersDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamOnCallRespondersDataTypeFromValue(v string) (*TeamOnCallRespondersDataType, error) {
	ev := TeamOnCallRespondersDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamOnCallRespondersDataType: valid values are %v", v, allowedTeamOnCallRespondersDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamOnCallRespondersDataType) IsValid() bool {
	for _, existing := range allowedTeamOnCallRespondersDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamOnCallRespondersDataType value.
func (v TeamOnCallRespondersDataType) Ptr() *TeamOnCallRespondersDataType {
	return &v
}
