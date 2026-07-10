// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnyInteractionType Type of an annotated interaction.
type LLMObsAnyInteractionType string

// List of LLMObsAnyInteractionType.
const (
	LLMOBSANYINTERACTIONTYPE_TRACE            LLMObsAnyInteractionType = "trace"
	LLMOBSANYINTERACTIONTYPE_EXPERIMENT_TRACE LLMObsAnyInteractionType = "experiment_trace"
	LLMOBSANYINTERACTIONTYPE_SESSION          LLMObsAnyInteractionType = "session"
	LLMOBSANYINTERACTIONTYPE_DISPLAY_BLOCK    LLMObsAnyInteractionType = "display_block"
)

var allowedLLMObsAnyInteractionTypeEnumValues = []LLMObsAnyInteractionType{
	LLMOBSANYINTERACTIONTYPE_TRACE,
	LLMOBSANYINTERACTIONTYPE_EXPERIMENT_TRACE,
	LLMOBSANYINTERACTIONTYPE_SESSION,
	LLMOBSANYINTERACTIONTYPE_DISPLAY_BLOCK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnyInteractionType) GetAllowedValues() []LLMObsAnyInteractionType {
	return allowedLLMObsAnyInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnyInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnyInteractionType(value)
	return nil
}

// NewLLMObsAnyInteractionTypeFromValue returns a pointer to a valid LLMObsAnyInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnyInteractionTypeFromValue(v string) (*LLMObsAnyInteractionType, error) {
	ev := LLMObsAnyInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnyInteractionType: valid values are %v", v, allowedLLMObsAnyInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnyInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsAnyInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnyInteractionType value.
func (v LLMObsAnyInteractionType) Ptr() *LLMObsAnyInteractionType {
	return &v
}
