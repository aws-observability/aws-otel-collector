// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSortFieldDirection Sort direction.
type LLMObsExperimentationSortFieldDirection string

// List of LLMObsExperimentationSortFieldDirection.
const (
	LLMOBSEXPERIMENTATIONSORTFIELDDIRECTION_ASC  LLMObsExperimentationSortFieldDirection = "asc"
	LLMOBSEXPERIMENTATIONSORTFIELDDIRECTION_DESC LLMObsExperimentationSortFieldDirection = "desc"
)

var allowedLLMObsExperimentationSortFieldDirectionEnumValues = []LLMObsExperimentationSortFieldDirection{
	LLMOBSEXPERIMENTATIONSORTFIELDDIRECTION_ASC,
	LLMOBSEXPERIMENTATIONSORTFIELDDIRECTION_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LLMObsExperimentationSortFieldDirection) GetAllowedValues() []LLMObsExperimentationSortFieldDirection {
	return allowedLLMObsExperimentationSortFieldDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LLMObsExperimentationSortFieldDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LLMObsExperimentationSortFieldDirection(value)
	return nil
}

// NewLLMObsExperimentationSortFieldDirectionFromValue returns a pointer to a valid LLMObsExperimentationSortFieldDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLLMObsExperimentationSortFieldDirectionFromValue(v string) (*LLMObsExperimentationSortFieldDirection, error) {
	ev := LLMObsExperimentationSortFieldDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LLMObsExperimentationSortFieldDirection: valid values are %v", v, allowedLLMObsExperimentationSortFieldDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LLMObsExperimentationSortFieldDirection) IsValid() bool {
	for _, existing := range allowedLLMObsExperimentationSortFieldDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LLMObsExperimentationSortFieldDirection value.
func (v LLMObsExperimentationSortFieldDirection) Ptr() *LLMObsExperimentationSortFieldDirection {
	return &v
}
