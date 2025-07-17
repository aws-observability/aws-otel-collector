// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleTagsCategory The category of the WAF Rule, can be either `business_logic`, `attack_attempt` or `security_response`.
type ApplicationSecurityWafCustomRuleTagsCategory string

// List of ApplicationSecurityWafCustomRuleTagsCategory.
const (
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_ATTACK_ATTEMPT    ApplicationSecurityWafCustomRuleTagsCategory = "attack_attempt"
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_BUSINESS_LOGIC    ApplicationSecurityWafCustomRuleTagsCategory = "business_logic"
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_SECURITY_RESPONSE ApplicationSecurityWafCustomRuleTagsCategory = "security_response"
)

var allowedApplicationSecurityWafCustomRuleTagsCategoryEnumValues = []ApplicationSecurityWafCustomRuleTagsCategory{
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_ATTACK_ATTEMPT,
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_BUSINESS_LOGIC,
	APPLICATIONSECURITYWAFCUSTOMRULETAGSCATEGORY_SECURITY_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafCustomRuleTagsCategory) GetAllowedValues() []ApplicationSecurityWafCustomRuleTagsCategory {
	return allowedApplicationSecurityWafCustomRuleTagsCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafCustomRuleTagsCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafCustomRuleTagsCategory(value)
	return nil
}

// NewApplicationSecurityWafCustomRuleTagsCategoryFromValue returns a pointer to a valid ApplicationSecurityWafCustomRuleTagsCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafCustomRuleTagsCategoryFromValue(v string) (*ApplicationSecurityWafCustomRuleTagsCategory, error) {
	ev := ApplicationSecurityWafCustomRuleTagsCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafCustomRuleTagsCategory: valid values are %v", v, allowedApplicationSecurityWafCustomRuleTagsCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafCustomRuleTagsCategory) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafCustomRuleTagsCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafCustomRuleTagsCategory value.
func (v ApplicationSecurityWafCustomRuleTagsCategory) Ptr() *ApplicationSecurityWafCustomRuleTagsCategory {
	return &v
}
