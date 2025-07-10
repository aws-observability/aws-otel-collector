// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGenerateMetricsProcessorType The processor type. Always `generate_datadog_metrics`.
type ObservabilityPipelineGenerateMetricsProcessorType string

// List of ObservabilityPipelineGenerateMetricsProcessorType.
const (
	OBSERVABILITYPIPELINEGENERATEMETRICSPROCESSORTYPE_GENERATE_DATADOG_METRICS ObservabilityPipelineGenerateMetricsProcessorType = "generate_datadog_metrics"
)

var allowedObservabilityPipelineGenerateMetricsProcessorTypeEnumValues = []ObservabilityPipelineGenerateMetricsProcessorType{
	OBSERVABILITYPIPELINEGENERATEMETRICSPROCESSORTYPE_GENERATE_DATADOG_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineGenerateMetricsProcessorType) GetAllowedValues() []ObservabilityPipelineGenerateMetricsProcessorType {
	return allowedObservabilityPipelineGenerateMetricsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineGenerateMetricsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineGenerateMetricsProcessorType(value)
	return nil
}

// NewObservabilityPipelineGenerateMetricsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineGenerateMetricsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineGenerateMetricsProcessorTypeFromValue(v string) (*ObservabilityPipelineGenerateMetricsProcessorType, error) {
	ev := ObservabilityPipelineGenerateMetricsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineGenerateMetricsProcessorType: valid values are %v", v, allowedObservabilityPipelineGenerateMetricsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineGenerateMetricsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineGenerateMetricsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineGenerateMetricsProcessorType value.
func (v ObservabilityPipelineGenerateMetricsProcessorType) Ptr() *ObservabilityPipelineGenerateMetricsProcessorType {
	return &v
}
