// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyIntegration The definition of the `FastlyIntegration` object.
type FastlyIntegration struct {
	// The definition of the `FastlyCredentials` object.
	Credentials FastlyCredentials `json:"credentials"`
	// The definition of the `FastlyIntegrationType` object.
	Type FastlyIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFastlyIntegration instantiates a new FastlyIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFastlyIntegration(credentials FastlyCredentials, typeVar FastlyIntegrationType) *FastlyIntegration {
	this := FastlyIntegration{}
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewFastlyIntegrationWithDefaults instantiates a new FastlyIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFastlyIntegrationWithDefaults() *FastlyIntegration {
	this := FastlyIntegration{}
	return &this
}

// GetCredentials returns the Credentials field value.
func (o *FastlyIntegration) GetCredentials() FastlyCredentials {
	if o == nil {
		var ret FastlyCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *FastlyIntegration) GetCredentialsOk() (*FastlyCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *FastlyIntegration) SetCredentials(v FastlyCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *FastlyIntegration) GetType() FastlyIntegrationType {
	if o == nil {
		var ret FastlyIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FastlyIntegration) GetTypeOk() (*FastlyIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FastlyIntegration) SetType(v FastlyIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FastlyIntegration) MarshalJSON() ([]byte, error) {
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
func (o *FastlyIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *FastlyCredentials     `json:"credentials"`
		Type        *FastlyIntegrationType `json:"type"`
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
