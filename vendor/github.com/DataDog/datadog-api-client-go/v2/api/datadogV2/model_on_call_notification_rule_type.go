// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallNotificationRuleType Indicates that the resource is of type 'notification_rules'.
type OnCallNotificationRuleType string

// List of OnCallNotificationRuleType.
const (
	ONCALLNOTIFICATIONRULETYPE_NOTIFICATION_RULES OnCallNotificationRuleType = "notification_rules"
)

var allowedOnCallNotificationRuleTypeEnumValues = []OnCallNotificationRuleType{
	ONCALLNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallNotificationRuleType) GetAllowedValues() []OnCallNotificationRuleType {
	return allowedOnCallNotificationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallNotificationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallNotificationRuleType(value)
	return nil
}

// NewOnCallNotificationRuleTypeFromValue returns a pointer to a valid OnCallNotificationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallNotificationRuleTypeFromValue(v string) (*OnCallNotificationRuleType, error) {
	ev := OnCallNotificationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallNotificationRuleType: valid values are %v", v, allowedOnCallNotificationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallNotificationRuleType) IsValid() bool {
	for _, existing := range allowedOnCallNotificationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallNotificationRuleType value.
func (v OnCallNotificationRuleType) Ptr() *OnCallNotificationRuleType {
	return &v
}
