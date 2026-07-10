// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAggregateProcessorType The processor type. The value must be `aggregate`.
type ObservabilityPipelineAggregateProcessorType string

// List of ObservabilityPipelineAggregateProcessorType.
const (
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORTYPE_AGGREGATE ObservabilityPipelineAggregateProcessorType = "aggregate"
)

var allowedObservabilityPipelineAggregateProcessorTypeEnumValues = []ObservabilityPipelineAggregateProcessorType{
	OBSERVABILITYPIPELINEAGGREGATEPROCESSORTYPE_AGGREGATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAggregateProcessorType) GetAllowedValues() []ObservabilityPipelineAggregateProcessorType {
	return allowedObservabilityPipelineAggregateProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAggregateProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAggregateProcessorType(value)
	return nil
}

// NewObservabilityPipelineAggregateProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineAggregateProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAggregateProcessorTypeFromValue(v string) (*ObservabilityPipelineAggregateProcessorType, error) {
	ev := ObservabilityPipelineAggregateProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAggregateProcessorType: valid values are %v", v, allowedObservabilityPipelineAggregateProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAggregateProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAggregateProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAggregateProcessorType value.
func (v ObservabilityPipelineAggregateProcessorType) Ptr() *ObservabilityPipelineAggregateProcessorType {
	return &v
}
