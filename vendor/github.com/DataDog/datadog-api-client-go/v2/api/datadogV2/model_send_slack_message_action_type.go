// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SendSlackMessageActionType Indicates that the action is a send Slack message action.
type SendSlackMessageActionType string

// List of SendSlackMessageActionType.
const (
	SENDSLACKMESSAGEACTIONTYPE_SEND_SLACK_MESSAGE SendSlackMessageActionType = "send_slack_message"
)

var allowedSendSlackMessageActionTypeEnumValues = []SendSlackMessageActionType{
	SENDSLACKMESSAGEACTIONTYPE_SEND_SLACK_MESSAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SendSlackMessageActionType) GetAllowedValues() []SendSlackMessageActionType {
	return allowedSendSlackMessageActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SendSlackMessageActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SendSlackMessageActionType(value)
	return nil
}

// NewSendSlackMessageActionTypeFromValue returns a pointer to a valid SendSlackMessageActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSendSlackMessageActionTypeFromValue(v string) (*SendSlackMessageActionType, error) {
	ev := SendSlackMessageActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SendSlackMessageActionType: valid values are %v", v, allowedSendSlackMessageActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SendSlackMessageActionType) IsValid() bool {
	for _, existing := range allowedSendSlackMessageActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SendSlackMessageActionType value.
func (v SendSlackMessageActionType) Ptr() *SendSlackMessageActionType {
	return &v
}
