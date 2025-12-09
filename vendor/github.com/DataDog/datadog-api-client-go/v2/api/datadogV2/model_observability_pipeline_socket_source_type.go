// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceType The source type. The value should always be `socket`.
type ObservabilityPipelineSocketSourceType string

// List of ObservabilityPipelineSocketSourceType.
const (
	OBSERVABILITYPIPELINESOCKETSOURCETYPE_SOCKET ObservabilityPipelineSocketSourceType = "socket"
)

var allowedObservabilityPipelineSocketSourceTypeEnumValues = []ObservabilityPipelineSocketSourceType{
	OBSERVABILITYPIPELINESOCKETSOURCETYPE_SOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketSourceType) GetAllowedValues() []ObservabilityPipelineSocketSourceType {
	return allowedObservabilityPipelineSocketSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketSourceType(value)
	return nil
}

// NewObservabilityPipelineSocketSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineSocketSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketSourceTypeFromValue(v string) (*ObservabilityPipelineSocketSourceType, error) {
	ev := ObservabilityPipelineSocketSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketSourceType: valid values are %v", v, allowedObservabilityPipelineSocketSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketSourceType value.
func (v ObservabilityPipelineSocketSourceType) Ptr() *ObservabilityPipelineSocketSourceType {
	return &v
}
