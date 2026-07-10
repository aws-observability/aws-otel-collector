// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsConfigType Resource type of an LLM Observability patterns configuration.
type LLMObsPatternsConfigType string

// List of LLMObsPatternsConfigType.
const (
	LLMOBSPATTERNSCONFIGTYPE_TOPIC_DISCOVERY_CONFIGS LLMObsPatternsConfigType = "topic_discovery_configs"
)

var allowedLLMObsPatternsConfigTypeEnumValues = []LLMObsPatternsConfigType{
	LLMOBSPATTERNSCONFIGTYPE_TOPIC_DISCOVERY_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsConfigType) GetAllowedValues() []LLMObsPatternsConfigType {
	return allowedLLMObsPatternsConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsConfigType(value)
	return nil
}

// NewLLMObsPatternsConfigTypeFromValue returns a pointer to a valid LLMObsPatternsConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsConfigTypeFromValue(v string) (*LLMObsPatternsConfigType, error) {
	ev := LLMObsPatternsConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsConfigType: valid values are %v", v, allowedLLMObsPatternsConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsConfigType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsConfigType value.
func (v LLMObsPatternsConfigType) Ptr() *LLMObsPatternsConfigType {
	return &v
}
