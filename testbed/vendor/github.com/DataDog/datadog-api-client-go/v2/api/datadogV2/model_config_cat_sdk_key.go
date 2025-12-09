// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatSDKKey The definition of the `ConfigCatSDKKey` object.
type ConfigCatSDKKey struct {
	// The `ConfigCatSDKKey` `api_password`.
	ApiPassword string `json:"api_password"`
	// The `ConfigCatSDKKey` `api_username`.
	ApiUsername string `json:"api_username"`
	// The `ConfigCatSDKKey` `sdk_key`.
	SdkKey string `json:"sdk_key"`
	// The definition of the `ConfigCatSDKKey` object.
	Type ConfigCatSDKKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfigCatSDKKey instantiates a new ConfigCatSDKKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigCatSDKKey(apiPassword string, apiUsername string, sdkKey string, typeVar ConfigCatSDKKeyType) *ConfigCatSDKKey {
	this := ConfigCatSDKKey{}
	this.ApiPassword = apiPassword
	this.ApiUsername = apiUsername
	this.SdkKey = sdkKey
	this.Type = typeVar
	return &this
}

// NewConfigCatSDKKeyWithDefaults instantiates a new ConfigCatSDKKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigCatSDKKeyWithDefaults() *ConfigCatSDKKey {
	this := ConfigCatSDKKey{}
	return &this
}

// GetApiPassword returns the ApiPassword field value.
func (o *ConfigCatSDKKey) GetApiPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiPassword
}

// GetApiPasswordOk returns a tuple with the ApiPassword field value
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKey) GetApiPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiPassword, true
}

// SetApiPassword sets field value.
func (o *ConfigCatSDKKey) SetApiPassword(v string) {
	o.ApiPassword = v
}

// GetApiUsername returns the ApiUsername field value.
func (o *ConfigCatSDKKey) GetApiUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiUsername
}

// GetApiUsernameOk returns a tuple with the ApiUsername field value
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKey) GetApiUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiUsername, true
}

// SetApiUsername sets field value.
func (o *ConfigCatSDKKey) SetApiUsername(v string) {
	o.ApiUsername = v
}

// GetSdkKey returns the SdkKey field value.
func (o *ConfigCatSDKKey) GetSdkKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SdkKey
}

// GetSdkKeyOk returns a tuple with the SdkKey field value
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKey) GetSdkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkKey, true
}

// SetSdkKey sets field value.
func (o *ConfigCatSDKKey) SetSdkKey(v string) {
	o.SdkKey = v
}

// GetType returns the Type field value.
func (o *ConfigCatSDKKey) GetType() ConfigCatSDKKeyType {
	if o == nil {
		var ret ConfigCatSDKKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKey) GetTypeOk() (*ConfigCatSDKKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ConfigCatSDKKey) SetType(v ConfigCatSDKKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfigCatSDKKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_password"] = o.ApiPassword
	toSerialize["api_username"] = o.ApiUsername
	toSerialize["sdk_key"] = o.SdkKey
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigCatSDKKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiPassword *string              `json:"api_password"`
		ApiUsername *string              `json:"api_username"`
		SdkKey      *string              `json:"sdk_key"`
		Type        *ConfigCatSDKKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiPassword == nil {
		return fmt.Errorf("required field api_password missing")
	}
	if all.ApiUsername == nil {
		return fmt.Errorf("required field api_username missing")
	}
	if all.SdkKey == nil {
		return fmt.Errorf("required field sdk_key missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_password", "api_username", "sdk_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiPassword = *all.ApiPassword
	o.ApiUsername = *all.ApiUsername
	o.SdkKey = *all.SdkKey
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
