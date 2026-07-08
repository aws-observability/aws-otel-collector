// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceAuthStrategy Authentication strategy for the WebSocket source connection.
type ObservabilityPipelineWebsocketSourceAuthStrategy string

// List of ObservabilityPipelineWebsocketSourceAuthStrategy.
const (
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_NONE   ObservabilityPipelineWebsocketSourceAuthStrategy = "none"
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_BASIC  ObservabilityPipelineWebsocketSourceAuthStrategy = "basic"
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_BEARER ObservabilityPipelineWebsocketSourceAuthStrategy = "bearer"
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_CUSTOM ObservabilityPipelineWebsocketSourceAuthStrategy = "custom"
)

var allowedObservabilityPipelineWebsocketSourceAuthStrategyEnumValues = []ObservabilityPipelineWebsocketSourceAuthStrategy{
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_NONE,
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_BASIC,
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_BEARER,
	OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineWebsocketSourceAuthStrategy) GetAllowedValues() []ObservabilityPipelineWebsocketSourceAuthStrategy {
	return allowedObservabilityPipelineWebsocketSourceAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineWebsocketSourceAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineWebsocketSourceAuthStrategy(value)
	return nil
}

// NewObservabilityPipelineWebsocketSourceAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelineWebsocketSourceAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineWebsocketSourceAuthStrategyFromValue(v string) (*ObservabilityPipelineWebsocketSourceAuthStrategy, error) {
	ev := ObservabilityPipelineWebsocketSourceAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineWebsocketSourceAuthStrategy: valid values are %v", v, allowedObservabilityPipelineWebsocketSourceAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineWebsocketSourceAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineWebsocketSourceAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineWebsocketSourceAuthStrategy value.
func (v ObservabilityPipelineWebsocketSourceAuthStrategy) Ptr() *ObservabilityPipelineWebsocketSourceAuthStrategy {
	return &v
}
