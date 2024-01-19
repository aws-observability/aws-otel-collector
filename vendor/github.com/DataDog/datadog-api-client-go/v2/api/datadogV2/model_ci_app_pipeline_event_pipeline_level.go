// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventPipelineLevel Used to distinguish between pipelines, stages, jobs, and steps.
type CIAppPipelineEventPipelineLevel string

// List of CIAppPipelineEventPipelineLevel.
const (
	CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE CIAppPipelineEventPipelineLevel = "pipeline"
)

var allowedCIAppPipelineEventPipelineLevelEnumValues = []CIAppPipelineEventPipelineLevel{
	CIAPPPIPELINEEVENTPIPELINELEVEL_PIPELINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventPipelineLevel) GetAllowedValues() []CIAppPipelineEventPipelineLevel {
	return allowedCIAppPipelineEventPipelineLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventPipelineLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventPipelineLevel(value)
	return nil
}

// NewCIAppPipelineEventPipelineLevelFromValue returns a pointer to a valid CIAppPipelineEventPipelineLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventPipelineLevelFromValue(v string) (*CIAppPipelineEventPipelineLevel, error) {
	ev := CIAppPipelineEventPipelineLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventPipelineLevel: valid values are %v", v, allowedCIAppPipelineEventPipelineLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventPipelineLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventPipelineLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventPipelineLevel value.
func (v CIAppPipelineEventPipelineLevel) Ptr() *CIAppPipelineEventPipelineLevel {
	return &v
}
