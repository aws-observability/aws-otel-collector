// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleBasedViewType The type of the resource. The value should always be `rule_based_view`.
type RuleBasedViewType string

// List of RuleBasedViewType.
const (
	RULEBASEDVIEWTYPE_RULE_BASED_VIEW RuleBasedViewType = "rule_based_view"
)

var allowedRuleBasedViewTypeEnumValues = []RuleBasedViewType{
	RULEBASEDVIEWTYPE_RULE_BASED_VIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RuleBasedViewType) GetAllowedValues() []RuleBasedViewType {
	return allowedRuleBasedViewTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RuleBasedViewType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RuleBasedViewType(value)
	return nil
}

// NewRuleBasedViewTypeFromValue returns a pointer to a valid RuleBasedViewType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRuleBasedViewTypeFromValue(v string) (*RuleBasedViewType, error) {
	ev := RuleBasedViewType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RuleBasedViewType: valid values are %v", v, allowedRuleBasedViewTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RuleBasedViewType) IsValid() bool {
	for _, existing := range allowedRuleBasedViewTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleBasedViewType value.
func (v RuleBasedViewType) Ptr() *RuleBasedViewType {
	return &v
}
