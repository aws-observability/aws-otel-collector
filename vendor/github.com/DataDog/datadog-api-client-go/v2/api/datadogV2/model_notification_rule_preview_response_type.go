// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRulePreviewResponseType The type of the notification preview response.
type NotificationRulePreviewResponseType string

// List of NotificationRulePreviewResponseType.
const (
	NOTIFICATIONRULEPREVIEWRESPONSETYPE_NOTIFICATION_PREVIEW_RESPONSE NotificationRulePreviewResponseType = "notification_preview_response"
)

var allowedNotificationRulePreviewResponseTypeEnumValues = []NotificationRulePreviewResponseType{
	NOTIFICATIONRULEPREVIEWRESPONSETYPE_NOTIFICATION_PREVIEW_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationRulePreviewResponseType) GetAllowedValues() []NotificationRulePreviewResponseType {
	return allowedNotificationRulePreviewResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationRulePreviewResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationRulePreviewResponseType(value)
	return nil
}

// NewNotificationRulePreviewResponseTypeFromValue returns a pointer to a valid NotificationRulePreviewResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationRulePreviewResponseTypeFromValue(v string) (*NotificationRulePreviewResponseType, error) {
	ev := NotificationRulePreviewResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationRulePreviewResponseType: valid values are %v", v, allowedNotificationRulePreviewResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationRulePreviewResponseType) IsValid() bool {
	for _, existing := range allowedNotificationRulePreviewResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationRulePreviewResponseType value.
func (v NotificationRulePreviewResponseType) Ptr() *NotificationRulePreviewResponseType {
	return &v
}
