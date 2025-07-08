// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleResourceType Monitor notification rule resource type.
type MonitorNotificationRuleResourceType string

// List of MonitorNotificationRuleResourceType.
const (
	MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE MonitorNotificationRuleResourceType = "monitor-notification-rule"
)

var allowedMonitorNotificationRuleResourceTypeEnumValues = []MonitorNotificationRuleResourceType{
	MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorNotificationRuleResourceType) GetAllowedValues() []MonitorNotificationRuleResourceType {
	return allowedMonitorNotificationRuleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorNotificationRuleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorNotificationRuleResourceType(value)
	return nil
}

// NewMonitorNotificationRuleResourceTypeFromValue returns a pointer to a valid MonitorNotificationRuleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorNotificationRuleResourceTypeFromValue(v string) (*MonitorNotificationRuleResourceType, error) {
	ev := MonitorNotificationRuleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorNotificationRuleResourceType: valid values are %v", v, allowedMonitorNotificationRuleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorNotificationRuleResourceType) IsValid() bool {
	for _, existing := range allowedMonitorNotificationRuleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorNotificationRuleResourceType value.
func (v MonitorNotificationRuleResourceType) Ptr() *MonitorNotificationRuleResourceType {
	return &v
}
