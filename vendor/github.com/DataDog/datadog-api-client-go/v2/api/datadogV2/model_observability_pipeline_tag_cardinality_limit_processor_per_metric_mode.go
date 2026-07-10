// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode How the per-metric override is applied. `tracked` enforces a custom limit; `excluded` skips the metric entirely.
type ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode string

// List of ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERMETRICMODE_TRACKED  ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode = "tracked"
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERMETRICMODE_EXCLUDED ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode = "excluded"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERMETRICMODE_TRACKED,
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORPERMETRICMODE_EXCLUDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorPerMetricModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorPerMetricMode {
	return &v
}
