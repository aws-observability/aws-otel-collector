// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppPipelineLevel Pipeline execution level.
type CIAppPipelineLevel string

// List of CIAppPipelineLevel.
const (
	CIAPPPIPELINELEVEL_PIPELINE CIAppPipelineLevel = "pipeline"
	CIAPPPIPELINELEVEL_STAGE    CIAppPipelineLevel = "stage"
	CIAPPPIPELINELEVEL_JOB      CIAppPipelineLevel = "job"
	CIAPPPIPELINELEVEL_STEP     CIAppPipelineLevel = "step"
	CIAPPPIPELINELEVEL_CUSTOM   CIAppPipelineLevel = "custom"
)

var allowedCIAppPipelineLevelEnumValues = []CIAppPipelineLevel{
	CIAPPPIPELINELEVEL_PIPELINE,
	CIAPPPIPELINELEVEL_STAGE,
	CIAPPPIPELINELEVEL_JOB,
	CIAPPPIPELINELEVEL_STEP,
	CIAPPPIPELINELEVEL_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppPipelineLevel) GetAllowedValues() []CIAppPipelineLevel {
	return allowedCIAppPipelineLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppPipelineLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppPipelineLevel(value)
	return nil
}

// NewCIAppPipelineLevelFromValue returns a pointer to a valid CIAppPipelineLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppPipelineLevelFromValue(v string) (*CIAppPipelineLevel, error) {
	ev := CIAppPipelineLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppPipelineLevel: valid values are %v", v, allowedCIAppPipelineLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppPipelineLevel) IsValid() bool {
	for _, existing := range allowedCIAppPipelineLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppPipelineLevel value.
func (v CIAppPipelineLevel) Ptr() *CIAppPipelineLevel {
	return &v
}
