// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineTagCardinalityLimitProcessorType The processor type. The value must be `tag_cardinality_limit`.
type ObservabilityPipelineTagCardinalityLimitProcessorType string

// List of ObservabilityPipelineTagCardinalityLimitProcessorType.
const (
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTYPE_TAG_CARDINALITY_LIMIT ObservabilityPipelineTagCardinalityLimitProcessorType = "tag_cardinality_limit"
)

var allowedObservabilityPipelineTagCardinalityLimitProcessorTypeEnumValues = []ObservabilityPipelineTagCardinalityLimitProcessorType{
	OBSERVABILITYPIPELINETAGCARDINALITYLIMITPROCESSORTYPE_TAG_CARDINALITY_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorType) GetAllowedValues() []ObservabilityPipelineTagCardinalityLimitProcessorType {
	return allowedObservabilityPipelineTagCardinalityLimitProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineTagCardinalityLimitProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineTagCardinalityLimitProcessorType(value)
	return nil
}

// NewObservabilityPipelineTagCardinalityLimitProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineTagCardinalityLimitProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineTagCardinalityLimitProcessorTypeFromValue(v string) (*ObservabilityPipelineTagCardinalityLimitProcessorType, error) {
	ev := ObservabilityPipelineTagCardinalityLimitProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineTagCardinalityLimitProcessorType: valid values are %v", v, allowedObservabilityPipelineTagCardinalityLimitProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineTagCardinalityLimitProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineTagCardinalityLimitProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineTagCardinalityLimitProcessorType value.
func (v ObservabilityPipelineTagCardinalityLimitProcessorType) Ptr() *ObservabilityPipelineTagCardinalityLimitProcessorType {
	return &v
}
