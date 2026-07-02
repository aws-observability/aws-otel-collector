// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleRoutingMode The routing mode for the notification rule. `manual` sends notifications to the configured targets.
type NotificationRuleRoutingMode string

// List of NotificationRuleRoutingMode.
const (
	NOTIFICATIONRULEROUTINGMODE_MANUAL NotificationRuleRoutingMode = "manual"
)

var allowedNotificationRuleRoutingModeEnumValues = []NotificationRuleRoutingMode{
	NOTIFICATIONRULEROUTINGMODE_MANUAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationRuleRoutingMode) GetAllowedValues() []NotificationRuleRoutingMode {
	return allowedNotificationRuleRoutingModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationRuleRoutingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationRuleRoutingMode(value)
	return nil
}

// NewNotificationRuleRoutingModeFromValue returns a pointer to a valid NotificationRuleRoutingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationRuleRoutingModeFromValue(v string) (*NotificationRuleRoutingMode, error) {
	ev := NotificationRuleRoutingMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationRuleRoutingMode: valid values are %v", v, allowedNotificationRuleRoutingModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationRuleRoutingMode) IsValid() bool {
	for _, existing := range allowedNotificationRuleRoutingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationRuleRoutingMode value.
func (v NotificationRuleRoutingMode) Ptr() *NotificationRuleRoutingMode {
	return &v
}
