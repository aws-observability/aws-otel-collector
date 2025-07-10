// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDedupeProcessorType The processor type. The value should always be `dedupe`.
type ObservabilityPipelineDedupeProcessorType string

// List of ObservabilityPipelineDedupeProcessorType.
const (
	OBSERVABILITYPIPELINEDEDUPEPROCESSORTYPE_DEDUPE ObservabilityPipelineDedupeProcessorType = "dedupe"
)

var allowedObservabilityPipelineDedupeProcessorTypeEnumValues = []ObservabilityPipelineDedupeProcessorType{
	OBSERVABILITYPIPELINEDEDUPEPROCESSORTYPE_DEDUPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDedupeProcessorType) GetAllowedValues() []ObservabilityPipelineDedupeProcessorType {
	return allowedObservabilityPipelineDedupeProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDedupeProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDedupeProcessorType(value)
	return nil
}

// NewObservabilityPipelineDedupeProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineDedupeProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDedupeProcessorTypeFromValue(v string) (*ObservabilityPipelineDedupeProcessorType, error) {
	ev := ObservabilityPipelineDedupeProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDedupeProcessorType: valid values are %v", v, allowedObservabilityPipelineDedupeProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDedupeProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDedupeProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDedupeProcessorType value.
func (v ObservabilityPipelineDedupeProcessorType) Ptr() *ObservabilityPipelineDedupeProcessorType {
	return &v
}
