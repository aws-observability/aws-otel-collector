// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDedupeProcessorMode The deduplication mode to apply to the fields.
type ObservabilityPipelineDedupeProcessorMode string

// List of ObservabilityPipelineDedupeProcessorMode.
const (
	OBSERVABILITYPIPELINEDEDUPEPROCESSORMODE_MATCH  ObservabilityPipelineDedupeProcessorMode = "match"
	OBSERVABILITYPIPELINEDEDUPEPROCESSORMODE_IGNORE ObservabilityPipelineDedupeProcessorMode = "ignore"
)

var allowedObservabilityPipelineDedupeProcessorModeEnumValues = []ObservabilityPipelineDedupeProcessorMode{
	OBSERVABILITYPIPELINEDEDUPEPROCESSORMODE_MATCH,
	OBSERVABILITYPIPELINEDEDUPEPROCESSORMODE_IGNORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDedupeProcessorMode) GetAllowedValues() []ObservabilityPipelineDedupeProcessorMode {
	return allowedObservabilityPipelineDedupeProcessorModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDedupeProcessorMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDedupeProcessorMode(value)
	return nil
}

// NewObservabilityPipelineDedupeProcessorModeFromValue returns a pointer to a valid ObservabilityPipelineDedupeProcessorMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDedupeProcessorModeFromValue(v string) (*ObservabilityPipelineDedupeProcessorMode, error) {
	ev := ObservabilityPipelineDedupeProcessorMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDedupeProcessorMode: valid values are %v", v, allowedObservabilityPipelineDedupeProcessorModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDedupeProcessorMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDedupeProcessorModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDedupeProcessorMode value.
func (v ObservabilityPipelineDedupeProcessorMode) Ptr() *ObservabilityPipelineDedupeProcessorMode {
	return &v
}
