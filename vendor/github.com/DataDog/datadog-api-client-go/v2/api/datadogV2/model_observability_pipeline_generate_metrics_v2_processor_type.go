// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGenerateMetricsV2ProcessorType The processor type. Always `generate_metrics`.
type ObservabilityPipelineGenerateMetricsV2ProcessorType string

// List of ObservabilityPipelineGenerateMetricsV2ProcessorType.
const (
	OBSERVABILITYPIPELINEGENERATEMETRICSV2PROCESSORTYPE_GENERATE_METRICS ObservabilityPipelineGenerateMetricsV2ProcessorType = "generate_metrics"
)

var allowedObservabilityPipelineGenerateMetricsV2ProcessorTypeEnumValues = []ObservabilityPipelineGenerateMetricsV2ProcessorType{
	OBSERVABILITYPIPELINEGENERATEMETRICSV2PROCESSORTYPE_GENERATE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGenerateMetricsV2ProcessorType) GetAllowedValues() []ObservabilityPipelineGenerateMetricsV2ProcessorType {
	return allowedObservabilityPipelineGenerateMetricsV2ProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGenerateMetricsV2ProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGenerateMetricsV2ProcessorType(value)
	return nil
}

// NewObservabilityPipelineGenerateMetricsV2ProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineGenerateMetricsV2ProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGenerateMetricsV2ProcessorTypeFromValue(v string) (*ObservabilityPipelineGenerateMetricsV2ProcessorType, error) {
	ev := ObservabilityPipelineGenerateMetricsV2ProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGenerateMetricsV2ProcessorType: valid values are %v", v, allowedObservabilityPipelineGenerateMetricsV2ProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGenerateMetricsV2ProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGenerateMetricsV2ProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGenerateMetricsV2ProcessorType value.
func (v ObservabilityPipelineGenerateMetricsV2ProcessorType) Ptr() *ObservabilityPipelineGenerateMetricsV2ProcessorType {
	return &v
}
