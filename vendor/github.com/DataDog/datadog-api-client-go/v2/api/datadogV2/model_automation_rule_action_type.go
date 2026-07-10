// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleActionType The type of automated action to perform when the rule triggers. `execute_workflow` runs a Datadog workflow; `assign_agent` assigns an AI agent to the case.
type AutomationRuleActionType string

// List of AutomationRuleActionType.
const (
	AUTOMATIONRULEACTIONTYPE_EXECUTE_WORKFLOW AutomationRuleActionType = "execute_workflow"
	AUTOMATIONRULEACTIONTYPE_ASSIGN_AGENT     AutomationRuleActionType = "assign_agent"
)

var allowedAutomationRuleActionTypeEnumValues = []AutomationRuleActionType{
	AUTOMATIONRULEACTIONTYPE_EXECUTE_WORKFLOW,
	AUTOMATIONRULEACTIONTYPE_ASSIGN_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AutomationRuleActionType) GetAllowedValues() []AutomationRuleActionType {
	return allowedAutomationRuleActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutomationRuleActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutomationRuleActionType(value)
	return nil
}

// NewAutomationRuleActionTypeFromValue returns a pointer to a valid AutomationRuleActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutomationRuleActionTypeFromValue(v string) (*AutomationRuleActionType, error) {
	ev := AutomationRuleActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutomationRuleActionType: valid values are %v", v, allowedAutomationRuleActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutomationRuleActionType) IsValid() bool {
	for _, existing := range allowedAutomationRuleActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutomationRuleActionType value.
func (v AutomationRuleActionType) Ptr() *AutomationRuleActionType {
	return &v
}
