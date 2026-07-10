// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsRequestType Resource type for triggering an LLM Observability patterns run.
type LLMObsPatternsRequestType string

// List of LLMObsPatternsRequestType.
const (
	LLMOBSPATTERNSREQUESTTYPE_TOPIC_DISCOVERY LLMObsPatternsRequestType = "topic_discovery"
)

var allowedLLMObsPatternsRequestTypeEnumValues = []LLMObsPatternsRequestType{
	LLMOBSPATTERNSREQUESTTYPE_TOPIC_DISCOVERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsRequestType) GetAllowedValues() []LLMObsPatternsRequestType {
	return allowedLLMObsPatternsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsRequestType(value)
	return nil
}

// NewLLMObsPatternsRequestTypeFromValue returns a pointer to a valid LLMObsPatternsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsRequestTypeFromValue(v string) (*LLMObsPatternsRequestType, error) {
	ev := LLMObsPatternsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsRequestType: valid values are %v", v, allowedLLMObsPatternsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsRequestType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsRequestType value.
func (v LLMObsPatternsRequestType) Ptr() *LLMObsPatternsRequestType {
	return &v
}
