// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsOpenAIReasoningSummary The verbosity of the reasoning summary.
type LLMObsOpenAIReasoningSummary string

// List of LLMObsOpenAIReasoningSummary.
const (
	LLMOBSOPENAIREASONINGSUMMARY_AUTO     LLMObsOpenAIReasoningSummary = "auto"
	LLMOBSOPENAIREASONINGSUMMARY_CONCISE  LLMObsOpenAIReasoningSummary = "concise"
	LLMOBSOPENAIREASONINGSUMMARY_DETAILED LLMObsOpenAIReasoningSummary = "detailed"
)

var allowedLLMObsOpenAIReasoningSummaryEnumValues = []LLMObsOpenAIReasoningSummary{
	LLMOBSOPENAIREASONINGSUMMARY_AUTO,
	LLMOBSOPENAIREASONINGSUMMARY_CONCISE,
	LLMOBSOPENAIREASONINGSUMMARY_DETAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsOpenAIReasoningSummary) GetAllowedValues() []LLMObsOpenAIReasoningSummary {
	return allowedLLMObsOpenAIReasoningSummaryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsOpenAIReasoningSummary) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsOpenAIReasoningSummary(value)
	return nil
}

// NewLLMObsOpenAIReasoningSummaryFromValue returns a pointer to a valid LLMObsOpenAIReasoningSummary
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsOpenAIReasoningSummaryFromValue(v string) (*LLMObsOpenAIReasoningSummary, error) {
	ev := LLMObsOpenAIReasoningSummary(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsOpenAIReasoningSummary: valid values are %v", v, allowedLLMObsOpenAIReasoningSummaryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsOpenAIReasoningSummary) IsValid() bool {
	for _, existing := range allowedLLMObsOpenAIReasoningSummaryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsOpenAIReasoningSummary value.
func (v LLMObsOpenAIReasoningSummary) Ptr() *LLMObsOpenAIReasoningSummary {
	return &v
}

// NullableLLMObsOpenAIReasoningSummary handles when a null is used for LLMObsOpenAIReasoningSummary.
type NullableLLMObsOpenAIReasoningSummary struct {
	value *LLMObsOpenAIReasoningSummary
	isSet bool
}

// Get returns the associated value.
func (v NullableLLMObsOpenAIReasoningSummary) Get() *LLMObsOpenAIReasoningSummary {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableLLMObsOpenAIReasoningSummary) Set(val *LLMObsOpenAIReasoningSummary) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableLLMObsOpenAIReasoningSummary) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableLLMObsOpenAIReasoningSummary) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLLMObsOpenAIReasoningSummary initializes the struct as if Set has been called.
func NewNullableLLMObsOpenAIReasoningSummary(val *LLMObsOpenAIReasoningSummary) *NullableLLMObsOpenAIReasoningSummary {
	return &NullableLLMObsOpenAIReasoningSummary{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableLLMObsOpenAIReasoningSummary) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableLLMObsOpenAIReasoningSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
