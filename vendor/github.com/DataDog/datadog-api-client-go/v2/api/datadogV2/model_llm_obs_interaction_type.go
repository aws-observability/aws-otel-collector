// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsInteractionType Type of interaction in an annotation queue.
type LLMObsInteractionType string

// List of LLMObsInteractionType.
const (
	LLMOBSINTERACTIONTYPE_TRACE            LLMObsInteractionType = "trace"
	LLMOBSINTERACTIONTYPE_EXPERIMENT_TRACE LLMObsInteractionType = "experiment_trace"
)

var allowedLLMObsInteractionTypeEnumValues = []LLMObsInteractionType{
	LLMOBSINTERACTIONTYPE_TRACE,
	LLMOBSINTERACTIONTYPE_EXPERIMENT_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsInteractionType) GetAllowedValues() []LLMObsInteractionType {
	return allowedLLMObsInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsInteractionType(value)
	return nil
}

// NewLLMObsInteractionTypeFromValue returns a pointer to a valid LLMObsInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsInteractionTypeFromValue(v string) (*LLMObsInteractionType, error) {
	ev := LLMObsInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsInteractionType: valid values are %v", v, allowedLLMObsInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsInteractionType value.
func (v LLMObsInteractionType) Ptr() *LLMObsInteractionType {
	return &v
}
