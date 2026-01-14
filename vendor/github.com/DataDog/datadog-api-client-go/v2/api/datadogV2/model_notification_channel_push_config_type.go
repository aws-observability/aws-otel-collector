// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelPushConfigType Indicates that the notification channel is a mobile device for push notifications
type NotificationChannelPushConfigType string

// List of NotificationChannelPushConfigType.
const (
	NOTIFICATIONCHANNELPUSHCONFIGTYPE_PUSH NotificationChannelPushConfigType = "push"
)

var allowedNotificationChannelPushConfigTypeEnumValues = []NotificationChannelPushConfigType{
	NOTIFICATIONCHANNELPUSHCONFIGTYPE_PUSH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationChannelPushConfigType) GetAllowedValues() []NotificationChannelPushConfigType {
	return allowedNotificationChannelPushConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationChannelPushConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationChannelPushConfigType(value)
	return nil
}

// NewNotificationChannelPushConfigTypeFromValue returns a pointer to a valid NotificationChannelPushConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationChannelPushConfigTypeFromValue(v string) (*NotificationChannelPushConfigType, error) {
	ev := NotificationChannelPushConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationChannelPushConfigType: valid values are %v", v, allowedNotificationChannelPushConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationChannelPushConfigType) IsValid() bool {
	for _, existing := range allowedNotificationChannelPushConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationChannelPushConfigType value.
func (v NotificationChannelPushConfigType) Ptr() *NotificationChannelPushConfigType {
	return &v
}
