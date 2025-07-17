// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSampleProcessorType The processor type. The value should always be `sample`.
type ObservabilityPipelineSampleProcessorType string

// List of ObservabilityPipelineSampleProcessorType.
const (
	OBSERVABILITYPIPELINESAMPLEPROCESSORTYPE_SAMPLE ObservabilityPipelineSampleProcessorType = "sample"
)

var allowedObservabilityPipelineSampleProcessorTypeEnumValues = []ObservabilityPipelineSampleProcessorType{
	OBSERVABILITYPIPELINESAMPLEPROCESSORTYPE_SAMPLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSampleProcessorType) GetAllowedValues() []ObservabilityPipelineSampleProcessorType {
	return allowedObservabilityPipelineSampleProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSampleProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSampleProcessorType(value)
	return nil
}

// NewObservabilityPipelineSampleProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineSampleProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSampleProcessorTypeFromValue(v string) (*ObservabilityPipelineSampleProcessorType, error) {
	ev := ObservabilityPipelineSampleProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSampleProcessorType: valid values are %v", v, allowedObservabilityPipelineSampleProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSampleProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSampleProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSampleProcessorType value.
func (v ObservabilityPipelineSampleProcessorType) Ptr() *ObservabilityPipelineSampleProcessorType {
	return &v
}
