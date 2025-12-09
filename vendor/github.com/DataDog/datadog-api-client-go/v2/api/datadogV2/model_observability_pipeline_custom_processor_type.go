// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCustomProcessorType The processor type. The value should always be `custom_processor`.
type ObservabilityPipelineCustomProcessorType string

// List of ObservabilityPipelineCustomProcessorType.
const (
	OBSERVABILITYPIPELINECUSTOMPROCESSORTYPE_CUSTOM_PROCESSOR ObservabilityPipelineCustomProcessorType = "custom_processor"
)

var allowedObservabilityPipelineCustomProcessorTypeEnumValues = []ObservabilityPipelineCustomProcessorType{
	OBSERVABILITYPIPELINECUSTOMPROCESSORTYPE_CUSTOM_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCustomProcessorType) GetAllowedValues() []ObservabilityPipelineCustomProcessorType {
	return allowedObservabilityPipelineCustomProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCustomProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCustomProcessorType(value)
	return nil
}

// NewObservabilityPipelineCustomProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineCustomProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCustomProcessorTypeFromValue(v string) (*ObservabilityPipelineCustomProcessorType, error) {
	ev := ObservabilityPipelineCustomProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCustomProcessorType: valid values are %v", v, allowedObservabilityPipelineCustomProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCustomProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCustomProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCustomProcessorType value.
func (v ObservabilityPipelineCustomProcessorType) Ptr() *ObservabilityPipelineCustomProcessorType {
	return &v
}
