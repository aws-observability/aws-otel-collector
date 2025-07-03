// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineParseJSONProcessorType The processor type. The value should always be `parse_json`.
type ObservabilityPipelineParseJSONProcessorType string

// List of ObservabilityPipelineParseJSONProcessorType.
const (
	OBSERVABILITYPIPELINEPARSEJSONPROCESSORTYPE_PARSE_JSON ObservabilityPipelineParseJSONProcessorType = "parse_json"
)

var allowedObservabilityPipelineParseJSONProcessorTypeEnumValues = []ObservabilityPipelineParseJSONProcessorType{
	OBSERVABILITYPIPELINEPARSEJSONPROCESSORTYPE_PARSE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineParseJSONProcessorType) GetAllowedValues() []ObservabilityPipelineParseJSONProcessorType {
	return allowedObservabilityPipelineParseJSONProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineParseJSONProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineParseJSONProcessorType(value)
	return nil
}

// NewObservabilityPipelineParseJSONProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineParseJSONProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineParseJSONProcessorTypeFromValue(v string) (*ObservabilityPipelineParseJSONProcessorType, error) {
	ev := ObservabilityPipelineParseJSONProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineParseJSONProcessorType: valid values are %v", v, allowedObservabilityPipelineParseJSONProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineParseJSONProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineParseJSONProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineParseJSONProcessorType value.
func (v ObservabilityPipelineParseJSONProcessorType) Ptr() *ObservabilityPipelineParseJSONProcessorType {
	return &v
}
