// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetricIncrementByOneStrategy Increments the metric by 1 for each matching event.
type ObservabilityPipelineGeneratedMetricIncrementByOneStrategy string

// List of ObservabilityPipelineGeneratedMetricIncrementByOneStrategy.
const (
	OBSERVABILITYPIPELINEGENERATEDMETRICINCREMENTBYONESTRATEGY_INCREMENT_BY_ONE ObservabilityPipelineGeneratedMetricIncrementByOneStrategy = "increment_by_one"
)

var allowedObservabilityPipelineGeneratedMetricIncrementByOneStrategyEnumValues = []ObservabilityPipelineGeneratedMetricIncrementByOneStrategy{
	OBSERVABILITYPIPELINEGENERATEDMETRICINCREMENTBYONESTRATEGY_INCREMENT_BY_ONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) GetAllowedValues() []ObservabilityPipelineGeneratedMetricIncrementByOneStrategy {
	return allowedObservabilityPipelineGeneratedMetricIncrementByOneStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGeneratedMetricIncrementByOneStrategy(value)
	return nil
}

// NewObservabilityPipelineGeneratedMetricIncrementByOneStrategyFromValue returns a pointer to a valid ObservabilityPipelineGeneratedMetricIncrementByOneStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGeneratedMetricIncrementByOneStrategyFromValue(v string) (*ObservabilityPipelineGeneratedMetricIncrementByOneStrategy, error) {
	ev := ObservabilityPipelineGeneratedMetricIncrementByOneStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGeneratedMetricIncrementByOneStrategy: valid values are %v", v, allowedObservabilityPipelineGeneratedMetricIncrementByOneStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGeneratedMetricIncrementByOneStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGeneratedMetricIncrementByOneStrategy value.
func (v ObservabilityPipelineGeneratedMetricIncrementByOneStrategy) Ptr() *ObservabilityPipelineGeneratedMetricIncrementByOneStrategy {
	return &v
}
