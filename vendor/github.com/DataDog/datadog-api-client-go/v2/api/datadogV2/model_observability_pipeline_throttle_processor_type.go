// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineThrottleProcessorType The processor type. The value should always be `throttle`.
type ObservabilityPipelineThrottleProcessorType string

// List of ObservabilityPipelineThrottleProcessorType.
const (
	OBSERVABILITYPIPELINETHROTTLEPROCESSORTYPE_THROTTLE ObservabilityPipelineThrottleProcessorType = "throttle"
)

var allowedObservabilityPipelineThrottleProcessorTypeEnumValues = []ObservabilityPipelineThrottleProcessorType{
	OBSERVABILITYPIPELINETHROTTLEPROCESSORTYPE_THROTTLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineThrottleProcessorType) GetAllowedValues() []ObservabilityPipelineThrottleProcessorType {
	return allowedObservabilityPipelineThrottleProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineThrottleProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineThrottleProcessorType(value)
	return nil
}

// NewObservabilityPipelineThrottleProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineThrottleProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineThrottleProcessorTypeFromValue(v string) (*ObservabilityPipelineThrottleProcessorType, error) {
	ev := ObservabilityPipelineThrottleProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineThrottleProcessorType: valid values are %v", v, allowedObservabilityPipelineThrottleProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineThrottleProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineThrottleProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineThrottleProcessorType value.
func (v ObservabilityPipelineThrottleProcessorType) Ptr() *ObservabilityPipelineThrottleProcessorType {
	return &v
}
