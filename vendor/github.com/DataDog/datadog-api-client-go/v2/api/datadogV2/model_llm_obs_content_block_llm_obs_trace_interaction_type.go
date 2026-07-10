// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsContentBlockLLMObsTraceInteractionType Upstream interaction type referenced by an `llmobs_trace` block.
// Restricted to `trace` or `experiment_trace`.
type LLMObsContentBlockLLMObsTraceInteractionType string

// List of LLMObsContentBlockLLMObsTraceInteractionType.
const (
	LLMOBSCONTENTBLOCKLLMOBSTRACEINTERACTIONTYPE_TRACE            LLMObsContentBlockLLMObsTraceInteractionType = "trace"
	LLMOBSCONTENTBLOCKLLMOBSTRACEINTERACTIONTYPE_EXPERIMENT_TRACE LLMObsContentBlockLLMObsTraceInteractionType = "experiment_trace"
)

var allowedLLMObsContentBlockLLMObsTraceInteractionTypeEnumValues = []LLMObsContentBlockLLMObsTraceInteractionType{
	LLMOBSCONTENTBLOCKLLMOBSTRACEINTERACTIONTYPE_TRACE,
	LLMOBSCONTENTBLOCKLLMOBSTRACEINTERACTIONTYPE_EXPERIMENT_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsContentBlockLLMObsTraceInteractionType) GetAllowedValues() []LLMObsContentBlockLLMObsTraceInteractionType {
	return allowedLLMObsContentBlockLLMObsTraceInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsContentBlockLLMObsTraceInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsContentBlockLLMObsTraceInteractionType(value)
	return nil
}

// NewLLMObsContentBlockLLMObsTraceInteractionTypeFromValue returns a pointer to a valid LLMObsContentBlockLLMObsTraceInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsContentBlockLLMObsTraceInteractionTypeFromValue(v string) (*LLMObsContentBlockLLMObsTraceInteractionType, error) {
	ev := LLMObsContentBlockLLMObsTraceInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsContentBlockLLMObsTraceInteractionType: valid values are %v", v, allowedLLMObsContentBlockLLMObsTraceInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsContentBlockLLMObsTraceInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsContentBlockLLMObsTraceInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsContentBlockLLMObsTraceInteractionType value.
func (v LLMObsContentBlockLLMObsTraceInteractionType) Ptr() *LLMObsContentBlockLLMObsTraceInteractionType {
	return &v
}
