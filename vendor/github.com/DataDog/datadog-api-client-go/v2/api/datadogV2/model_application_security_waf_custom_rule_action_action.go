// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleActionAction Override the default action to take when the WAF custom rule would block.
type ApplicationSecurityWafCustomRuleActionAction string

// List of ApplicationSecurityWafCustomRuleActionAction.
const (
	APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_REDIRECT_REQUEST ApplicationSecurityWafCustomRuleActionAction = "redirect_request"
	APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_BLOCK_REQUEST    ApplicationSecurityWafCustomRuleActionAction = "block_request"
)

var allowedApplicationSecurityWafCustomRuleActionActionEnumValues = []ApplicationSecurityWafCustomRuleActionAction{
	APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_REDIRECT_REQUEST,
	APPLICATIONSECURITYWAFCUSTOMRULEACTIONACTION_BLOCK_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafCustomRuleActionAction) GetAllowedValues() []ApplicationSecurityWafCustomRuleActionAction {
	return allowedApplicationSecurityWafCustomRuleActionActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafCustomRuleActionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafCustomRuleActionAction(value)
	return nil
}

// NewApplicationSecurityWafCustomRuleActionActionFromValue returns a pointer to a valid ApplicationSecurityWafCustomRuleActionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafCustomRuleActionActionFromValue(v string) (*ApplicationSecurityWafCustomRuleActionAction, error) {
	ev := ApplicationSecurityWafCustomRuleActionAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafCustomRuleActionAction: valid values are %v", v, allowedApplicationSecurityWafCustomRuleActionActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafCustomRuleActionAction) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafCustomRuleActionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafCustomRuleActionAction value.
func (v ApplicationSecurityWafCustomRuleActionAction) Ptr() *ApplicationSecurityWafCustomRuleActionAction {
	return &v
}
