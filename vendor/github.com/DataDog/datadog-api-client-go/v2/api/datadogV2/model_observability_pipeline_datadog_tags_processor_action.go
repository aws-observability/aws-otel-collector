// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogTagsProcessorAction The action to take on tags with matching keys.
type ObservabilityPipelineDatadogTagsProcessorAction string

// List of ObservabilityPipelineDatadogTagsProcessorAction.
const (
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORACTION_INCLUDE ObservabilityPipelineDatadogTagsProcessorAction = "include"
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORACTION_EXCLUDE ObservabilityPipelineDatadogTagsProcessorAction = "exclude"
)

var allowedObservabilityPipelineDatadogTagsProcessorActionEnumValues = []ObservabilityPipelineDatadogTagsProcessorAction{
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORACTION_INCLUDE,
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORACTION_EXCLUDE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogTagsProcessorAction) GetAllowedValues() []ObservabilityPipelineDatadogTagsProcessorAction {
	return allowedObservabilityPipelineDatadogTagsProcessorActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogTagsProcessorAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogTagsProcessorAction(value)
	return nil
}

// NewObservabilityPipelineDatadogTagsProcessorActionFromValue returns a pointer to a valid ObservabilityPipelineDatadogTagsProcessorAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogTagsProcessorActionFromValue(v string) (*ObservabilityPipelineDatadogTagsProcessorAction, error) {
	ev := ObservabilityPipelineDatadogTagsProcessorAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogTagsProcessorAction: valid values are %v", v, allowedObservabilityPipelineDatadogTagsProcessorActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogTagsProcessorAction) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogTagsProcessorActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogTagsProcessorAction value.
func (v ObservabilityPipelineDatadogTagsProcessorAction) Ptr() *ObservabilityPipelineDatadogTagsProcessorAction {
	return &v
}
