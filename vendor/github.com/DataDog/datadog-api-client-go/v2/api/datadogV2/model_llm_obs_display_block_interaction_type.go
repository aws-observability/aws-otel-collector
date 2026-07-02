// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDisplayBlockInteractionType Type discriminator for a `display_block` interaction.
type LLMObsDisplayBlockInteractionType string

// List of LLMObsDisplayBlockInteractionType.
const (
	LLMOBSDISPLAYBLOCKINTERACTIONTYPE_DISPLAY_BLOCK LLMObsDisplayBlockInteractionType = "display_block"
)

var allowedLLMObsDisplayBlockInteractionTypeEnumValues = []LLMObsDisplayBlockInteractionType{
	LLMOBSDISPLAYBLOCKINTERACTIONTYPE_DISPLAY_BLOCK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsDisplayBlockInteractionType) GetAllowedValues() []LLMObsDisplayBlockInteractionType {
	return allowedLLMObsDisplayBlockInteractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsDisplayBlockInteractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsDisplayBlockInteractionType(value)
	return nil
}

// NewLLMObsDisplayBlockInteractionTypeFromValue returns a pointer to a valid LLMObsDisplayBlockInteractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsDisplayBlockInteractionTypeFromValue(v string) (*LLMObsDisplayBlockInteractionType, error) {
	ev := LLMObsDisplayBlockInteractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsDisplayBlockInteractionType: valid values are %v", v, allowedLLMObsDisplayBlockInteractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsDisplayBlockInteractionType) IsValid() bool {
	for _, existing := range allowedLLMObsDisplayBlockInteractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsDisplayBlockInteractionType value.
func (v LLMObsDisplayBlockInteractionType) Ptr() *LLMObsDisplayBlockInteractionType {
	return &v
}
