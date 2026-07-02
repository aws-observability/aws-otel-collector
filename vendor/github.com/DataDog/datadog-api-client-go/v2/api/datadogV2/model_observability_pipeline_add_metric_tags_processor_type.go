// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddMetricTagsProcessorType The processor type. The value must be `add_metric_tags`.
type ObservabilityPipelineAddMetricTagsProcessorType string

// List of ObservabilityPipelineAddMetricTagsProcessorType.
const (
	OBSERVABILITYPIPELINEADDMETRICTAGSPROCESSORTYPE_ADD_METRIC_TAGS ObservabilityPipelineAddMetricTagsProcessorType = "add_metric_tags"
)

var allowedObservabilityPipelineAddMetricTagsProcessorTypeEnumValues = []ObservabilityPipelineAddMetricTagsProcessorType{
	OBSERVABILITYPIPELINEADDMETRICTAGSPROCESSORTYPE_ADD_METRIC_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAddMetricTagsProcessorType) GetAllowedValues() []ObservabilityPipelineAddMetricTagsProcessorType {
	return allowedObservabilityPipelineAddMetricTagsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAddMetricTagsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAddMetricTagsProcessorType(value)
	return nil
}

// NewObservabilityPipelineAddMetricTagsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineAddMetricTagsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAddMetricTagsProcessorTypeFromValue(v string) (*ObservabilityPipelineAddMetricTagsProcessorType, error) {
	ev := ObservabilityPipelineAddMetricTagsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAddMetricTagsProcessorType: valid values are %v", v, allowedObservabilityPipelineAddMetricTagsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAddMetricTagsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAddMetricTagsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAddMetricTagsProcessorType value.
func (v ObservabilityPipelineAddMetricTagsProcessorType) Ptr() *ObservabilityPipelineAddMetricTagsProcessorType {
	return &v
}
