// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAddFieldsProcessorType The processor type. The value should always be `add_fields`.
type ObservabilityPipelineAddFieldsProcessorType string

// List of ObservabilityPipelineAddFieldsProcessorType.
const (
	OBSERVABILITYPIPELINEADDFIELDSPROCESSORTYPE_ADD_FIELDS ObservabilityPipelineAddFieldsProcessorType = "add_fields"
)

var allowedObservabilityPipelineAddFieldsProcessorTypeEnumValues = []ObservabilityPipelineAddFieldsProcessorType{
	OBSERVABILITYPIPELINEADDFIELDSPROCESSORTYPE_ADD_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAddFieldsProcessorType) GetAllowedValues() []ObservabilityPipelineAddFieldsProcessorType {
	return allowedObservabilityPipelineAddFieldsProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAddFieldsProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAddFieldsProcessorType(value)
	return nil
}

// NewObservabilityPipelineAddFieldsProcessorTypeFromValue returns a pointer to a valid ObservabilityPipelineAddFieldsProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAddFieldsProcessorTypeFromValue(v string) (*ObservabilityPipelineAddFieldsProcessorType, error) {
	ev := ObservabilityPipelineAddFieldsProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAddFieldsProcessorType: valid values are %v", v, allowedObservabilityPipelineAddFieldsProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAddFieldsProcessorType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAddFieldsProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAddFieldsProcessorType value.
func (v ObservabilityPipelineAddFieldsProcessorType) Ptr() *ObservabilityPipelineAddFieldsProcessorType {
	return &v
}
