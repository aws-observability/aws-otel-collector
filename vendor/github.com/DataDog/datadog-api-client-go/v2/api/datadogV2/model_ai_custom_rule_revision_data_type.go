// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleRevisionDataType AI custom rule revision resource type.
type AiCustomRuleRevisionDataType string

// List of AiCustomRuleRevisionDataType.
const (
	AICUSTOMRULEREVISIONDATATYPE_AI_RULE_REVISION AiCustomRuleRevisionDataType = "ai_rule_revision"
)

var allowedAiCustomRuleRevisionDataTypeEnumValues = []AiCustomRuleRevisionDataType{
	AICUSTOMRULEREVISIONDATATYPE_AI_RULE_REVISION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiCustomRuleRevisionDataType) GetAllowedValues() []AiCustomRuleRevisionDataType {
	return allowedAiCustomRuleRevisionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiCustomRuleRevisionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiCustomRuleRevisionDataType(value)
	return nil
}

// NewAiCustomRuleRevisionDataTypeFromValue returns a pointer to a valid AiCustomRuleRevisionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiCustomRuleRevisionDataTypeFromValue(v string) (*AiCustomRuleRevisionDataType, error) {
	ev := AiCustomRuleRevisionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiCustomRuleRevisionDataType: valid values are %v", v, allowedAiCustomRuleRevisionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiCustomRuleRevisionDataType) IsValid() bool {
	for _, existing := range allowedAiCustomRuleRevisionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiCustomRuleRevisionDataType value.
func (v AiCustomRuleRevisionDataType) Ptr() *AiCustomRuleRevisionDataType {
	return &v
}
