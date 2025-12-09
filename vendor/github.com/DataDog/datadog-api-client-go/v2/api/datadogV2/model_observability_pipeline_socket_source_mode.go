// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceMode Protocol used to receive logs.
type ObservabilityPipelineSocketSourceMode string

// List of ObservabilityPipelineSocketSourceMode.
const (
	OBSERVABILITYPIPELINESOCKETSOURCEMODE_TCP ObservabilityPipelineSocketSourceMode = "tcp"
	OBSERVABILITYPIPELINESOCKETSOURCEMODE_UDP ObservabilityPipelineSocketSourceMode = "udp"
)

var allowedObservabilityPipelineSocketSourceModeEnumValues = []ObservabilityPipelineSocketSourceMode{
	OBSERVABILITYPIPELINESOCKETSOURCEMODE_TCP,
	OBSERVABILITYPIPELINESOCKETSOURCEMODE_UDP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketSourceMode) GetAllowedValues() []ObservabilityPipelineSocketSourceMode {
	return allowedObservabilityPipelineSocketSourceModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketSourceMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketSourceMode(value)
	return nil
}

// NewObservabilityPipelineSocketSourceModeFromValue returns a pointer to a valid ObservabilityPipelineSocketSourceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketSourceModeFromValue(v string) (*ObservabilityPipelineSocketSourceMode, error) {
	ev := ObservabilityPipelineSocketSourceMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketSourceMode: valid values are %v", v, allowedObservabilityPipelineSocketSourceModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketSourceMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketSourceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketSourceMode value.
func (v ObservabilityPipelineSocketSourceMode) Ptr() *ObservabilityPipelineSocketSourceMode {
	return &v
}
