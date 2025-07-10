// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionOperator Operator to use for the WAF Condition.
type ApplicationSecurityWafCustomRuleConditionOperator string

// List of ApplicationSecurityWafCustomRuleConditionOperator.
const (
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_MATCH_REGEX      ApplicationSecurityWafCustomRuleConditionOperator = "match_regex"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_MATCH_REGEX  ApplicationSecurityWafCustomRuleConditionOperator = "!match_regex"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_PHRASE_MATCH     ApplicationSecurityWafCustomRuleConditionOperator = "phrase_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_PHRASE_MATCH ApplicationSecurityWafCustomRuleConditionOperator = "!phrase_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IS_XSS           ApplicationSecurityWafCustomRuleConditionOperator = "is_xss"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IS_SQLI          ApplicationSecurityWafCustomRuleConditionOperator = "is_sqli"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_EXACT_MATCH      ApplicationSecurityWafCustomRuleConditionOperator = "exact_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_EXACT_MATCH  ApplicationSecurityWafCustomRuleConditionOperator = "!exact_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IP_MATCH         ApplicationSecurityWafCustomRuleConditionOperator = "ip_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_IP_MATCH     ApplicationSecurityWafCustomRuleConditionOperator = "!ip_match"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_CAPTURE_DATA     ApplicationSecurityWafCustomRuleConditionOperator = "capture_data"
)

var allowedApplicationSecurityWafCustomRuleConditionOperatorEnumValues = []ApplicationSecurityWafCustomRuleConditionOperator{
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_MATCH_REGEX,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_MATCH_REGEX,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_PHRASE_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_PHRASE_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IS_XSS,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IS_SQLI,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_EXACT_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_EXACT_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_IP_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_NOT_IP_MATCH,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONOPERATOR_CAPTURE_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafCustomRuleConditionOperator) GetAllowedValues() []ApplicationSecurityWafCustomRuleConditionOperator {
	return allowedApplicationSecurityWafCustomRuleConditionOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafCustomRuleConditionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafCustomRuleConditionOperator(value)
	return nil
}

// NewApplicationSecurityWafCustomRuleConditionOperatorFromValue returns a pointer to a valid ApplicationSecurityWafCustomRuleConditionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafCustomRuleConditionOperatorFromValue(v string) (*ApplicationSecurityWafCustomRuleConditionOperator, error) {
	ev := ApplicationSecurityWafCustomRuleConditionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafCustomRuleConditionOperator: valid values are %v", v, allowedApplicationSecurityWafCustomRuleConditionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafCustomRuleConditionOperator) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafCustomRuleConditionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafCustomRuleConditionOperator value.
func (v ApplicationSecurityWafCustomRuleConditionOperator) Ptr() *ApplicationSecurityWafCustomRuleConditionOperator {
	return &v
}
