// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTriggerResponseType Resource type of an LLM Observability patterns trigger response.
type LLMObsPatternsTriggerResponseType string

// List of LLMObsPatternsTriggerResponseType.
const (
	LLMOBSPATTERNSTRIGGERRESPONSETYPE_TOPIC_DISCOVERY_RUN LLMObsPatternsTriggerResponseType = "topic_discovery_run"
)

var allowedLLMObsPatternsTriggerResponseTypeEnumValues = []LLMObsPatternsTriggerResponseType{
	LLMOBSPATTERNSTRIGGERRESPONSETYPE_TOPIC_DISCOVERY_RUN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsPatternsTriggerResponseType) GetAllowedValues() []LLMObsPatternsTriggerResponseType {
	return allowedLLMObsPatternsTriggerResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsPatternsTriggerResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsPatternsTriggerResponseType(value)
	return nil
}

// NewLLMObsPatternsTriggerResponseTypeFromValue returns a pointer to a valid LLMObsPatternsTriggerResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsPatternsTriggerResponseTypeFromValue(v string) (*LLMObsPatternsTriggerResponseType, error) {
	ev := LLMObsPatternsTriggerResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsPatternsTriggerResponseType: valid values are %v", v, allowedLLMObsPatternsTriggerResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsPatternsTriggerResponseType) IsValid() bool {
	for _, existing := range allowedLLMObsPatternsTriggerResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsPatternsTriggerResponseType value.
func (v LLMObsPatternsTriggerResponseType) Ptr() *LLMObsPatternsTriggerResponseType {
	return &v
}
