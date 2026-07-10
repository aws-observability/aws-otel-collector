// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRunsListType Resource type of a list of LLM Observability patterns runs.
type LLMObsPatternsRunsListType string

// List of LLMObsPatternsRunsListType.
const (
	LLMOBSPATTERNSRUNSLISTTYPE_LIST_TOPIC_DISCOVERY_RUNS_RESPONSE LLMObsPatternsRunsListType = "list_topic_discovery_runs_response"
)

var allowedLLMObsPatternsRunsListTypeEnumValues = []LLMObsPatternsRunsListType{
	LLMOBSPATTERNSRUNSLISTTYPE_LIST_TOPIC_DISCOVERY_RUNS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsRunsListType) GetAllowedValues() []LLMObsPatternsRunsListType {
	return allowedLLMObsPatternsRunsListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsRunsListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsRunsListType(value)
	return nil
}

// NewLLMObsPatternsRunsListTypeFromValue returns a pointer to a valid LLMObsPatternsRunsListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsRunsListTypeFromValue(v string) (*LLMObsPatternsRunsListType, error) {
	ev := LLMObsPatternsRunsListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsRunsListType: valid values are %v", v, allowedLLMObsPatternsRunsListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsRunsListType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsRunsListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsRunsListType value.
func (v LLMObsPatternsRunsListType) Ptr() *LLMObsPatternsRunsListType {
	return &v
}
