// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseAutomationRuleResourceType JSON:API resource type for case automation rules.
type CaseAutomationRuleResourceType string

// List of CaseAutomationRuleResourceType.
const (
	CASEAUTOMATIONRULERESOURCETYPE_RULE CaseAutomationRuleResourceType = "rule"
)

var allowedCaseAutomationRuleResourceTypeEnumValues = []CaseAutomationRuleResourceType{
	CASEAUTOMATIONRULERESOURCETYPE_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseAutomationRuleResourceType) GetAllowedValues() []CaseAutomationRuleResourceType {
	return allowedCaseAutomationRuleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseAutomationRuleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseAutomationRuleResourceType(value)
	return nil
}

// NewCaseAutomationRuleResourceTypeFromValue returns a pointer to a valid CaseAutomationRuleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseAutomationRuleResourceTypeFromValue(v string) (*CaseAutomationRuleResourceType, error) {
	ev := CaseAutomationRuleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseAutomationRuleResourceType: valid values are %v", v, allowedCaseAutomationRuleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseAutomationRuleResourceType) IsValid() bool {
	for _, existing := range allowedCaseAutomationRuleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseAutomationRuleResourceType value.
func (v CaseAutomationRuleResourceType) Ptr() *CaseAutomationRuleResourceType {
	return &v
}
