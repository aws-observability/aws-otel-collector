// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SendTeamsMessageActionType Indicates that the action is a send Microsoft Teams message action.
type SendTeamsMessageActionType string

// List of SendTeamsMessageActionType.
const (
	SENDTEAMSMESSAGEACTIONTYPE_SEND_TEAMS_MESSAGE SendTeamsMessageActionType = "send_teams_message"
)

var allowedSendTeamsMessageActionTypeEnumValues = []SendTeamsMessageActionType{
	SENDTEAMSMESSAGEACTIONTYPE_SEND_TEAMS_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SendTeamsMessageActionType) GetAllowedValues() []SendTeamsMessageActionType {
	return allowedSendTeamsMessageActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SendTeamsMessageActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SendTeamsMessageActionType(value)
	return nil
}

// NewSendTeamsMessageActionTypeFromValue returns a pointer to a valid SendTeamsMessageActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSendTeamsMessageActionTypeFromValue(v string) (*SendTeamsMessageActionType, error) {
	ev := SendTeamsMessageActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SendTeamsMessageActionType: valid values are %v", v, allowedSendTeamsMessageActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SendTeamsMessageActionType) IsValid() bool {
	for _, existing := range allowedSendTeamsMessageActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SendTeamsMessageActionType value.
func (v SendTeamsMessageActionType) Ptr() *SendTeamsMessageActionType {
	return &v
}
