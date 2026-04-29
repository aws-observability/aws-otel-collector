// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionParametersType The type of the value to compare against. Only used with the equals and !equals operator.
type ApplicationSecurityWafCustomRuleConditionParametersType string

// List of ApplicationSecurityWafCustomRuleConditionParametersType.
const (
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_BOOLEAN  ApplicationSecurityWafCustomRuleConditionParametersType = "boolean"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_SIGNED   ApplicationSecurityWafCustomRuleConditionParametersType = "signed"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_UNSIGNED ApplicationSecurityWafCustomRuleConditionParametersType = "unsigned"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_FLOAT    ApplicationSecurityWafCustomRuleConditionParametersType = "float"
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_STRING   ApplicationSecurityWafCustomRuleConditionParametersType = "string"
)

var allowedApplicationSecurityWafCustomRuleConditionParametersTypeEnumValues = []ApplicationSecurityWafCustomRuleConditionParametersType{
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_BOOLEAN,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_SIGNED,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_UNSIGNED,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_FLOAT,
	APPLICATIONSECURITYWAFCUSTOMRULECONDITIONPARAMETERSTYPE_STRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafCustomRuleConditionParametersType) GetAllowedValues() []ApplicationSecurityWafCustomRuleConditionParametersType {
	return allowedApplicationSecurityWafCustomRuleConditionParametersTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafCustomRuleConditionParametersType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafCustomRuleConditionParametersType(value)
	return nil
}

// NewApplicationSecurityWafCustomRuleConditionParametersTypeFromValue returns a pointer to a valid ApplicationSecurityWafCustomRuleConditionParametersType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafCustomRuleConditionParametersTypeFromValue(v string) (*ApplicationSecurityWafCustomRuleConditionParametersType, error) {
	ev := ApplicationSecurityWafCustomRuleConditionParametersType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafCustomRuleConditionParametersType: valid values are %v", v, allowedApplicationSecurityWafCustomRuleConditionParametersTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafCustomRuleConditionParametersType) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafCustomRuleConditionParametersTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafCustomRuleConditionParametersType value.
func (v ApplicationSecurityWafCustomRuleConditionParametersType) Ptr() *ApplicationSecurityWafCustomRuleConditionParametersType {
	return &v
}
