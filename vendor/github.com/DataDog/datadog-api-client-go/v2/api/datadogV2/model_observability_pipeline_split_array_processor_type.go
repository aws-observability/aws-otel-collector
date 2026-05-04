// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplitArrayProcessorType The processor type. The value should always be `split_array`.
type ObservabilityPipelineSplitArrayProcessorType string

// List of ObservabilityPipelineSplitArrayProcessorType.
const (
	OBSERVABILITYPIPELINESPLITARRAYPROCESSORTYPE_SPLIT_ARRAY ObservabilityPipelineSplitArrayProcessorType = "split_array"
)

var allowedObservabilityPipelineSplitArrayProcessorTypeEnumValues = []ObservabilityPipelineSplitArrayProcessorType{
	OBSERVABILITYPIPELINESPLITARRAYPROCESSORTYPE_SPLIT_ARRAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplitArrayProcessorType) GetAllowedValues() []ObservabilityPipelineSplitArrayProcessorType {
	return allowedObservabilityPipelineSplitArrayProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplitArrayProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplitArrayProcessorType(value)
	return nil
}

// NewObservabilityPipelineSplitArrayProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineSplitArrayProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplitArrayProcessorTypeFromValue(v string) (*ObservabilityPipelineSplitArrayProcessorType, error) {
	ev := ObservabilityPipelineSplitArrayProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplitArrayProcessorType: valid values are %v", v, allowedObservabilityPipelineSplitArrayProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplitArrayProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplitArrayProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplitArrayProcessorType value.
func (v ObservabilityPipelineSplitArrayProcessorType) Ptr() *ObservabilityPipelineSplitArrayProcessorType {
	return &v
}
