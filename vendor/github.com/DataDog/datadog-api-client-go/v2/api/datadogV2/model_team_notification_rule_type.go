// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRuleType Team notification rule type
type TeamNotificationRuleType string

// List of TeamNotificationRuleType.
const (
	TEAMNOTIFICATIONRULETYPE_TEAM_NOTIFICATION_RULES TeamNotificationRuleType = "team_notification_rules"
)

var allowedTeamNotificationRuleTypeEnumValues = []TeamNotificationRuleType{
	TEAMNOTIFICATIONRULETYPE_TEAM_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamNotificationRuleType) GetAllowedValues() []TeamNotificationRuleType {
	return allowedTeamNotificationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamNotificationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamNotificationRuleType(value)
	return nil
}

// NewTeamNotificationRuleTypeFromValue returns a pointer to a valid TeamNotificationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamNotificationRuleTypeFromValue(v string) (*TeamNotificationRuleType, error) {
	ev := TeamNotificationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamNotificationRuleType: valid values are %v", v, allowedTeamNotificationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamNotificationRuleType) IsValid() bool {
	for _, existing := range allowedTeamNotificationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamNotificationRuleType value.
func (v TeamNotificationRuleType) Ptr() *TeamNotificationRuleType {
	return &v
}
