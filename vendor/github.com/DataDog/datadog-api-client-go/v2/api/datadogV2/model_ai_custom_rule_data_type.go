// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleDataType AI custom rule resource type.
type AiCustomRuleDataType string

// List of AiCustomRuleDataType.
const (
	AICUSTOMRULEDATATYPE_AI_RULE AiCustomRuleDataType = "ai_rule"
)

var allowedAiCustomRuleDataTypeEnumValues = []AiCustomRuleDataType{
	AICUSTOMRULEDATATYPE_AI_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiCustomRuleDataType) GetAllowedValues() []AiCustomRuleDataType {
	return allowedAiCustomRuleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiCustomRuleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiCustomRuleDataType(value)
	return nil
}

// NewAiCustomRuleDataTypeFromValue returns a pointer to a valid AiCustomRuleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiCustomRuleDataTypeFromValue(v string) (*AiCustomRuleDataType, error) {
	ev := AiCustomRuleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiCustomRuleDataType: valid values are %v", v, allowedAiCustomRuleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiCustomRuleDataType) IsValid() bool {
	for _, existing := range allowedAiCustomRuleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiCustomRuleDataType value.
func (v AiCustomRuleDataType) Ptr() *AiCustomRuleDataType {
	return &v
}
