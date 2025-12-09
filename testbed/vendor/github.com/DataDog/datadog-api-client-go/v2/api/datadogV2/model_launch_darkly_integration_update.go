// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyIntegrationUpdate The definition of the `LaunchDarklyIntegrationUpdate` object.
type LaunchDarklyIntegrationUpdate struct {
	// The definition of the `LaunchDarklyCredentialsUpdate` object.
	Credentials *LaunchDarklyCredentialsUpdate `json:"credentials,omitempty"`
	// The definition of the `LaunchDarklyIntegrationType` object.
	Type LaunchDarklyIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLaunchDarklyIntegrationUpdate instantiates a new LaunchDarklyIntegrationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLaunchDarklyIntegrationUpdate(typeVar LaunchDarklyIntegrationType) *LaunchDarklyIntegrationUpdate {
	this := LaunchDarklyIntegrationUpdate{}
	this.Type = typeVar
	return &this
}

// NewLaunchDarklyIntegrationUpdateWithDefaults instantiates a new LaunchDarklyIntegrationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLaunchDarklyIntegrationUpdateWithDefaults() *LaunchDarklyIntegrationUpdate {
	this := LaunchDarklyIntegrationUpdate{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *LaunchDarklyIntegrationUpdate) GetCredentials() LaunchDarklyCredentialsUpdate {
	if o == nil || o.Credentials == nil {
		var ret LaunchDarklyCredentialsUpdate
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchDarklyIntegrationUpdate) GetCredentialsOk() (*LaunchDarklyCredentialsUpdate, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *LaunchDarklyIntegrationUpdate) HasCredentials() bool {
	return o != nil && o.Credentials != nil
}

// SetCredentials gets a reference to the given LaunchDarklyCredentialsUpdate and assigns it to the Credentials field.
func (o *LaunchDarklyIntegrationUpdate) SetCredentials(v LaunchDarklyCredentialsUpdate) {
	o.Credentials = &v
}

// GetType returns the Type field value.
func (o *LaunchDarklyIntegrationUpdate) GetType() LaunchDarklyIntegrationType {
	if o == nil {
		var ret LaunchDarklyIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LaunchDarklyIntegrationUpdate) GetTypeOk() (*LaunchDarklyIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LaunchDarklyIntegrationUpdate) SetType(v LaunchDarklyIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LaunchDarklyIntegrationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LaunchDarklyIntegrationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *LaunchDarklyCredentialsUpdate `json:"credentials,omitempty"`
		Type        *LaunchDarklyIntegrationType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Credentials = all.Credentials
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
