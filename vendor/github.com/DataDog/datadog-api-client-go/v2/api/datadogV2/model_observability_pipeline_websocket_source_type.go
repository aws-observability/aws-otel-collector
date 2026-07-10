// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceType The source type. The value should always be `websocket`.
type ObservabilityPipelineWebsocketSourceType string

// List of ObservabilityPipelineWebsocketSourceType.
const (
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETYPE_WEBSOCKET ObservabilityPipelineWebsocketSourceType = "websocket"
)

var allowedObservabilityPipelineWebsocketSourceTypeEnumValues = []ObservabilityPipelineWebsocketSourceType{
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETYPE_WEBSOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineWebsocketSourceType) GetAllowedValues() []ObservabilityPipelineWebsocketSourceType {
	return allowedObservabilityPipelineWebsocketSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineWebsocketSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineWebsocketSourceType(value)
	return nil
}

// NewObservabilityPipelineWebsocketSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineWebsocketSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineWebsocketSourceTypeFromValue(v string) (*ObservabilityPipelineWebsocketSourceType, error) {
	ev := ObservabilityPipelineWebsocketSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineWebsocketSourceType: valid values are %v", v, allowedObservabilityPipelineWebsocketSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineWebsocketSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineWebsocketSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineWebsocketSourceType value.
func (v ObservabilityPipelineWebsocketSourceType) Ptr() *ObservabilityPipelineWebsocketSourceType {
	return &v
}
