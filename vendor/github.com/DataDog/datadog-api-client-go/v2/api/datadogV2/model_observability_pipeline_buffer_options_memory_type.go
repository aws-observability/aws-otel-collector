// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineBufferOptionsMemoryType The type of the buffer that will be configured, a memory buffer.
type ObservabilityPipelineBufferOptionsMemoryType string

// List of ObservabilityPipelineBufferOptionsMemoryType.
const (
	OBSERVABILITYPIPELINEBUFFEROPTIONSMEMORYTYPE_MEMORY ObservabilityPipelineBufferOptionsMemoryType = "memory"
)

var allowedObservabilityPipelineBufferOptionsMemoryTypeEnumValues = []ObservabilityPipelineBufferOptionsMemoryType{
	OBSERVABILITYPIPELINEBUFFEROPTIONSMEMORYTYPE_MEMORY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineBufferOptionsMemoryType) GetAllowedValues() []ObservabilityPipelineBufferOptionsMemoryType {
	return allowedObservabilityPipelineBufferOptionsMemoryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineBufferOptionsMemoryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineBufferOptionsMemoryType(value)
	return nil
}

// NewObservabilityPipelineBufferOptionsMemoryTypeFromValue returns a pointer to a valid ObservabilityPipelineBufferOptionsMemoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineBufferOptionsMemoryTypeFromValue(v string) (*ObservabilityPipelineBufferOptionsMemoryType, error) {
	ev := ObservabilityPipelineBufferOptionsMemoryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineBufferOptionsMemoryType: valid values are %v", v, allowedObservabilityPipelineBufferOptionsMemoryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineBufferOptionsMemoryType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineBufferOptionsMemoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineBufferOptionsMemoryType value.
func (v ObservabilityPipelineBufferOptionsMemoryType) Ptr() *ObservabilityPipelineBufferOptionsMemoryType {
	return &v
}
