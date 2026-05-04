// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricTagsProcessorRuleMode The processing mode for tag filtering.
type ObservabilityPipelineMetricTagsProcessorRuleMode string

// List of ObservabilityPipelineMetricTagsProcessorRuleMode.
const (
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEMODE_FILTER ObservabilityPipelineMetricTagsProcessorRuleMode = "filter"
)

var allowedObservabilityPipelineMetricTagsProcessorRuleModeEnumValues = []ObservabilityPipelineMetricTagsProcessorRuleMode{
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEMODE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineMetricTagsProcessorRuleMode) GetAllowedValues() []ObservabilityPipelineMetricTagsProcessorRuleMode {
	return allowedObservabilityPipelineMetricTagsProcessorRuleModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineMetricTagsProcessorRuleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineMetricTagsProcessorRuleMode(value)
	return nil
}

// NewObservabilityPipelineMetricTagsProcessorRuleModeFromValue returns a pointer to a valid ObservabilityPipelineMetricTagsProcessorRuleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineMetricTagsProcessorRuleModeFromValue(v string) (*ObservabilityPipelineMetricTagsProcessorRuleMode, error) {
	ev := ObservabilityPipelineMetricTagsProcessorRuleMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineMetricTagsProcessorRuleMode: valid values are %v", v, allowedObservabilityPipelineMetricTagsProcessorRuleModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineMetricTagsProcessorRuleMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineMetricTagsProcessorRuleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineMetricTagsProcessorRuleMode value.
func (v ObservabilityPipelineMetricTagsProcessorRuleMode) Ptr() *ObservabilityPipelineMetricTagsProcessorRuleMode {
	return &v
}
