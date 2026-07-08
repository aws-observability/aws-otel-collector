// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleTriggerType The case event that activates the automation rule.
type AutomationRuleTriggerType string

// List of AutomationRuleTriggerType.
const (
	AUTOMATIONRULETRIGGERTYPE_CASE_CREATED                        AutomationRuleTriggerType = "case_created"
	AUTOMATIONRULETRIGGERTYPE_STATUS_TRANSITIONED                 AutomationRuleTriggerType = "status_transitioned"
	AUTOMATIONRULETRIGGERTYPE_ATTRIBUTE_VALUE_CHANGED             AutomationRuleTriggerType = "attribute_value_changed"
	AUTOMATIONRULETRIGGERTYPE_EVENT_CORRELATION_SIGNAL_CORRELATED AutomationRuleTriggerType = "event_correlation_signal_correlated"
	AUTOMATIONRULETRIGGERTYPE_CASE_REVIEW_APPROVED                AutomationRuleTriggerType = "case_review_approved"
	AUTOMATIONRULETRIGGERTYPE_COMMENT_ADDED                       AutomationRuleTriggerType = "comment_added"
)

var allowedAutomationRuleTriggerTypeEnumValues = []AutomationRuleTriggerType{
	AUTOMATIONRULETRIGGERTYPE_CASE_CREATED,
	AUTOMATIONRULETRIGGERTYPE_STATUS_TRANSITIONED,
	AUTOMATIONRULETRIGGERTYPE_ATTRIBUTE_VALUE_CHANGED,
	AUTOMATIONRULETRIGGERTYPE_EVENT_CORRELATION_SIGNAL_CORRELATED,
	AUTOMATIONRULETRIGGERTYPE_CASE_REVIEW_APPROVED,
	AUTOMATIONRULETRIGGERTYPE_COMMENT_ADDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AutomationRuleTriggerType) GetAllowedValues() []AutomationRuleTriggerType {
	return allowedAutomationRuleTriggerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AutomationRuleTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AutomationRuleTriggerType(value)
	return nil
}

// NewAutomationRuleTriggerTypeFromValue returns a pointer to a valid AutomationRuleTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAutomationRuleTriggerTypeFromValue(v string) (*AutomationRuleTriggerType, error) {
	ev := AutomationRuleTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AutomationRuleTriggerType: valid values are %v", v, allowedAutomationRuleTriggerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AutomationRuleTriggerType) IsValid() bool {
	for _, existing := range allowedAutomationRuleTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutomationRuleTriggerType value.
func (v AutomationRuleTriggerType) Ptr() *AutomationRuleTriggerType {
	return &v
}
