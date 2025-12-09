// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogTagsProcessorType The processor type. The value should always be `datadog_tags`.
type ObservabilityPipelineDatadogTagsProcessorType string

// List of ObservabilityPipelineDatadogTagsProcessorType.
const (
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORTYPE_DATADOG_TAGS ObservabilityPipelineDatadogTagsProcessorType = "datadog_tags"
)

var allowedObservabilityPipelineDatadogTagsProcessorTypeEnumValues = []ObservabilityPipelineDatadogTagsProcessorType{
	OBSERVABILITYPIPELINEDATADOGTAGSPROCESSORTYPE_DATADOG_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogTagsProcessorType) GetAllowedValues() []ObservabilityPipelineDatadogTagsProcessorType {
	return allowedObservabilityPipelineDatadogTagsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogTagsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogTagsProcessorType(value)
	return nil
}

// NewObservabilityPipelineDatadogTagsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineDatadogTagsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogTagsProcessorTypeFromValue(v string) (*ObservabilityPipelineDatadogTagsProcessorType, error) {
	ev := ObservabilityPipelineDatadogTagsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogTagsProcessorType: valid values are %v", v, allowedObservabilityPipelineDatadogTagsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogTagsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogTagsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogTagsProcessorType value.
func (v ObservabilityPipelineDatadogTagsProcessorType) Ptr() *ObservabilityPipelineDatadogTagsProcessorType {
	return &v
}
