// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationAssessment Assessment result for a label value.
type LLMObsAnnotationAssessment string

// List of LLMObsAnnotationAssessment.
const (
	LLMOBSANNOTATIONASSESSMENT_PASS LLMObsAnnotationAssessment = "pass"
	LLMOBSANNOTATIONASSESSMENT_FAIL LLMObsAnnotationAssessment = "fail"
)

var allowedLLMObsAnnotationAssessmentEnumValues = []LLMObsAnnotationAssessment{
	LLMOBSANNOTATIONASSESSMENT_PASS,
	LLMOBSANNOTATIONASSESSMENT_FAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsAnnotationAssessment) GetAllowedValues() []LLMObsAnnotationAssessment {
	return allowedLLMObsAnnotationAssessmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsAnnotationAssessment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsAnnotationAssessment(value)
	return nil
}

// NewLLMObsAnnotationAssessmentFromValue returns a pointer to a valid LLMObsAnnotationAssessment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsAnnotationAssessmentFromValue(v string) (*LLMObsAnnotationAssessment, error) {
	ev := LLMObsAnnotationAssessment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsAnnotationAssessment: valid values are %v", v, allowedLLMObsAnnotationAssessmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsAnnotationAssessment) IsValid() bool {
	for _, existing := range allowedLLMObsAnnotationAssessmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsAnnotationAssessment value.
func (v LLMObsAnnotationAssessment) Ptr() *LLMObsAnnotationAssessment {
	return &v
}
