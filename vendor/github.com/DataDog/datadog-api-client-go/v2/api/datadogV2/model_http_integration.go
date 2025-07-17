// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPIntegration The definition of `HTTPIntegration` object.
type HTTPIntegration struct {
	// Base HTTP url for the integration
	BaseUrl string `json:"base_url"`
	// The definition of `HTTPCredentials` object.
	Credentials HTTPCredentials `json:"credentials"`
	// The definition of `HTTPIntegrationType` object.
	Type HTTPIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPIntegration instantiates a new HTTPIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPIntegration(baseUrl string, credentials HTTPCredentials, typeVar HTTPIntegrationType) *HTTPIntegration {
	this := HTTPIntegration{}
	this.BaseUrl = baseUrl
	this.Credentials = credentials
	this.Type = typeVar
	return &this
}

// NewHTTPIntegrationWithDefaults instantiates a new HTTPIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPIntegrationWithDefaults() *HTTPIntegration {
	this := HTTPIntegration{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value.
func (o *HTTPIntegration) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *HTTPIntegration) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value.
func (o *HTTPIntegration) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetCredentials returns the Credentials field value.
func (o *HTTPIntegration) GetCredentials() HTTPCredentials {
	if o == nil {
		var ret HTTPCredentials
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *HTTPIntegration) GetCredentialsOk() (*HTTPCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value.
func (o *HTTPIntegration) SetCredentials(v HTTPCredentials) {
	o.Credentials = v
}

// GetType returns the Type field value.
func (o *HTTPIntegration) GetType() HTTPIntegrationType {
	if o == nil {
		var ret HTTPIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPIntegration) GetTypeOk() (*HTTPIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPIntegration) SetType(v HTTPIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_url"] = o.BaseUrl
	toSerialize["credentials"] = o.Credentials
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseUrl     *string              `json:"base_url"`
		Credentials *HTTPCredentials     `json:"credentials"`
		Type        *HTTPIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseUrl == nil {
		return fmt.Errorf("required field base_url missing")
	}
	if all.Credentials == nil {
		return fmt.Errorf("required field credentials missing")
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
	o.BaseUrl = *all.BaseUrl
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
