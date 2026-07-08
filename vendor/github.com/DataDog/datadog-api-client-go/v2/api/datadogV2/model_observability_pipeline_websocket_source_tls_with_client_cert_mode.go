// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceTlsWithClientCertMode TLS mode. Must be `with_client_cert`.
type ObservabilityPipelineWebsocketSourceTlsWithClientCertMode string

// List of ObservabilityPipelineWebsocketSourceTlsWithClientCertMode.
const (
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETLSWITHCLIENTCERTMODE_WITH_CLIENT_CERT ObservabilityPipelineWebsocketSourceTlsWithClientCertMode = "with_client_cert"
)

var allowedObservabilityPipelineWebsocketSourceTlsWithClientCertModeEnumValues = []ObservabilityPipelineWebsocketSourceTlsWithClientCertMode{
	OBSERVABILITYPIPELINEWEBSOCKETSOURCETLSWITHCLIENTCERTMODE_WITH_CLIENT_CERT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) GetAllowedValues() []ObservabilityPipelineWebsocketSourceTlsWithClientCertMode {
	return allowedObservabilityPipelineWebsocketSourceTlsWithClientCertModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineWebsocketSourceTlsWithClientCertMode(value)
	return nil
}

// NewObservabilityPipelineWebsocketSourceTlsWithClientCertModeFromValue returns a pointer to a valid ObservabilityPipelineWebsocketSourceTlsWithClientCertMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineWebsocketSourceTlsWithClientCertModeFromValue(v string) (*ObservabilityPipelineWebsocketSourceTlsWithClientCertMode, error) {
	ev := ObservabilityPipelineWebsocketSourceTlsWithClientCertMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineWebsocketSourceTlsWithClientCertMode: valid values are %v", v, allowedObservabilityPipelineWebsocketSourceTlsWithClientCertModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineWebsocketSourceTlsWithClientCertModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineWebsocketSourceTlsWithClientCertMode value.
func (v ObservabilityPipelineWebsocketSourceTlsWithClientCertMode) Ptr() *ObservabilityPipelineWebsocketSourceTlsWithClientCertMode {
	return &v
}
