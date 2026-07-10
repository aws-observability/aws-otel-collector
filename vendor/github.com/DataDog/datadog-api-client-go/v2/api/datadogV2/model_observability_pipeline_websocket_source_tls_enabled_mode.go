// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceTlsEnabledMode TLS mode. Must be `enabled`.
type ObservabilityPipelineWebsocketSourceTlsEnabledMode string

// List of ObservabilityPipelineWebsocketSourceTlsEnabledMode.
const (
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETLSENABLEDMODE_ENABLED ObservabilityPipelineWebsocketSourceTlsEnabledMode = "enabled"
)

var allowedObservabilityPipelineWebsocketSourceTlsEnabledModeEnumValues = []ObservabilityPipelineWebsocketSourceTlsEnabledMode{
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETLSENABLEDMODE_ENABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineWebsocketSourceTlsEnabledMode) GetAllowedValues() []ObservabilityPipelineWebsocketSourceTlsEnabledMode {
	return allowedObservabilityPipelineWebsocketSourceTlsEnabledModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineWebsocketSourceTlsEnabledMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineWebsocketSourceTlsEnabledMode(value)
	return nil
}

// NewObservabilityPipelineWebsocketSourceTlsEnabledModeFromValue returns a pointer to a valid ObservabilityPipelineWebsocketSourceTlsEnabledMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineWebsocketSourceTlsEnabledModeFromValue(v string) (*ObservabilityPipelineWebsocketSourceTlsEnabledMode, error) {
	ev := ObservabilityPipelineWebsocketSourceTlsEnabledMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineWebsocketSourceTlsEnabledMode: valid values are %v", v, allowedObservabilityPipelineWebsocketSourceTlsEnabledModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineWebsocketSourceTlsEnabledMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineWebsocketSourceTlsEnabledModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineWebsocketSourceTlsEnabledMode value.
func (v ObservabilityPipelineWebsocketSourceTlsEnabledMode) Ptr() *ObservabilityPipelineWebsocketSourceTlsEnabledMode {
	return &v
}
