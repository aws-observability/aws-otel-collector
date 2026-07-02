// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnthropicThinkingType The thinking mode for Anthropic extended thinking.
type LLMObsAnthropicThinkingType string

// List of LLMObsAnthropicThinkingType.
const (
	LLMOBSANTHROPICTHINKINGTYPE_ENABLED  LLMObsAnthropicThinkingType = "enabled"
	LLMOBSANTHROPICTHINKINGTYPE_DISABLED LLMObsAnthropicThinkingType = "disabled"
	LLMOBSANTHROPICTHINKINGTYPE_ADAPTIVE LLMObsAnthropicThinkingType = "adaptive"
)

var allowedLLMObsAnthropicThinkingTypeEnumValues = []LLMObsAnthropicThinkingType{
	LLMOBSANTHROPICTHINKINGTYPE_ENABLED,
	LLMOBSANTHROPICTHINKINGTYPE_DISABLED,
	LLMOBSANTHROPICTHINKINGTYPE_ADAPTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnthropicThinkingType) GetAllowedValues() []LLMObsAnthropicThinkingType {
	return allowedLLMObsAnthropicThinkingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnthropicThinkingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnthropicThinkingType(value)
	return nil
}

// NewLLMObsAnthropicThinkingTypeFromValue returns a pointer to a valid LLMObsAnthropicThinkingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnthropicThinkingTypeFromValue(v string) (*LLMObsAnthropicThinkingType, error) {
	ev := LLMObsAnthropicThinkingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnthropicThinkingType: valid values are %v", v, allowedLLMObsAnthropicThinkingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnthropicThinkingType) IsValid() bool {
	for _, existing := range allowedLLMObsAnthropicThinkingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnthropicThinkingType value.
func (v LLMObsAnthropicThinkingType) Ptr() *LLMObsAnthropicThinkingType {
	return &v
}
