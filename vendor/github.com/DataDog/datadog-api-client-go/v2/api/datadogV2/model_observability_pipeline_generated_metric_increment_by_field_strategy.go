// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy Uses a numeric field in the log event as the metric increment.
type ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy string

// List of ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy.
const (
	OBSERVABILITYPIPELINEGENERATEDMETRICINCREMENTBYFIELDSTRATEGY_INCREMENT_BY_FIELD ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy = "increment_by_field"
)

var allowedObservabilityPipelineGeneratedMetricIncrementByFieldStrategyEnumValues = []ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy{
	OBSERVABILITYPIPELINEGENERATEDMETRICINCREMENTBYFIELDSTRATEGY_INCREMENT_BY_FIELD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) GetAllowedValues() []ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy {
	return allowedObservabilityPipelineGeneratedMetricIncrementByFieldStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy(value)
	return nil
}

// NewObservabilityPipelineGeneratedMetricIncrementByFieldStrategyFromValue returns a pointer to a valid ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGeneratedMetricIncrementByFieldStrategyFromValue(v string) (*ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy, error) {
	ev := ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy: valid values are %v", v, allowedObservabilityPipelineGeneratedMetricIncrementByFieldStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGeneratedMetricIncrementByFieldStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy value.
func (v ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy) Ptr() *ObservabilityPipelineGeneratedMetricIncrementByFieldStrategy {
	return &v
}
