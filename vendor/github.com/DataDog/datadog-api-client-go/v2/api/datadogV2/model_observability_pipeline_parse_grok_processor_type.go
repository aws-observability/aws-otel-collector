// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseGrokProcessorType The processor type. The value should always be `parse_grok`.
type ObservabilityPipelineParseGrokProcessorType string

// List of ObservabilityPipelineParseGrokProcessorType.
const (
	OBSERVABILITYPIPELINEPARSEGROKPROCESSORTYPE_PARSE_GROK ObservabilityPipelineParseGrokProcessorType = "parse_grok"
)

var allowedObservabilityPipelineParseGrokProcessorTypeEnumValues = []ObservabilityPipelineParseGrokProcessorType{
	OBSERVABILITYPIPELINEPARSEGROKPROCESSORTYPE_PARSE_GROK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineParseGrokProcessorType) GetAllowedValues() []ObservabilityPipelineParseGrokProcessorType {
	return allowedObservabilityPipelineParseGrokProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineParseGrokProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineParseGrokProcessorType(value)
	return nil
}

// NewObservabilityPipelineParseGrokProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineParseGrokProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineParseGrokProcessorTypeFromValue(v string) (*ObservabilityPipelineParseGrokProcessorType, error) {
	ev := ObservabilityPipelineParseGrokProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineParseGrokProcessorType: valid values are %v", v, allowedObservabilityPipelineParseGrokProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineParseGrokProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineParseGrokProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineParseGrokProcessorType value.
func (v ObservabilityPipelineParseGrokProcessorType) Ptr() *ObservabilityPipelineParseGrokProcessorType {
	return &v
}
