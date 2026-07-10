// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpanType Resource type for a span item in an experiment spans response.
type LLMObsExperimentSpanType string

// List of LLMObsExperimentSpanType.
const (
	LLMOBSEXPERIMENTSPANTYPE_EXPERIMENTS_SPAN LLMObsExperimentSpanType = "experiments"
)

var allowedLLMObsExperimentSpanTypeEnumValues = []LLMObsExperimentSpanType{
	LLMOBSEXPERIMENTSPANTYPE_EXPERIMENTS_SPAN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentSpanType) GetAllowedValues() []LLMObsExperimentSpanType {
	return allowedLLMObsExperimentSpanTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentSpanType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentSpanType(value)
	return nil
}

// NewLLMObsExperimentSpanTypeFromValue returns a pointer to a valid LLMObsExperimentSpanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentSpanTypeFromValue(v string) (*LLMObsExperimentSpanType, error) {
	ev := LLMObsExperimentSpanType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentSpanType: valid values are %v", v, allowedLLMObsExperimentSpanTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentSpanType) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentSpanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentSpanType value.
func (v LLMObsExperimentSpanType) Ptr() *LLMObsExperimentSpanType {
	return &v
}
