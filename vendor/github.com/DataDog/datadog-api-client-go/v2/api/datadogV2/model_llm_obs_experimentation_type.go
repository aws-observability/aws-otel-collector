// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationType Resource type for experimentation search and analytics operations.
type LLMObsExperimentationType string

// List of LLMObsExperimentationType.
const (
	LLMOBSEXPERIMENTATIONTYPE_EXPERIMENTATION LLMObsExperimentationType = "experimentation"
)

var allowedLLMObsExperimentationTypeEnumValues = []LLMObsExperimentationType{
	LLMOBSEXPERIMENTATIONTYPE_EXPERIMENTATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentationType) GetAllowedValues() []LLMObsExperimentationType {
	return allowedLLMObsExperimentationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentationType(value)
	return nil
}

// NewLLMObsExperimentationTypeFromValue returns a pointer to a valid LLMObsExperimentationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentationTypeFromValue(v string) (*LLMObsExperimentationType, error) {
	ev := LLMObsExperimentationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentationType: valid values are %v", v, allowedLLMObsExperimentationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentationType) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentationType value.
func (v LLMObsExperimentationType) Ptr() *LLMObsExperimentationType {
	return &v
}
