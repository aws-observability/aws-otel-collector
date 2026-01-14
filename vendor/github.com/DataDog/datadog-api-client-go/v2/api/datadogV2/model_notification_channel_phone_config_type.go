// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationChannelPhoneConfigType Indicates that the notification channel is a phone
type NotificationChannelPhoneConfigType string

// List of NotificationChannelPhoneConfigType.
const (
	NOTIFICATIONCHANNELPHONECONFIGTYPE_PHONE NotificationChannelPhoneConfigType = "phone"
)

var allowedNotificationChannelPhoneConfigTypeEnumValues = []NotificationChannelPhoneConfigType{
	NOTIFICATIONCHANNELPHONECONFIGTYPE_PHONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationChannelPhoneConfigType) GetAllowedValues() []NotificationChannelPhoneConfigType {
	return allowedNotificationChannelPhoneConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationChannelPhoneConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationChannelPhoneConfigType(value)
	return nil
}

// NewNotificationChannelPhoneConfigTypeFromValue returns a pointer to a valid NotificationChannelPhoneConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationChannelPhoneConfigTypeFromValue(v string) (*NotificationChannelPhoneConfigType, error) {
	ev := NotificationChannelPhoneConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationChannelPhoneConfigType: valid values are %v", v, allowedNotificationChannelPhoneConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationChannelPhoneConfigType) IsValid() bool {
	for _, existing := range allowedNotificationChannelPhoneConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationChannelPhoneConfigType value.
func (v NotificationChannelPhoneConfigType) Ptr() *NotificationChannelPhoneConfigType {
	return &v
}
