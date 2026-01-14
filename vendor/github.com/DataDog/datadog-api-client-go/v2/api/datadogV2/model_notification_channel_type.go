// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelType Indicates that the resource is of type 'notification_channels'.
type NotificationChannelType string

// List of NotificationChannelType.
const (
	NOTIFICATIONCHANNELTYPE_NOTIFICATION_CHANNELS NotificationChannelType = "notification_channels"
)

var allowedNotificationChannelTypeEnumValues = []NotificationChannelType{
	NOTIFICATIONCHANNELTYPE_NOTIFICATION_CHANNELS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationChannelType) GetAllowedValues() []NotificationChannelType {
	return allowedNotificationChannelTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationChannelType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationChannelType(value)
	return nil
}

// NewNotificationChannelTypeFromValue returns a pointer to a valid NotificationChannelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationChannelTypeFromValue(v string) (*NotificationChannelType, error) {
	ev := NotificationChannelType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationChannelType: valid values are %v", v, allowedNotificationChannelTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationChannelType) IsValid() bool {
	for _, existing := range allowedNotificationChannelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationChannelType value.
func (v NotificationChannelType) Ptr() *NotificationChannelType {
	return &v
}
