// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationMode Protocol used to send logs.
type ObservabilityPipelineSocketDestinationMode string

// List of ObservabilityPipelineSocketDestinationMode.
const (
	OBSERVABILITYPIPELINESOCKETDESTINATIONMODE_TCP ObservabilityPipelineSocketDestinationMode = "tcp"
	OBSERVABILITYPIPELINESOCKETDESTINATIONMODE_UDP ObservabilityPipelineSocketDestinationMode = "udp"
)

var allowedObservabilityPipelineSocketDestinationModeEnumValues = []ObservabilityPipelineSocketDestinationMode{
	OBSERVABILITYPIPELINESOCKETDESTINATIONMODE_TCP,
	OBSERVABILITYPIPELINESOCKETDESTINATIONMODE_UDP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSocketDestinationMode) GetAllowedValues() []ObservabilityPipelineSocketDestinationMode {
	return allowedObservabilityPipelineSocketDestinationModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSocketDestinationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSocketDestinationMode(value)
	return nil
}

// NewObservabilityPipelineSocketDestinationModeFromValue returns a pointer to a valid ObservabilityPipelineSocketDestinationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSocketDestinationModeFromValue(v string) (*ObservabilityPipelineSocketDestinationMode, error) {
	ev := ObservabilityPipelineSocketDestinationMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSocketDestinationMode: valid values are %v", v, allowedObservabilityPipelineSocketDestinationModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSocketDestinationMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSocketDestinationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSocketDestinationMode value.
func (v ObservabilityPipelineSocketDestinationMode) Ptr() *ObservabilityPipelineSocketDestinationMode {
	return &v
}
