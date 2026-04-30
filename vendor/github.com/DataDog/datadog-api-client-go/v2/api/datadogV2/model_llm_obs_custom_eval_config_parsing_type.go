// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCustomEvalConfigParsingType Output parsing type for a custom LLM judge evaluator.
type LLMObsCustomEvalConfigParsingType string

// List of LLMObsCustomEvalConfigParsingType.
const (
	LLMOBSCUSTOMEVALCONFIGPARSINGTYPE_STRUCTURED_OUTPUT LLMObsCustomEvalConfigParsingType = "structured_output"
	LLMOBSCUSTOMEVALCONFIGPARSINGTYPE_JSON              LLMObsCustomEvalConfigParsingType = "json"
)

var allowedLLMObsCustomEvalConfigParsingTypeEnumValues = []LLMObsCustomEvalConfigParsingType{
	LLMOBSCUSTOMEVALCONFIGPARSINGTYPE_STRUCTURED_OUTPUT,
	LLMOBSCUSTOMEVALCONFIGPARSINGTYPE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsCustomEvalConfigParsingType) GetAllowedValues() []LLMObsCustomEvalConfigParsingType {
	return allowedLLMObsCustomEvalConfigParsingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsCustomEvalConfigParsingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsCustomEvalConfigParsingType(value)
	return nil
}

// NewLLMObsCustomEvalConfigParsingTypeFromValue returns a pointer to a valid LLMObsCustomEvalConfigParsingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsCustomEvalConfigParsingTypeFromValue(v string) (*LLMObsCustomEvalConfigParsingType, error) {
	ev := LLMObsCustomEvalConfigParsingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsCustomEvalConfigParsingType: valid values are %v", v, allowedLLMObsCustomEvalConfigParsingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsCustomEvalConfigParsingType) IsValid() bool {
	for _, existing := range allowedLLMObsCustomEvalConfigParsingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsCustomEvalConfigParsingType value.
func (v LLMObsCustomEvalConfigParsingType) Ptr() *LLMObsCustomEvalConfigParsingType {
	return &v
}
