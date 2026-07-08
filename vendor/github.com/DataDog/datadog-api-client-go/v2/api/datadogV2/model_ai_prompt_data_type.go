// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiPromptDataType AI prompt resource type.
type AiPromptDataType string

// List of AiPromptDataType.
const (
	AIPROMPTDATATYPE_AI_PROMPT AiPromptDataType = "ai_prompt"
)

var allowedAiPromptDataTypeEnumValues = []AiPromptDataType{
	AIPROMPTDATATYPE_AI_PROMPT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiPromptDataType) GetAllowedValues() []AiPromptDataType {
	return allowedAiPromptDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiPromptDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiPromptDataType(value)
	return nil
}

// NewAiPromptDataTypeFromValue returns a pointer to a valid AiPromptDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiPromptDataTypeFromValue(v string) (*AiPromptDataType, error) {
	ev := AiPromptDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiPromptDataType: valid values are %v", v, allowedAiPromptDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiPromptDataType) IsValid() bool {
	for _, existing := range allowedAiPromptDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiPromptDataType value.
func (v AiPromptDataType) Ptr() *AiPromptDataType {
	return &v
}
