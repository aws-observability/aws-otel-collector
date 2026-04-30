// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricTagsProcessorType The processor type. The value should always be `metric_tags`.
type ObservabilityPipelineMetricTagsProcessorType string

// List of ObservabilityPipelineMetricTagsProcessorType.
const (
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORTYPE_METRIC_TAGS ObservabilityPipelineMetricTagsProcessorType = "metric_tags"
)

var allowedObservabilityPipelineMetricTagsProcessorTypeEnumValues = []ObservabilityPipelineMetricTagsProcessorType{
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORTYPE_METRIC_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineMetricTagsProcessorType) GetAllowedValues() []ObservabilityPipelineMetricTagsProcessorType {
	return allowedObservabilityPipelineMetricTagsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineMetricTagsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineMetricTagsProcessorType(value)
	return nil
}

// NewObservabilityPipelineMetricTagsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineMetricTagsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineMetricTagsProcessorTypeFromValue(v string) (*ObservabilityPipelineMetricTagsProcessorType, error) {
	ev := ObservabilityPipelineMetricTagsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineMetricTagsProcessorType: valid values are %v", v, allowedObservabilityPipelineMetricTagsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineMetricTagsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineMetricTagsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineMetricTagsProcessorType value.
func (v ObservabilityPipelineMetricTagsProcessorType) Ptr() *ObservabilityPipelineMetricTagsProcessorType {
	return &v
}
