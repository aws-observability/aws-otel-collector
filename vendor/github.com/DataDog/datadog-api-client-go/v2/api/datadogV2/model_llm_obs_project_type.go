// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsProjectType Resource type of an LLM Observability project.
type LLMObsProjectType string

// List of LLMObsProjectType.
const (
	LLMOBSPROJECTTYPE_PROJECTS LLMObsProjectType = "projects"
)

var allowedLLMObsProjectTypeEnumValues = []LLMObsProjectType{
	LLMOBSPROJECTTYPE_PROJECTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsProjectType) GetAllowedValues() []LLMObsProjectType {
	return allowedLLMObsProjectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsProjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsProjectType(value)
	return nil
}

// NewLLMObsProjectTypeFromValue returns a pointer to a valid LLMObsProjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsProjectTypeFromValue(v string) (*LLMObsProjectType, error) {
	ev := LLMObsProjectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsProjectType: valid values are %v", v, allowedLLMObsProjectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsProjectType) IsValid() bool {
	for _, existing := range allowedLLMObsProjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsProjectType value.
func (v LLMObsProjectType) Ptr() *LLMObsProjectType {
	return &v
}
