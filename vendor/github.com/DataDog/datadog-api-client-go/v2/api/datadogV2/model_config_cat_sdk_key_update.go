// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatSDKKeyUpdate The definition of the `ConfigCatSDKKey` object.
type ConfigCatSDKKeyUpdate struct {
	// The `ConfigCatSDKKeyUpdate` `api_password`.
	ApiPassword *string `json:"api_password,omitempty"`
	// The `ConfigCatSDKKeyUpdate` `api_username`.
	ApiUsername *string `json:"api_username,omitempty"`
	// The `ConfigCatSDKKeyUpdate` `sdk_key`.
	SdkKey *string `json:"sdk_key,omitempty"`
	// The definition of the `ConfigCatSDKKey` object.
	Type ConfigCatSDKKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfigCatSDKKeyUpdate instantiates a new ConfigCatSDKKeyUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigCatSDKKeyUpdate(typeVar ConfigCatSDKKeyType) *ConfigCatSDKKeyUpdate {
	this := ConfigCatSDKKeyUpdate{}
	this.Type = typeVar
	return &this
}

// NewConfigCatSDKKeyUpdateWithDefaults instantiates a new ConfigCatSDKKeyUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigCatSDKKeyUpdateWithDefaults() *ConfigCatSDKKeyUpdate {
	this := ConfigCatSDKKeyUpdate{}
	return &this
}

// GetApiPassword returns the ApiPassword field value if set, zero value otherwise.
func (o *ConfigCatSDKKeyUpdate) GetApiPassword() string {
	if o == nil || o.ApiPassword == nil {
		var ret string
		return ret
	}
	return *o.ApiPassword
}

// GetApiPasswordOk returns a tuple with the ApiPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKeyUpdate) GetApiPasswordOk() (*string, bool) {
	if o == nil || o.ApiPassword == nil {
		return nil, false
	}
	return o.ApiPassword, true
}

// HasApiPassword returns a boolean if a field has been set.
func (o *ConfigCatSDKKeyUpdate) HasApiPassword() bool {
	return o != nil && o.ApiPassword != nil
}

// SetApiPassword gets a reference to the given string and assigns it to the ApiPassword field.
func (o *ConfigCatSDKKeyUpdate) SetApiPassword(v string) {
	o.ApiPassword = &v
}

// GetApiUsername returns the ApiUsername field value if set, zero value otherwise.
func (o *ConfigCatSDKKeyUpdate) GetApiUsername() string {
	if o == nil || o.ApiUsername == nil {
		var ret string
		return ret
	}
	return *o.ApiUsername
}

// GetApiUsernameOk returns a tuple with the ApiUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKeyUpdate) GetApiUsernameOk() (*string, bool) {
	if o == nil || o.ApiUsername == nil {
		return nil, false
	}
	return o.ApiUsername, true
}

// HasApiUsername returns a boolean if a field has been set.
func (o *ConfigCatSDKKeyUpdate) HasApiUsername() bool {
	return o != nil && o.ApiUsername != nil
}

// SetApiUsername gets a reference to the given string and assigns it to the ApiUsername field.
func (o *ConfigCatSDKKeyUpdate) SetApiUsername(v string) {
	o.ApiUsername = &v
}

// GetSdkKey returns the SdkKey field value if set, zero value otherwise.
func (o *ConfigCatSDKKeyUpdate) GetSdkKey() string {
	if o == nil || o.SdkKey == nil {
		var ret string
		return ret
	}
	return *o.SdkKey
}

// GetSdkKeyOk returns a tuple with the SdkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKeyUpdate) GetSdkKeyOk() (*string, bool) {
	if o == nil || o.SdkKey == nil {
		return nil, false
	}
	return o.SdkKey, true
}

// HasSdkKey returns a boolean if a field has been set.
func (o *ConfigCatSDKKeyUpdate) HasSdkKey() bool {
	return o != nil && o.SdkKey != nil
}

// SetSdkKey gets a reference to the given string and assigns it to the SdkKey field.
func (o *ConfigCatSDKKeyUpdate) SetSdkKey(v string) {
	o.SdkKey = &v
}

// GetType returns the Type field value.
func (o *ConfigCatSDKKeyUpdate) GetType() ConfigCatSDKKeyType {
	if o == nil {
		var ret ConfigCatSDKKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigCatSDKKeyUpdate) GetTypeOk() (*ConfigCatSDKKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ConfigCatSDKKeyUpdate) SetType(v ConfigCatSDKKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfigCatSDKKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiPassword != nil {
		toSerialize["api_password"] = o.ApiPassword
	}
	if o.ApiUsername != nil {
		toSerialize["api_username"] = o.ApiUsername
	}
	if o.SdkKey != nil {
		toSerialize["sdk_key"] = o.SdkKey
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigCatSDKKeyUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiPassword *string              `json:"api_password,omitempty"`
		ApiUsername *string              `json:"api_username,omitempty"`
		SdkKey      *string              `json:"sdk_key,omitempty"`
		Type        *ConfigCatSDKKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.ApiPassword = all.ApiPassword
	o.ApiUsername = all.ApiUsername
	o.SdkKey = all.SdkKey
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
