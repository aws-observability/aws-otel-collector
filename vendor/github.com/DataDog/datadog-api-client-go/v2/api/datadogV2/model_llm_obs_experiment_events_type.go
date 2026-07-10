// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentEventsType Resource type for an experiment events collection.
type LLMObsExperimentEventsType string

// List of LLMObsExperimentEventsType.
const (
	LLMOBSEXPERIMENTEVENTSTYPE_EXPERIMENT_EVENTS LLMObsExperimentEventsType = "experiment_events"
)

var allowedLLMObsExperimentEventsTypeEnumValues = []LLMObsExperimentEventsType{
	LLMOBSEXPERIMENTEVENTSTYPE_EXPERIMENT_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentEventsType) GetAllowedValues() []LLMObsExperimentEventsType {
	return allowedLLMObsExperimentEventsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentEventsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentEventsType(value)
	return nil
}

// NewLLMObsExperimentEventsTypeFromValue returns a pointer to a valid LLMObsExperimentEventsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentEventsTypeFromValue(v string) (*LLMObsExperimentEventsType, error) {
	ev := LLMObsExperimentEventsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentEventsType: valid values are %v", v, allowedLLMObsExperimentEventsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentEventsType) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentEventsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentEventsType value.
func (v LLMObsExperimentEventsType) Ptr() *LLMObsExperimentEventsType {
	return &v
}
