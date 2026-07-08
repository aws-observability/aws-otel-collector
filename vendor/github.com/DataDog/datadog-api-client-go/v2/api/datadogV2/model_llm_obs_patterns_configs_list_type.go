// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsConfigsListType Resource type of a list of LLM Observability patterns configurations.
type LLMObsPatternsConfigsListType string

// List of LLMObsPatternsConfigsListType.
const (
	LLMOBSPATTERNSCONFIGSLISTTYPE_LIST_TOPIC_DISCOVERY_CONFIGS_RESPONSE LLMObsPatternsConfigsListType = "list_topic_discovery_configs_response"
)

var allowedLLMObsPatternsConfigsListTypeEnumValues = []LLMObsPatternsConfigsListType{
	LLMOBSPATTERNSCONFIGSLISTTYPE_LIST_TOPIC_DISCOVERY_CONFIGS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsConfigsListType) GetAllowedValues() []LLMObsPatternsConfigsListType {
	return allowedLLMObsPatternsConfigsListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsConfigsListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsConfigsListType(value)
	return nil
}

// NewLLMObsPatternsConfigsListTypeFromValue returns a pointer to a valid LLMObsPatternsConfigsListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsConfigsListTypeFromValue(v string) (*LLMObsPatternsConfigsListType, error) {
	ev := LLMObsPatternsConfigsListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsConfigsListType: valid values are %v", v, allowedLLMObsPatternsConfigsListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsConfigsListType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsConfigsListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsConfigsListType value.
func (v LLMObsPatternsConfigsListType) Ptr() *LLMObsPatternsConfigsListType {
	return &v
}
