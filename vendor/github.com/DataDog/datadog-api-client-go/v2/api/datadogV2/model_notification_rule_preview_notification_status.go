// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRulePreviewNotificationStatus The notification status for the given rule type. `SUCCESS` means a matching event was found and the notification was sent successfully. `DEFAULT` means no matching event was found and a default placeholder notification was sent instead. `ERROR` means an error occurred while sending the notification.
type NotificationRulePreviewNotificationStatus string

// List of NotificationRulePreviewNotificationStatus.
const (
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_SUCCESS NotificationRulePreviewNotificationStatus = "SUCCESS"
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_DEFAULT NotificationRulePreviewNotificationStatus = "DEFAULT"
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_ERROR   NotificationRulePreviewNotificationStatus = "ERROR"
)

var allowedNotificationRulePreviewNotificationStatusEnumValues = []NotificationRulePreviewNotificationStatus{
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_SUCCESS,
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_DEFAULT,
	NOTIFICATIONRULEPREVIEWNOTIFICATIONSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationRulePreviewNotificationStatus) GetAllowedValues() []NotificationRulePreviewNotificationStatus {
	return allowedNotificationRulePreviewNotificationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationRulePreviewNotificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationRulePreviewNotificationStatus(value)
	return nil
}

// NewNotificationRulePreviewNotificationStatusFromValue returns a pointer to a valid NotificationRulePreviewNotificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationRulePreviewNotificationStatusFromValue(v string) (*NotificationRulePreviewNotificationStatus, error) {
	ev := NotificationRulePreviewNotificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationRulePreviewNotificationStatus: valid values are %v", v, allowedNotificationRulePreviewNotificationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationRulePreviewNotificationStatus) IsValid() bool {
	for _, existing := range allowedNotificationRulePreviewNotificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationRulePreviewNotificationStatus value.
func (v NotificationRulePreviewNotificationStatus) Ptr() *NotificationRulePreviewNotificationStatus {
	return &v
}
