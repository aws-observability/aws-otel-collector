// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRunStatusType Resource type of an LLM Observability patterns run status.
type LLMObsPatternsRunStatusType string

// List of LLMObsPatternsRunStatusType.
const (
	LLMOBSPATTERNSRUNSTATUSTYPE_TOPIC_DISCOVERY_RUN_STATUS LLMObsPatternsRunStatusType = "topic_discovery_run_status"
)

var allowedLLMObsPatternsRunStatusTypeEnumValues = []LLMObsPatternsRunStatusType{
	LLMOBSPATTERNSRUNSTATUSTYPE_TOPIC_DISCOVERY_RUN_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsRunStatusType) GetAllowedValues() []LLMObsPatternsRunStatusType {
	return allowedLLMObsPatternsRunStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsRunStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsRunStatusType(value)
	return nil
}

// NewLLMObsPatternsRunStatusTypeFromValue returns a pointer to a valid LLMObsPatternsRunStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsRunStatusTypeFromValue(v string) (*LLMObsPatternsRunStatusType, error) {
	ev := LLMObsPatternsRunStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsRunStatusType: valid values are %v", v, allowedLLMObsPatternsRunStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsRunStatusType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsRunStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsRunStatusType value.
func (v LLMObsPatternsRunStatusType) Ptr() *LLMObsPatternsRunStatusType {
	return &v
}
