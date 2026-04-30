// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpentelemetrySourceType The source type. The value should always be `opentelemetry`.
type ObservabilityPipelineOpentelemetrySourceType string

// List of ObservabilityPipelineOpentelemetrySourceType.
const (
	OBSERVABILITYPIPELINEOPENTELEMETRYSOURCETYPE_OPENTELEMETRY ObservabilityPipelineOpentelemetrySourceType = "opentelemetry"
)

var allowedObservabilityPipelineOpentelemetrySourceTypeEnumValues = []ObservabilityPipelineOpentelemetrySourceType{
	OBSERVABILITYPIPELINEOPENTELEMETRYSOURCETYPE_OPENTELEMETRY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineOpentelemetrySourceType) GetAllowedValues() []ObservabilityPipelineOpentelemetrySourceType {
	return allowedObservabilityPipelineOpentelemetrySourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineOpentelemetrySourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineOpentelemetrySourceType(value)
	return nil
}

// NewObservabilityPipelineOpentelemetrySourceTypeFromValue returns a pointer to a valid ObservabilityPipelineOpentelemetrySourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineOpentelemetrySourceTypeFromValue(v string) (*ObservabilityPipelineOpentelemetrySourceType, error) {
	ev := ObservabilityPipelineOpentelemetrySourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineOpentelemetrySourceType: valid values are %v", v, allowedObservabilityPipelineOpentelemetrySourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineOpentelemetrySourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineOpentelemetrySourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineOpentelemetrySourceType value.
func (v ObservabilityPipelineOpentelemetrySourceType) Ptr() *ObservabilityPipelineOpentelemetrySourceType {
	return &v
}
