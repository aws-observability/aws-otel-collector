// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAggregateProcessorMode The aggregation mode applied to metrics that share the same name and tags within the interval.
type ObservabilityPipelineAggregateProcessorMode string

// List of ObservabilityPipelineAggregateProcessorMode.
const (
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_AUTO   ObservabilityPipelineAggregateProcessorMode = "auto"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_SUM    ObservabilityPipelineAggregateProcessorMode = "sum"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_LATEST ObservabilityPipelineAggregateProcessorMode = "latest"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_COUNT  ObservabilityPipelineAggregateProcessorMode = "count"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MAX    ObservabilityPipelineAggregateProcessorMode = "max"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MIN    ObservabilityPipelineAggregateProcessorMode = "min"
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MEAN   ObservabilityPipelineAggregateProcessorMode = "mean"
)

var allowedObservabilityPipelineAggregateProcessorModeEnumValues = []ObservabilityPipelineAggregateProcessorMode{
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_AUTO,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_SUM,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_LATEST,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_COUNT,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MAX,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MIN,
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORMODE_MEAN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAggregateProcessorMode) GetAllowedValues() []ObservabilityPipelineAggregateProcessorMode {
	return allowedObservabilityPipelineAggregateProcessorModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAggregateProcessorMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAggregateProcessorMode(value)
	return nil
}

// NewObservabilityPipelineAggregateProcessorModeFromValue returns a pointer to a valid ObservabilityPipelineAggregateProcessorMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAggregateProcessorModeFromValue(v string) (*ObservabilityPipelineAggregateProcessorMode, error) {
	ev := ObservabilityPipelineAggregateProcessorMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAggregateProcessorMode: valid values are %v", v, allowedObservabilityPipelineAggregateProcessorModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAggregateProcessorMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAggregateProcessorModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAggregateProcessorMode value.
func (v ObservabilityPipelineAggregateProcessorMode) Ptr() *ObservabilityPipelineAggregateProcessorMode {
	return &v
}
