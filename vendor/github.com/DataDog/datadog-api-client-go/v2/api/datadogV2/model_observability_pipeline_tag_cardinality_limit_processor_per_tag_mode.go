// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode How the per-tag override is applied. `limit_override` enforces a custom limit on the tag; `excluded` skips the tag from cardinality tracking.
type ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode string

// List of ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERTAGMODE_LIMIT_OVERRIDE ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode = "limit_override"
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERTAGMODE_EXCLUDED       ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode = "excluded"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorPerTagModeEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERTAGMODE_LIMIT_OVERRIDE,
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERTAGMODE_EXCLUDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorPerTagModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerTagModeFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerTagModeFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorPerTagModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorPerTagModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorPerTagMode {
	return &v
}
