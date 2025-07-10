// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineReduceProcessorType The processor type. The value should always be `reduce`.
type ObservabilityPipelineReduceProcessorType string

// List of ObservabilityPipelineReduceProcessorType.
const (
	OBSERVABILITYPIPELINEREDUCEPROCESSORTYPE_REDUCE ObservabilityPipelineReduceProcessorType = "reduce"
)

var allowedObservabilityPipelineReduceProcessorTypeEnumValues = []ObservabilityPipelineReduceProcessorType{
	OBSERVABILITYPIPELINEREDUCEPROCESSORTYPE_REDUCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineReduceProcessorType) GetAllowedValues() []ObservabilityPipelineReduceProcessorType {
	return allowedObservabilityPipelineReduceProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineReduceProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineReduceProcessorType(value)
	return nil
}

// NewObservabilityPipelineReduceProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineReduceProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineReduceProcessorTypeFromValue(v string) (*ObservabilityPipelineReduceProcessorType, error) {
	ev := ObservabilityPipelineReduceProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineReduceProcessorType: valid values are %v", v, allowedObservabilityPipelineReduceProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineReduceProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineReduceProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineReduceProcessorType value.
func (v ObservabilityPipelineReduceProcessorType) Ptr() *ObservabilityPipelineReduceProcessorType {
	return &v
}
