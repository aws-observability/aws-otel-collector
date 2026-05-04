// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelEmailConfigType Indicates that the notification channel is an e-mail address
type NotificationChannelEmailConfigType string

// List of NotificationChannelEmailConfigType.
const (
	NOTIFICATIONCHANNELEMAILCONFIGTYPE_EMAIL NotificationChannelEmailConfigType = "email"
)

var allowedNotificationChannelEmailConfigTypeEnumValues = []NotificationChannelEmailConfigType{
	NOTIFICATIONCHANNELEMAILCONFIGTYPE_EMAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationChannelEmailConfigType) GetAllowedValues() []NotificationChannelEmailConfigType {
	return allowedNotificationChannelEmailConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationChannelEmailConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationChannelEmailConfigType(value)
	return nil
}

// NewNotificationChannelEmailConfigTypeFromValue returns a pointer to a valid NotificationChannelEmailConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationChannelEmailConfigTypeFromValue(v string) (*NotificationChannelEmailConfigType, error) {
	ev := NotificationChannelEmailConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationChannelEmailConfigType: valid values are %v", v, allowedNotificationChannelEmailConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationChannelEmailConfigType) IsValid() bool {
	for _, existing := range allowedNotificationChannelEmailConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationChannelEmailConfigType value.
func (v NotificationChannelEmailConfigType) Ptr() *NotificationChannelEmailConfigType {
	return &v
}
