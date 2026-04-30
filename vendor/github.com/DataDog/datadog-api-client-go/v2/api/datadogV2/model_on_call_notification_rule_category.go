// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallNotificationRuleCategory Specifies the category a notification rule will apply to
type OnCallNotificationRuleCategory string

// List of OnCallNotificationRuleCategory.
const (
	ONCALLNOTIFICATIONRULECATEGORY_HIGH_URGENCY OnCallNotificationRuleCategory = "high_urgency"
	ONCALLNOTIFICATIONRULECATEGORY_LOW_URGENCY  OnCallNotificationRuleCategory = "low_urgency"
)

var allowedOnCallNotificationRuleCategoryEnumValues = []OnCallNotificationRuleCategory{
	ONCALLNOTIFICATIONRULECATEGORY_HIGH_URGENCY,
	ONCALLNOTIFICATIONRULECATEGORY_LOW_URGENCY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallNotificationRuleCategory) GetAllowedValues() []OnCallNotificationRuleCategory {
	return allowedOnCallNotificationRuleCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallNotificationRuleCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallNotificationRuleCategory(value)
	return nil
}

// NewOnCallNotificationRuleCategoryFromValue returns a pointer to a valid OnCallNotificationRuleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallNotificationRuleCategoryFromValue(v string) (*OnCallNotificationRuleCategory, error) {
	ev := OnCallNotificationRuleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallNotificationRuleCategory: valid values are %v", v, allowedOnCallNotificationRuleCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallNotificationRuleCategory) IsValid() bool {
	for _, existing := range allowedOnCallNotificationRuleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallNotificationRuleCategory value.
func (v OnCallNotificationRuleCategory) Ptr() *OnCallNotificationRuleCategory {
	return &v
}
