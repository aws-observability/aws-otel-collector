// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSourceTlsEnabled TLS configuration that enables encryption without a client certificate. Use this for standard `wss://` connections that do not require mutual TLS.
type ObservabilityPipelineWebsocketSourceTlsEnabled struct {
	// TLS mode. Must be `enabled`.
	Mode ObservabilityPipelineWebsocketSourceTlsEnabledMode `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineWebsocketSourceTlsEnabled instantiates a new ObservabilityPipelineWebsocketSourceTlsEnabled object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineWebsocketSourceTlsEnabled(mode ObservabilityPipelineWebsocketSourceTlsEnabledMode) *ObservabilityPipelineWebsocketSourceTlsEnabled {
	this := ObservabilityPipelineWebsocketSourceTlsEnabled{}
	this.Mode = mode
	return &this
}

// NewObservabilityPipelineWebsocketSourceTlsEnabledWithDefaults instantiates a new ObservabilityPipelineWebsocketSourceTlsEnabled object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineWebsocketSourceTlsEnabledWithDefaults() *ObservabilityPipelineWebsocketSourceTlsEnabled {
	this := ObservabilityPipelineWebsocketSourceTlsEnabled{}
	return &this
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineWebsocketSourceTlsEnabled) GetMode() ObservabilityPipelineWebsocketSourceTlsEnabledMode {
	if o == nil {
		var ret ObservabilityPipelineWebsocketSourceTlsEnabledMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSourceTlsEnabled) GetModeOk() (*ObservabilityPipelineWebsocketSourceTlsEnabledMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineWebsocketSourceTlsEnabled) SetMode(v ObservabilityPipelineWebsocketSourceTlsEnabledMode) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineWebsocketSourceTlsEnabled) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineWebsocketSourceTlsEnabled) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode *ObservabilityPipelineWebsocketSourceTlsEnabledMode `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
