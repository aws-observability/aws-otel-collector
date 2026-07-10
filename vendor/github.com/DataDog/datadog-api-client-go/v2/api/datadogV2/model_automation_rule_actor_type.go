// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleActorType Whether the actor is a user or the Datadog system.
type AutomationRuleActorType string

// List of AutomationRuleActorType.
const (
	AUTOMATIONRULEACTORTYPE_USER   AutomationRuleActorType = "user"
	AUTOMATIONRULEACTORTYPE_SYSTEM AutomationRuleActorType = "system"
)

var allowedAutomationRuleActorTypeEnumValues = []AutomationRuleActorType{
	AUTOMATIONRULEACTORTYPE_USER,
	AUTOMATIONRULEACTORTYPE_SYSTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AutomationRuleActorType) GetAllowedValues() []AutomationRuleActorType {
	return allowedAutomationRuleActorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutomationRuleActorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutomationRuleActorType(value)
	return nil
}

// NewAutomationRuleActorTypeFromValue returns a pointer to a valid AutomationRuleActorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutomationRuleActorTypeFromValue(v string) (*AutomationRuleActorType, error) {
	ev := AutomationRuleActorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutomationRuleActorType: valid values are %v", v, allowedAutomationRuleActorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutomationRuleActorType) IsValid() bool {
	for _, existing := range allowedAutomationRuleActorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutomationRuleActorType value.
func (v AutomationRuleActorType) Ptr() *AutomationRuleActorType {
	return &v
}
