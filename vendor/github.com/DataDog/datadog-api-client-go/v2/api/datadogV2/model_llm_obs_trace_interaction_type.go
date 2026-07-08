// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsTraceInteractionType Type of an upstream-entity interaction.
type LLMObsTraceInteractionType string

// List of LLMObsTraceInteractionType.
const (
	LLMOBSTRACEINTERACTIONTYPE_TRACE            LLMObsTraceInteractionType = "trace"
	LLMOBSTRACEINTERACTIONTYPE_EXPERIMENT_TRACE LLMObsTraceInteractionType = "experiment_trace"
	LLMOBSTRACEINTERACTIONTYPE_SESSION          LLMObsTraceInteractionType = "session"
)

var allowedLLMObsTraceInteractionTypeEnumValues = []LLMObsTraceInteractionType{
	LLMOBSTRACEINTERACTIONTYPE_TRACE,
	LLMOBSTRACEINTERACTIONTYPE_EXPERIMENT_TRACE,
	LLMOBSTRACEINTERACTIONTYPE_SESSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsTraceInteractionType) GetAllowedValues() []LLMObsTraceInteractionType {
	return allowedLLMObsTraceInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsTraceInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsTraceInteractionType(value)
	return nil
}

// NewLLMObsTraceInteractionTypeFromValue returns a pointer to a valid LLMObsTraceInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsTraceInteractionTypeFromValue(v string) (*LLMObsTraceInteractionType, error) {
	ev := LLMObsTraceInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsTraceInteractionType: valid values are %v", v, allowedLLMObsTraceInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsTraceInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsTraceInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsTraceInteractionType value.
func (v LLMObsTraceInteractionType) Ptr() *LLMObsTraceInteractionType {
	return &v
}
