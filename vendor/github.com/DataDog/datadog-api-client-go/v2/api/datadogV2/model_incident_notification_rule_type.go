// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleType Notification rules resource type.
type IncidentNotificationRuleType string

// List of IncidentNotificationRuleType.
const (
	INCIDENTNOTIFICATIONRULETYPE_INCIDENT_NOTIFICATION_RULES IncidentNotificationRuleType = "incident_notification_rules"
)

var allowedIncidentNotificationRuleTypeEnumValues = []IncidentNotificationRuleType{
	INCIDENTNOTIFICATIONRULETYPE_INCIDENT_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentNotificationRuleType) GetAllowedValues() []IncidentNotificationRuleType {
	return allowedIncidentNotificationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentNotificationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentNotificationRuleType(value)
	return nil
}

// NewIncidentNotificationRuleTypeFromValue returns a pointer to a valid IncidentNotificationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentNotificationRuleTypeFromValue(v string) (*IncidentNotificationRuleType, error) {
	ev := IncidentNotificationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentNotificationRuleType: valid values are %v", v, allowedIncidentNotificationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentNotificationRuleType) IsValid() bool {
	for _, existing := range allowedIncidentNotificationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentNotificationRuleType value.
func (v IncidentNotificationRuleType) Ptr() *IncidentNotificationRuleType {
	return &v
}
