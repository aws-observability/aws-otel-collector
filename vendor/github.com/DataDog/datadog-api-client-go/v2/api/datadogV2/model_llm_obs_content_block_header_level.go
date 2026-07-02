// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsContentBlockHeaderLevel Visual size for a `header` block.
type LLMObsContentBlockHeaderLevel string

// List of LLMObsContentBlockHeaderLevel.
const (
	LLMOBSCONTENTBLOCKHEADERLEVEL_SM LLMObsContentBlockHeaderLevel = "sm"
	LLMOBSCONTENTBLOCKHEADERLEVEL_MD LLMObsContentBlockHeaderLevel = "md"
	LLMOBSCONTENTBLOCKHEADERLEVEL_LG LLMObsContentBlockHeaderLevel = "lg"
	LLMOBSCONTENTBLOCKHEADERLEVEL_XL LLMObsContentBlockHeaderLevel = "xl"
)

var allowedLLMObsContentBlockHeaderLevelEnumValues = []LLMObsContentBlockHeaderLevel{
	LLMOBSCONTENTBLOCKHEADERLEVEL_SM,
	LLMOBSCONTENTBLOCKHEADERLEVEL_MD,
	LLMOBSCONTENTBLOCKHEADERLEVEL_LG,
	LLMOBSCONTENTBLOCKHEADERLEVEL_XL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsContentBlockHeaderLevel) GetAllowedValues() []LLMObsContentBlockHeaderLevel {
	return allowedLLMObsContentBlockHeaderLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsContentBlockHeaderLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsContentBlockHeaderLevel(value)
	return nil
}

// NewLLMObsContentBlockHeaderLevelFromValue returns a pointer to a valid LLMObsContentBlockHeaderLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsContentBlockHeaderLevelFromValue(v string) (*LLMObsContentBlockHeaderLevel, error) {
	ev := LLMObsContentBlockHeaderLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsContentBlockHeaderLevel: valid values are %v", v, allowedLLMObsContentBlockHeaderLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsContentBlockHeaderLevel) IsValid() bool {
	for _, existing := range allowedLLMObsContentBlockHeaderLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsContentBlockHeaderLevel value.
func (v LLMObsContentBlockHeaderLevel) Ptr() *LLMObsContentBlockHeaderLevel {
	return &v
}
