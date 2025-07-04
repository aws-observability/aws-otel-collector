// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineFilterProcessorType The processor type. The value should always be `filter`.
type ObservabilityPipelineFilterProcessorType string

// List of ObservabilityPipelineFilterProcessorType.
const (
	OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER ObservabilityPipelineFilterProcessorType = "filter"
)

var allowedObservabilityPipelineFilterProcessorTypeEnumValues = []ObservabilityPipelineFilterProcessorType{
	OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineFilterProcessorType) GetAllowedValues() []ObservabilityPipelineFilterProcessorType {
	return allowedObservabilityPipelineFilterProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineFilterProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineFilterProcessorType(value)
	return nil
}

// NewObservabilityPipelineFilterProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineFilterProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineFilterProcessorTypeFromValue(v string) (*ObservabilityPipelineFilterProcessorType, error) {
	ev := ObservabilityPipelineFilterProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineFilterProcessorType: valid values are %v", v, allowedObservabilityPipelineFilterProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineFilterProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineFilterProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineFilterProcessorType value.
func (v ObservabilityPipelineFilterProcessorType) Ptr() *ObservabilityPipelineFilterProcessorType {
	return &v
}
