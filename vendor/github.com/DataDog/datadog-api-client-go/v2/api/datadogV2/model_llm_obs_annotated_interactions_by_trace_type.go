// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionsByTraceType Resource type for cross-queue annotated interactions lookup.
type LLMObsAnnotatedInteractionsByTraceType string

// List of LLMObsAnnotatedInteractionsByTraceType.
const (
	LLMOBSANNOTATEDINTERACTIONSBYTRACETYPE_ANNOTATED_INTERACTIONS_BY_TRACE LLMObsAnnotatedInteractionsByTraceType = "annotated_interactions_by_trace"
)

var allowedLLMObsAnnotatedInteractionsByTraceTypeEnumValues = []LLMObsAnnotatedInteractionsByTraceType{
	LLMOBSANNOTATEDINTERACTIONSBYTRACETYPE_ANNOTATED_INTERACTIONS_BY_TRACE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnnotatedInteractionsByTraceType) GetAllowedValues() []LLMObsAnnotatedInteractionsByTraceType {
	return allowedLLMObsAnnotatedInteractionsByTraceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnnotatedInteractionsByTraceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnnotatedInteractionsByTraceType(value)
	return nil
}

// NewLLMObsAnnotatedInteractionsByTraceTypeFromValue returns a pointer to a valid LLMObsAnnotatedInteractionsByTraceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnnotatedInteractionsByTraceTypeFromValue(v string) (*LLMObsAnnotatedInteractionsByTraceType, error) {
	ev := LLMObsAnnotatedInteractionsByTraceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnnotatedInteractionsByTraceType: valid values are %v", v, allowedLLMObsAnnotatedInteractionsByTraceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnnotatedInteractionsByTraceType) IsValid() bool {
	for _, existing := range allowedLLMObsAnnotatedInteractionsByTraceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnnotatedInteractionsByTraceType value.
func (v LLMObsAnnotatedInteractionsByTraceType) Ptr() *LLMObsAnnotatedInteractionsByTraceType {
	return &v
}
