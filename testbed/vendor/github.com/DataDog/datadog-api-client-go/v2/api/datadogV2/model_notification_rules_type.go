// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRulesType The rule type associated to notification rules.
type NotificationRulesType string

// List of NotificationRulesType.
const (
	NOTIFICATIONRULESTYPE_NOTIFICATION_RULES NotificationRulesType = "notification_rules"
)

var allowedNotificationRulesTypeEnumValues = []NotificationRulesType{
	NOTIFICATIONRULESTYPE_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationRulesType) GetAllowedValues() []NotificationRulesType {
	return allowedNotificationRulesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationRulesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationRulesType(value)
	return nil
}

// NewNotificationRulesTypeFromValue returns a pointer to a valid NotificationRulesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationRulesTypeFromValue(v string) (*NotificationRulesType, error) {
	ev := NotificationRulesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationRulesType: valid values are %v", v, allowedNotificationRulesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationRulesType) IsValid() bool {
	for _, existing := range allowedNotificationRulesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationRulesType value.
func (v NotificationRulesType) Ptr() *NotificationRulesType {
	return &v
}
