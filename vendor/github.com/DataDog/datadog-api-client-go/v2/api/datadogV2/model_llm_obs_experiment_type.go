// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentType Resource type of an LLM Observability experiment.
type LLMObsExperimentType string

// List of LLMObsExperimentType.
const (
	LLMOBSEXPERIMENTTYPE_EXPERIMENTS LLMObsExperimentType = "experiments"
)

var allowedLLMObsExperimentTypeEnumValues = []LLMObsExperimentType{
	LLMOBSEXPERIMENTTYPE_EXPERIMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentType) GetAllowedValues() []LLMObsExperimentType {
	return allowedLLMObsExperimentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentType(value)
	return nil
}

// NewLLMObsExperimentTypeFromValue returns a pointer to a valid LLMObsExperimentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentTypeFromValue(v string) (*LLMObsExperimentType, error) {
	ev := LLMObsExperimentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentType: valid values are %v", v, allowedLLMObsExperimentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentType) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentType value.
func (v LLMObsExperimentType) Ptr() *LLMObsExperimentType {
	return &v
}
