// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineReduceProcessorMergeStrategyStrategy The merge strategy to apply.
type ObservabilityPipelineReduceProcessorMergeStrategyStrategy string

// List of ObservabilityPipelineReduceProcessorMergeStrategyStrategy.
const (
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_DISCARD        ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "discard"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_RETAIN         ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "retain"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_SUM            ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "sum"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_MAX            ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "max"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_MIN            ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "min"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_ARRAY          ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "array"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT         ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "concat"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT_NEWLINE ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "concat_newline"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT_RAW     ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "concat_raw"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_SHORTEST_ARRAY ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "shortest_array"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_LONGEST_ARRAY  ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "longest_array"
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_FLAT_UNIQUE    ObservabilityPipelineReduceProcessorMergeStrategyStrategy = "flat_unique"
)

var allowedObservabilityPipelineReduceProcessorMergeStrategyStrategyEnumValues = []ObservabilityPipelineReduceProcessorMergeStrategyStrategy{
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_DISCARD,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_RETAIN,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_SUM,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_MAX,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_MIN,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_ARRAY,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT_NEWLINE,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_CONCAT_RAW,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_SHORTEST_ARRAY,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_LONGEST_ARRAY,
	OBSERVABILITYPIPELINEREDUCEPROCESSORMERGESTRATEGYSTRATEGY_FLAT_UNIQUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineReduceProcessorMergeStrategyStrategy) GetAllowedValues() []ObservabilityPipelineReduceProcessorMergeStrategyStrategy {
	return allowedObservabilityPipelineReduceProcessorMergeStrategyStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineReduceProcessorMergeStrategyStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineReduceProcessorMergeStrategyStrategy(value)
	return nil
}

// NewObservabilityPipelineReduceProcessorMergeStrategyStrategyFromValue returns a pointer to a valid ObservabilityPipelineReduceProcessorMergeStrategyStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineReduceProcessorMergeStrategyStrategyFromValue(v string) (*ObservabilityPipelineReduceProcessorMergeStrategyStrategy, error) {
	ev := ObservabilityPipelineReduceProcessorMergeStrategyStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineReduceProcessorMergeStrategyStrategy: valid values are %v", v, allowedObservabilityPipelineReduceProcessorMergeStrategyStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineReduceProcessorMergeStrategyStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineReduceProcessorMergeStrategyStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineReduceProcessorMergeStrategyStrategy value.
func (v ObservabilityPipelineReduceProcessorMergeStrategyStrategy) Ptr() *ObservabilityPipelineReduceProcessorMergeStrategyStrategy {
	return &v
}
