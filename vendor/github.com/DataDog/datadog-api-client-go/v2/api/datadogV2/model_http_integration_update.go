// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPIntegrationUpdate The definition of `HTTPIntegrationUpdate` object.
type HTTPIntegrationUpdate struct {
	// Base HTTP url for the integration
	BaseUrl *string `json:"base_url,omitempty"`
	// The definition of `HTTPCredentialsUpdate` object.
	Credentials *HTTPCredentialsUpdate `json:"credentials,omitempty"`
	// The definition of `HTTPIntegrationType` object.
	Type HTTPIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPIntegrationUpdate instantiates a new HTTPIntegrationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPIntegrationUpdate(typeVar HTTPIntegrationType) *HTTPIntegrationUpdate {
	this := HTTPIntegrationUpdate{}
	this.Type = typeVar
	return &this
}

// NewHTTPIntegrationUpdateWithDefaults instantiates a new HTTPIntegrationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPIntegrationUpdateWithDefaults() *HTTPIntegrationUpdate {
	this := HTTPIntegrationUpdate{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *HTTPIntegrationUpdate) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPIntegrationUpdate) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *HTTPIntegrationUpdate) HasBaseUrl() bool {
	return o != nil && o.BaseUrl != nil
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *HTTPIntegrationUpdate) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *HTTPIntegrationUpdate) GetCredentials() HTTPCredentialsUpdate {
	if o == nil || o.Credentials == nil {
		var ret HTTPCredentialsUpdate
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPIntegrationUpdate) GetCredentialsOk() (*HTTPCredentialsUpdate, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *HTTPIntegrationUpdate) HasCredentials() bool {
	return o != nil && o.Credentials != nil
}

// SetCredentials gets a reference to the given HTTPCredentialsUpdate and assigns it to the Credentials field.
func (o *HTTPIntegrationUpdate) SetCredentials(v HTTPCredentialsUpdate) {
	o.Credentials = &v
}

// GetType returns the Type field value.
func (o *HTTPIntegrationUpdate) GetType() HTTPIntegrationType {
	if o == nil {
		var ret HTTPIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPIntegrationUpdate) GetTypeOk() (*HTTPIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPIntegrationUpdate) SetType(v HTTPIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPIntegrationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
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
func (o *HTTPIntegrationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseUrl     *string                `json:"base_url,omitempty"`
		Credentials *HTTPCredentialsUpdate `json:"credentials,omitempty"`
		Type        *HTTPIntegrationType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"base_url", "credentials", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BaseUrl = all.BaseUrl
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
