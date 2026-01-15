// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelEmailFormatType Specifies the format of the e-mail that is sent for On-Call notifications
type NotificationChannelEmailFormatType string

// List of NotificationChannelEmailFormatType.
const (
	NOTIFICATIONCHANNELEMAILFORMATTYPE_HTML NotificationChannelEmailFormatType = "html"
	NOTIFICATIONCHANNELEMAILFORMATTYPE_TEXT NotificationChannelEmailFormatType = "text"
)

var allowedNotificationChannelEmailFormatTypeEnumValues = []NotificationChannelEmailFormatType{
	NOTIFICATIONCHANNELEMAILFORMATTYPE_HTML,
	NOTIFICATIONCHANNELEMAILFORMATTYPE_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationChannelEmailFormatType) GetAllowedValues() []NotificationChannelEmailFormatType {
	return allowedNotificationChannelEmailFormatTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationChannelEmailFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationChannelEmailFormatType(value)
	return nil
}

// NewNotificationChannelEmailFormatTypeFromValue returns a pointer to a valid NotificationChannelEmailFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationChannelEmailFormatTypeFromValue(v string) (*NotificationChannelEmailFormatType, error) {
	ev := NotificationChannelEmailFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationChannelEmailFormatType: valid values are %v", v, allowedNotificationChannelEmailFormatTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationChannelEmailFormatType) IsValid() bool {
	for _, existing := range allowedNotificationChannelEmailFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationChannelEmailFormatType value.
func (v NotificationChannelEmailFormatType) Ptr() *NotificationChannelEmailFormatType {
	return &v
}
