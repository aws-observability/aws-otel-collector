// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueInteractionsType Resource type for annotation queue interactions.
type LLMObsAnnotationQueueInteractionsType string

// List of LLMObsAnnotationQueueInteractionsType.
const (
	LLMOBSANNOTATIONQUEUEINTERACTIONSTYPE_INTERACTIONS LLMObsAnnotationQueueInteractionsType = "interactions"
)

var allowedLLMObsAnnotationQueueInteractionsTypeEnumValues = []LLMObsAnnotationQueueInteractionsType{
	LLMOBSANNOTATIONQUEUEINTERACTIONSTYPE_INTERACTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnnotationQueueInteractionsType) GetAllowedValues() []LLMObsAnnotationQueueInteractionsType {
	return allowedLLMObsAnnotationQueueInteractionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnnotationQueueInteractionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnnotationQueueInteractionsType(value)
	return nil
}

// NewLLMObsAnnotationQueueInteractionsTypeFromValue returns a pointer to a valid LLMObsAnnotationQueueInteractionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnnotationQueueInteractionsTypeFromValue(v string) (*LLMObsAnnotationQueueInteractionsType, error) {
	ev := LLMObsAnnotationQueueInteractionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnnotationQueueInteractionsType: valid values are %v", v, allowedLLMObsAnnotationQueueInteractionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnnotationQueueInteractionsType) IsValid() bool {
	for _, existing := range allowedLLMObsAnnotationQueueInteractionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnnotationQueueInteractionsType value.
func (v LLMObsAnnotationQueueInteractionsType) Ptr() *LLMObsAnnotationQueueInteractionsType {
	return &v
}
