// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineRenameMetricTagsProcessorType The processor type. The value must be `rename_metric_tags`.
type ObservabilityPipelineRenameMetricTagsProcessorType string

// List of ObservabilityPipelineRenameMetricTagsProcessorType.
const (
	OBSERVABILITYPIPELINERENAMEMETRICTAGSPROCESSORTYPE_RENAME_METRIC_TAGS ObservabilityPipelineRenameMetricTagsProcessorType = "rename_metric_tags"
)

var allowedObservabilityPipelineRenameMetricTagsProcessorTypeEnumValues = []ObservabilityPipelineRenameMetricTagsProcessorType{
	OBSERVABILITYPIPELINERENAMEMETRICTAGSPROCESSORTYPE_RENAME_METRIC_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineRenameMetricTagsProcessorType) GetAllowedValues() []ObservabilityPipelineRenameMetricTagsProcessorType {
	return allowedObservabilityPipelineRenameMetricTagsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineRenameMetricTagsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineRenameMetricTagsProcessorType(value)
	return nil
}

// NewObservabilityPipelineRenameMetricTagsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineRenameMetricTagsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineRenameMetricTagsProcessorTypeFromValue(v string) (*ObservabilityPipelineRenameMetricTagsProcessorType, error) {
	ev := ObservabilityPipelineRenameMetricTagsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineRenameMetricTagsProcessorType: valid values are %v", v, allowedObservabilityPipelineRenameMetricTagsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineRenameMetricTagsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineRenameMetricTagsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineRenameMetricTagsProcessorType value.
func (v ObservabilityPipelineRenameMetricTagsProcessorType) Ptr() *ObservabilityPipelineRenameMetricTagsProcessorType {
	return &v
}
