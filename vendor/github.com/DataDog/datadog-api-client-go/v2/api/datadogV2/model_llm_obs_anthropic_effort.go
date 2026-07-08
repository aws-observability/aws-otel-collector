// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnthropicEffort The effort level for Anthropic inference.
type LLMObsAnthropicEffort string

// List of LLMObsAnthropicEffort.
const (
	LLMOBSANTHROPICEFFORT_LOW    LLMObsAnthropicEffort = "low"
	LLMOBSANTHROPICEFFORT_MEDIUM LLMObsAnthropicEffort = "medium"
	LLMOBSANTHROPICEFFORT_HIGH   LLMObsAnthropicEffort = "high"
	LLMOBSANTHROPICEFFORT_MAX    LLMObsAnthropicEffort = "max"
)

var allowedLLMObsAnthropicEffortEnumValues = []LLMObsAnthropicEffort{
	LLMOBSANTHROPICEFFORT_LOW,
	LLMOBSANTHROPICEFFORT_MEDIUM,
	LLMOBSANTHROPICEFFORT_HIGH,
	LLMOBSANTHROPICEFFORT_MAX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnthropicEffort) GetAllowedValues() []LLMObsAnthropicEffort {
	return allowedLLMObsAnthropicEffortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnthropicEffort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnthropicEffort(value)
	return nil
}

// NewLLMObsAnthropicEffortFromValue returns a pointer to a valid LLMObsAnthropicEffort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnthropicEffortFromValue(v string) (*LLMObsAnthropicEffort, error) {
	ev := LLMObsAnthropicEffort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnthropicEffort: valid values are %v", v, allowedLLMObsAnthropicEffortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnthropicEffort) IsValid() bool {
	for _, existing := range allowedLLMObsAnthropicEffortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnthropicEffort value.
func (v LLMObsAnthropicEffort) Ptr() *LLMObsAnthropicEffort {
	return &v
}

// NullableLLMObsAnthropicEffort handles when a null is used for LLMObsAnthropicEffort.
type NullableLLMObsAnthropicEffort struct {
	value *LLMObsAnthropicEffort
	isSet bool
}

// Get returns the associated value.
func (v NullableLLMObsAnthropicEffort) Get() *LLMObsAnthropicEffort {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableLLMObsAnthropicEffort) Set(val *LLMObsAnthropicEffort) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableLLMObsAnthropicEffort) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableLLMObsAnthropicEffort) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLLMObsAnthropicEffort initializes the struct as if Set has been called.
func NewNullableLLMObsAnthropicEffort(val *LLMObsAnthropicEffort) *NullableLLMObsAnthropicEffort {
	return &NullableLLMObsAnthropicEffort{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableLLMObsAnthropicEffort) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableLLMObsAnthropicEffort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
