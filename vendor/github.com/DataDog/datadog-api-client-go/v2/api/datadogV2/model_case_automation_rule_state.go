// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAutomationRuleState Whether the automation rule is active. Enabled rules trigger on matching case events; disabled rules are inactive but preserve their configuration.
type CaseAutomationRuleState string

// List of CaseAutomationRuleState.
const (
	CASEAUTOMATIONRULESTATE_ENABLED  CaseAutomationRuleState = "ENABLED"
	CASEAUTOMATIONRULESTATE_DISABLED CaseAutomationRuleState = "DISABLED"
)

var allowedCaseAutomationRuleStateEnumValues = []CaseAutomationRuleState{
	CASEAUTOMATIONRULESTATE_ENABLED,
	CASEAUTOMATIONRULESTATE_DISABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseAutomationRuleState) GetAllowedValues() []CaseAutomationRuleState {
	return allowedCaseAutomationRuleStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseAutomationRuleState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseAutomationRuleState(value)
	return nil
}

// NewCaseAutomationRuleStateFromValue returns a pointer to a valid CaseAutomationRuleState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseAutomationRuleStateFromValue(v string) (*CaseAutomationRuleState, error) {
	ev := CaseAutomationRuleState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseAutomationRuleState: valid values are %v", v, allowedCaseAutomationRuleStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseAutomationRuleState) IsValid() bool {
	for _, existing := range allowedCaseAutomationRuleStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseAutomationRuleState value.
func (v CaseAutomationRuleState) Ptr() *CaseAutomationRuleState {
	return &v
}
