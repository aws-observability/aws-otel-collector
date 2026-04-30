// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentSpanStatus Status of the span.
type LLMObsExperimentSpanStatus string

// List of LLMObsExperimentSpanStatus.
const (
	LLMOBSEXPERIMENTSPANSTATUS_OK    LLMObsExperimentSpanStatus = "ok"
	LLMOBSEXPERIMENTSPANSTATUS_ERROR LLMObsExperimentSpanStatus = "error"
)

var allowedLLMObsExperimentSpanStatusEnumValues = []LLMObsExperimentSpanStatus{
	LLMOBSEXPERIMENTSPANSTATUS_OK,
	LLMOBSEXPERIMENTSPANSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentSpanStatus) GetAllowedValues() []LLMObsExperimentSpanStatus {
	return allowedLLMObsExperimentSpanStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentSpanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentSpanStatus(value)
	return nil
}

// NewLLMObsExperimentSpanStatusFromValue returns a pointer to a valid LLMObsExperimentSpanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentSpanStatusFromValue(v string) (*LLMObsExperimentSpanStatus, error) {
	ev := LLMObsExperimentSpanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentSpanStatus: valid values are %v", v, allowedLLMObsExperimentSpanStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentSpanStatus) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentSpanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentSpanStatus value.
func (v LLMObsExperimentSpanStatus) Ptr() *LLMObsExperimentSpanStatus {
	return &v
}
