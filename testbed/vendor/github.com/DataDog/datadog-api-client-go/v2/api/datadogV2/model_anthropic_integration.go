// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicIntegration The definition of the `AnthropicIntegration` object.
type AnthropicIntegration struct {
	// The definition of the `AnthropicCredentials` object.
	Credentials AnthropicCredentials `json:"credentials"`
	// The definition of the `AnthropicIntegrationType` object.
	Type AnthropicIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnthropicIntegration instantiates a new AnthropicIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicIntegration(credentials AnthropicCredentials, typeVar AnthropicIntegrationType) *AnthropicIntegration {
	this := AnthropicIntegration{}
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewAnthropicIntegrationWithDefaults instantiates a new AnthropicIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicIntegrationWithDefaults() *AnthropicIntegration {
	this := AnthropicIntegration{}
	return &this
}

// GetCredentials returns the Credentials field value.
func (o *AnthropicIntegration) GetCredentials() AnthropicCredentials {
	if o == nil {
		var ret AnthropicCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *AnthropicIntegration) GetCredentialsOk() (*AnthropicCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *AnthropicIntegration) SetCredentials(v AnthropicCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *AnthropicIntegration) GetType() AnthropicIntegrationType {
	if o == nil {
		var ret AnthropicIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnthropicIntegration) GetTypeOk() (*AnthropicIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnthropicIntegration) SetType(v AnthropicIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["credentials"] = o.Credentials
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *AnthropicCredentials     `json:"credentials"`
		Type        *AnthropicIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Credentials == nil {
		return fmt.Errorf("required field credentials missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"credentials", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Credentials = *all.Credentials
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
