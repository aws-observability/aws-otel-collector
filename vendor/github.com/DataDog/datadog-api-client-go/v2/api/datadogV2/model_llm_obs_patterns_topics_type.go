// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTopicsType Resource type of an LLM Observability patterns topics response.
type LLMObsPatternsTopicsType string

// List of LLMObsPatternsTopicsType.
const (
	LLMOBSPATTERNSTOPICSTYPE_GET_TOPICS_RESPONSE LLMObsPatternsTopicsType = "get_topics_response"
)

var allowedLLMObsPatternsTopicsTypeEnumValues = []LLMObsPatternsTopicsType{
	LLMOBSPATTERNSTOPICSTYPE_GET_TOPICS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsTopicsType) GetAllowedValues() []LLMObsPatternsTopicsType {
	return allowedLLMObsPatternsTopicsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsTopicsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsTopicsType(value)
	return nil
}

// NewLLMObsPatternsTopicsTypeFromValue returns a pointer to a valid LLMObsPatternsTopicsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsTopicsTypeFromValue(v string) (*LLMObsPatternsTopicsType, error) {
	ev := LLMObsPatternsTopicsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsTopicsType: valid values are %v", v, allowedLLMObsPatternsTopicsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsTopicsType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsTopicsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsTopicsType value.
func (v LLMObsPatternsTopicsType) Ptr() *LLMObsPatternsTopicsType {
	return &v
}
