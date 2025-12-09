// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyAPIKey The definition of the `LaunchDarklyAPIKey` object.
type LaunchDarklyAPIKey struct {
	// The `LaunchDarklyAPIKey` `api_token`.
	ApiToken string `json:"api_token"`
	// The definition of the `LaunchDarklyAPIKey` object.
	Type LaunchDarklyAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLaunchDarklyAPIKey instantiates a new LaunchDarklyAPIKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLaunchDarklyAPIKey(apiToken string, typeVar LaunchDarklyAPIKeyType) *LaunchDarklyAPIKey {
	this := LaunchDarklyAPIKey{}
	this.ApiToken = apiToken
	this.Type = typeVar
	return &this
}

// NewLaunchDarklyAPIKeyWithDefaults instantiates a new LaunchDarklyAPIKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLaunchDarklyAPIKeyWithDefaults() *LaunchDarklyAPIKey {
	this := LaunchDarklyAPIKey{}
	return &this
}

// GetApiToken returns the ApiToken field value.
func (o *LaunchDarklyAPIKey) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *LaunchDarklyAPIKey) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value.
func (o *LaunchDarklyAPIKey) SetApiToken(v string) {
	o.ApiToken = v
}

// GetType returns the Type field value.
func (o *LaunchDarklyAPIKey) GetType() LaunchDarklyAPIKeyType {
	if o == nil {
		var ret LaunchDarklyAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LaunchDarklyAPIKey) GetTypeOk() (*LaunchDarklyAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LaunchDarklyAPIKey) SetType(v LaunchDarklyAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LaunchDarklyAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_token"] = o.ApiToken
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LaunchDarklyAPIKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string                 `json:"api_token"`
		Type     *LaunchDarklyAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiToken == nil {
		return fmt.Errorf("required field api_token missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_token", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiToken = *all.ApiToken
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
