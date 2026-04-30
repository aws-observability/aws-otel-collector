// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionsType Resource type for annotated interactions.
type LLMObsAnnotatedInteractionsType string

// List of LLMObsAnnotatedInteractionsType.
const (
	LLMOBSANNOTATEDINTERACTIONSTYPE_ANNOTATED_INTERACTIONS LLMObsAnnotatedInteractionsType = "annotated_interactions"
)

var allowedLLMObsAnnotatedInteractionsTypeEnumValues = []LLMObsAnnotatedInteractionsType{
	LLMOBSANNOTATEDINTERACTIONSTYPE_ANNOTATED_INTERACTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnnotatedInteractionsType) GetAllowedValues() []LLMObsAnnotatedInteractionsType {
	return allowedLLMObsAnnotatedInteractionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnnotatedInteractionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnnotatedInteractionsType(value)
	return nil
}

// NewLLMObsAnnotatedInteractionsTypeFromValue returns a pointer to a valid LLMObsAnnotatedInteractionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnnotatedInteractionsTypeFromValue(v string) (*LLMObsAnnotatedInteractionsType, error) {
	ev := LLMObsAnnotatedInteractionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnnotatedInteractionsType: valid values are %v", v, allowedLLMObsAnnotatedInteractionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnnotatedInteractionsType) IsValid() bool {
	for _, existing := range allowedLLMObsAnnotatedInteractionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnnotatedInteractionsType value.
func (v LLMObsAnnotatedInteractionsType) Ptr() *LLMObsAnnotatedInteractionsType {
	return &v
}
