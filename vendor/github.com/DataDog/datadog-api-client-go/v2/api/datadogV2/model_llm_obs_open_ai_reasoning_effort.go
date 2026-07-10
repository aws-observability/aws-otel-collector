// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsOpenAIReasoningEffort The reasoning effort level for OpenAI models that support it.
type LLMObsOpenAIReasoningEffort string

// List of LLMObsOpenAIReasoningEffort.
const (
	LLMOBSOPENAIREASONINGEFFORT_NONE   LLMObsOpenAIReasoningEffort = "none"
	LLMOBSOPENAIREASONINGEFFORT_LOW    LLMObsOpenAIReasoningEffort = "low"
	LLMOBSOPENAIREASONINGEFFORT_MEDIUM LLMObsOpenAIReasoningEffort = "medium"
	LLMOBSOPENAIREASONINGEFFORT_HIGH   LLMObsOpenAIReasoningEffort = "high"
	LLMOBSOPENAIREASONINGEFFORT_XHIGH  LLMObsOpenAIReasoningEffort = "xhigh"
)

var allowedLLMObsOpenAIReasoningEffortEnumValues = []LLMObsOpenAIReasoningEffort{
	LLMOBSOPENAIREASONINGEFFORT_NONE,
	LLMOBSOPENAIREASONINGEFFORT_LOW,
	LLMOBSOPENAIREASONINGEFFORT_MEDIUM,
	LLMOBSOPENAIREASONINGEFFORT_HIGH,
	LLMOBSOPENAIREASONINGEFFORT_XHIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsOpenAIReasoningEffort) GetAllowedValues() []LLMObsOpenAIReasoningEffort {
	return allowedLLMObsOpenAIReasoningEffortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsOpenAIReasoningEffort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsOpenAIReasoningEffort(value)
	return nil
}

// NewLLMObsOpenAIReasoningEffortFromValue returns a pointer to a valid LLMObsOpenAIReasoningEffort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsOpenAIReasoningEffortFromValue(v string) (*LLMObsOpenAIReasoningEffort, error) {
	ev := LLMObsOpenAIReasoningEffort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsOpenAIReasoningEffort: valid values are %v", v, allowedLLMObsOpenAIReasoningEffortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsOpenAIReasoningEffort) IsValid() bool {
	for _, existing := range allowedLLMObsOpenAIReasoningEffortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsOpenAIReasoningEffort value.
func (v LLMObsOpenAIReasoningEffort) Ptr() *LLMObsOpenAIReasoningEffort {
	return &v
}

// NullableLLMObsOpenAIReasoningEffort handles when a null is used for LLMObsOpenAIReasoningEffort.
type NullableLLMObsOpenAIReasoningEffort struct {
	value *LLMObsOpenAIReasoningEffort
	isSet bool
}

// Get returns the associated value.
func (v NullableLLMObsOpenAIReasoningEffort) Get() *LLMObsOpenAIReasoningEffort {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableLLMObsOpenAIReasoningEffort) Set(val *LLMObsOpenAIReasoningEffort) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableLLMObsOpenAIReasoningEffort) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableLLMObsOpenAIReasoningEffort) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLLMObsOpenAIReasoningEffort initializes the struct as if Set has been called.
func NewNullableLLMObsOpenAIReasoningEffort(val *LLMObsOpenAIReasoningEffort) *NullableLLMObsOpenAIReasoningEffort {
	return &NullableLLMObsOpenAIReasoningEffort{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableLLMObsOpenAIReasoningEffort) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableLLMObsOpenAIReasoningEffort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
