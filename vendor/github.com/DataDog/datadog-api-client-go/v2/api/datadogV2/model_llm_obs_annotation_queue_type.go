// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueType Resource type of an LLM Observability annotation queue.
type LLMObsAnnotationQueueType string

// List of LLMObsAnnotationQueueType.
const (
	LLMOBSANNOTATIONQUEUETYPE_QUEUES LLMObsAnnotationQueueType = "queues"
)

var allowedLLMObsAnnotationQueueTypeEnumValues = []LLMObsAnnotationQueueType{
	LLMOBSANNOTATIONQUEUETYPE_QUEUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnnotationQueueType) GetAllowedValues() []LLMObsAnnotationQueueType {
	return allowedLLMObsAnnotationQueueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnnotationQueueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnnotationQueueType(value)
	return nil
}

// NewLLMObsAnnotationQueueTypeFromValue returns a pointer to a valid LLMObsAnnotationQueueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnnotationQueueTypeFromValue(v string) (*LLMObsAnnotationQueueType, error) {
	ev := LLMObsAnnotationQueueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnnotationQueueType: valid values are %v", v, allowedLLMObsAnnotationQueueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnnotationQueueType) IsValid() bool {
	for _, existing := range allowedLLMObsAnnotationQueueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnnotationQueueType value.
func (v LLMObsAnnotationQueueType) Ptr() *LLMObsAnnotationQueueType {
	return &v
}
