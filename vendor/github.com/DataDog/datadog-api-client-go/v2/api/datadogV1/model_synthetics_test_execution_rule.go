// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestExecutionRule Execution rule for a Synthetic test.
type SyntheticsTestExecutionRule string

// List of SyntheticsTestExecutionRule.
const (
	SYNTHETICSTESTEXECUTIONRULE_BLOCKING     SyntheticsTestExecutionRule = "blocking"
	SYNTHETICSTESTEXECUTIONRULE_NON_BLOCKING SyntheticsTestExecutionRule = "non_blocking"
	SYNTHETICSTESTEXECUTIONRULE_SKIPPED      SyntheticsTestExecutionRule = "skipped"
)

var allowedSyntheticsTestExecutionRuleEnumValues = []SyntheticsTestExecutionRule{
	SYNTHETICSTESTEXECUTIONRULE_BLOCKING,
	SYNTHETICSTESTEXECUTIONRULE_NON_BLOCKING,
	SYNTHETICSTESTEXECUTIONRULE_SKIPPED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestExecutionRule) GetAllowedValues() []SyntheticsTestExecutionRule {
	return allowedSyntheticsTestExecutionRuleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestExecutionRule) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestExecutionRule(value)
	return nil
}

// NewSyntheticsTestExecutionRuleFromValue returns a pointer to a valid SyntheticsTestExecutionRule
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestExecutionRuleFromValue(v string) (*SyntheticsTestExecutionRule, error) {
	ev := SyntheticsTestExecutionRule(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestExecutionRule: valid values are %v", v, allowedSyntheticsTestExecutionRuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestExecutionRule) IsValid() bool {
	for _, existing := range allowedSyntheticsTestExecutionRuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestExecutionRule value.
func (v SyntheticsTestExecutionRule) Ptr() *SyntheticsTestExecutionRule {
	return &v
}
