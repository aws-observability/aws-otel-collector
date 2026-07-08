// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceTls - TLS configuration for the WebSocket source. Use `enabled` for standard `wss://` connections, or `with_client_cert` to present a client certificate for mutual TLS.
type ObservabilityPipelineWebsocketSourceTls struct {
	ObservabilityPipelineWebsocketSourceTlsEnabled        *ObservabilityPipelineWebsocketSourceTlsEnabled
	ObservabilityPipelineWebsocketSourceTlsWithClientCert *ObservabilityPipelineWebsocketSourceTlsWithClientCert

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineWebsocketSourceTlsEnabledAsObservabilityPipelineWebsocketSourceTls is a convenience function that returns ObservabilityPipelineWebsocketSourceTlsEnabled wrapped in ObservabilityPipelineWebsocketSourceTls.
func ObservabilityPipelineWebsocketSourceTlsEnabledAsObservabilityPipelineWebsocketSourceTls(v *ObservabilityPipelineWebsocketSourceTlsEnabled) ObservabilityPipelineWebsocketSourceTls {
	return ObservabilityPipelineWebsocketSourceTls{ObservabilityPipelineWebsocketSourceTlsEnabled: v}
}

// ObservabilityPipelineWebsocketSourceTlsWithClientCertAsObservabilityPipelineWebsocketSourceTls is a convenience function that returns ObservabilityPipelineWebsocketSourceTlsWithClientCert wrapped in ObservabilityPipelineWebsocketSourceTls.
func ObservabilityPipelineWebsocketSourceTlsWithClientCertAsObservabilityPipelineWebsocketSourceTls(v *ObservabilityPipelineWebsocketSourceTlsWithClientCert) ObservabilityPipelineWebsocketSourceTls {
	return ObservabilityPipelineWebsocketSourceTls{ObservabilityPipelineWebsocketSourceTlsWithClientCert: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineWebsocketSourceTls) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineWebsocketSourceTlsEnabled
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineWebsocketSourceTlsEnabled)
	if err == nil {
		if obj.ObservabilityPipelineWebsocketSourceTlsEnabled != nil && obj.ObservabilityPipelineWebsocketSourceTlsEnabled.UnparsedObject == nil {
			jsonObservabilityPipelineWebsocketSourceTlsEnabled, _ := datadog.Marshal(obj.ObservabilityPipelineWebsocketSourceTlsEnabled)
			if string(jsonObservabilityPipelineWebsocketSourceTlsEnabled) == "{}" { // empty struct
				obj.ObservabilityPipelineWebsocketSourceTlsEnabled = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineWebsocketSourceTlsEnabled = nil
		}
	} else {
		obj.ObservabilityPipelineWebsocketSourceTlsEnabled = nil
	}

	// try to unmarshal data into ObservabilityPipelineWebsocketSourceTlsWithClientCert
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert)
	if err == nil {
		if obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert != nil && obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert.UnparsedObject == nil {
			jsonObservabilityPipelineWebsocketSourceTlsWithClientCert, _ := datadog.Marshal(obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert)
			if string(jsonObservabilityPipelineWebsocketSourceTlsWithClientCert) == "{}" { // empty struct
				obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert = nil
		}
	} else {
		obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineWebsocketSourceTlsEnabled = nil
		obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineWebsocketSourceTls) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineWebsocketSourceTlsEnabled != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineWebsocketSourceTlsEnabled)
	}

	if obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineWebsocketSourceTls) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineWebsocketSourceTlsEnabled != nil {
		return obj.ObservabilityPipelineWebsocketSourceTlsEnabled
	}

	if obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert != nil {
		return obj.ObservabilityPipelineWebsocketSourceTlsWithClientCert
	}

	// all schemas are nil
	return nil
}
