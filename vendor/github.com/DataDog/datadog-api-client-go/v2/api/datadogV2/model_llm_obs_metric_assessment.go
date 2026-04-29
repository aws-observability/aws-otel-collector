// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsMetricAssessment Assessment result for an LLM Observability experiment metric.
type LLMObsMetricAssessment string

// List of LLMObsMetricAssessment.
const (
	LLMOBSMETRICASSESSMENT_PASS LLMObsMetricAssessment = "pass"
	LLMOBSMETRICASSESSMENT_FAIL LLMObsMetricAssessment = "fail"
)

var allowedLLMObsMetricAssessmentEnumValues = []LLMObsMetricAssessment{
	LLMOBSMETRICASSESSMENT_PASS,
	LLMOBSMETRICASSESSMENT_FAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsMetricAssessment) GetAllowedValues() []LLMObsMetricAssessment {
	return allowedLLMObsMetricAssessmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsMetricAssessment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsMetricAssessment(value)
	return nil
}

// NewLLMObsMetricAssessmentFromValue returns a pointer to a valid LLMObsMetricAssessment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsMetricAssessmentFromValue(v string) (*LLMObsMetricAssessment, error) {
	ev := LLMObsMetricAssessment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsMetricAssessment: valid values are %v", v, allowedLLMObsMetricAssessmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsMetricAssessment) IsValid() bool {
	for _, existing := range allowedLLMObsMetricAssessmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsMetricAssessment value.
func (v LLMObsMetricAssessment) Ptr() *LLMObsMetricAssessment {
	return &v
}
