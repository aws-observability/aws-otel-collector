// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineMetricTagsProcessorRuleAction The action to take on tags with matching keys.
type ObservabilityPipelineMetricTagsProcessorRuleAction string

// List of ObservabilityPipelineMetricTagsProcessorRuleAction.
const (
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEACTION_INCLUDE ObservabilityPipelineMetricTagsProcessorRuleAction = "include"
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEACTION_EXCLUDE ObservabilityPipelineMetricTagsProcessorRuleAction = "exclude"
)

var allowedObservabilityPipelineMetricTagsProcessorRuleActionEnumValues = []ObservabilityPipelineMetricTagsProcessorRuleAction{
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEACTION_INCLUDE,
	OBSERVABILITYPIPELINEMETRICTAGSPROCESSORRULEACTION_EXCLUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineMetricTagsProcessorRuleAction) GetAllowedValues() []ObservabilityPipelineMetricTagsProcessorRuleAction {
	return allowedObservabilityPipelineMetricTagsProcessorRuleActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineMetricTagsProcessorRuleAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineMetricTagsProcessorRuleAction(value)
	return nil
}

// NewObservabilityPipelineMetricTagsProcessorRuleActionFromValue returns a pointer to a valid ObservabilityPipelineMetricTagsProcessorRuleAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineMetricTagsProcessorRuleActionFromValue(v string) (*ObservabilityPipelineMetricTagsProcessorRuleAction, error) {
	ev := ObservabilityPipelineMetricTagsProcessorRuleAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineMetricTagsProcessorRuleAction: valid values are %v", v, allowedObservabilityPipelineMetricTagsProcessorRuleActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineMetricTagsProcessorRuleAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineMetricTagsProcessorRuleActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineMetricTagsProcessorRuleAction value.
func (v ObservabilityPipelineMetricTagsProcessorRuleAction) Ptr() *ObservabilityPipelineMetricTagsProcessorRuleAction {
	return &v
}
