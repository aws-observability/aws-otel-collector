// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineEventStageLevel Used to distinguish between pipelines, stages, jobs and steps.
type CIAppPipelineEventStageLevel string

// List of CIAppPipelineEventStageLevel.
const (
	CIAPPPIPELINEEVENTSTAGELEVEL_STAGE CIAppPipelineEventStageLevel = "stage"
)

var allowedCIAppPipelineEventStageLevelEnumValues = []CIAppPipelineEventStageLevel{
	CIAPPPIPELINEEVENTSTAGELEVEL_STAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineEventStageLevel) GetAllowedValues() []CIAppPipelineEventStageLevel {
	return allowedCIAppPipelineEventStageLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineEventStageLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineEventStageLevel(value)
	return nil
}

// NewCIAppPipelineEventStageLevelFromValue returns a pointer to a valid CIAppPipelineEventStageLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineEventStageLevelFromValue(v string) (*CIAppPipelineEventStageLevel, error) {
	ev := CIAppPipelineEventStageLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineEventStageLevel: valid values are %v", v, allowedCIAppPipelineEventStageLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineEventStageLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineEventStageLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineEventStageLevel value.
func (v CIAppPipelineEventStageLevel) Ptr() *CIAppPipelineEventStageLevel {
	return &v
}
