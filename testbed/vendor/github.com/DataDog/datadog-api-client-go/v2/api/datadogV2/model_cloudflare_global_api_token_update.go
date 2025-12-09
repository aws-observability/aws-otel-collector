// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareGlobalAPITokenUpdate The definition of the `CloudflareGlobalAPIToken` object.
type CloudflareGlobalAPITokenUpdate struct {
	// The `CloudflareGlobalAPITokenUpdate` `auth_email`.
	AuthEmail *string `json:"auth_email,omitempty"`
	// The `CloudflareGlobalAPITokenUpdate` `global_api_key`.
	GlobalApiKey *string `json:"global_api_key,omitempty"`
	// The definition of the `CloudflareGlobalAPIToken` object.
	Type CloudflareGlobalAPITokenType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudflareGlobalAPITokenUpdate instantiates a new CloudflareGlobalAPITokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudflareGlobalAPITokenUpdate(typeVar CloudflareGlobalAPITokenType) *CloudflareGlobalAPITokenUpdate {
	this := CloudflareGlobalAPITokenUpdate{}
	this.Type = typeVar
	return &this
}

// NewCloudflareGlobalAPITokenUpdateWithDefaults instantiates a new CloudflareGlobalAPITokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudflareGlobalAPITokenUpdateWithDefaults() *CloudflareGlobalAPITokenUpdate {
	this := CloudflareGlobalAPITokenUpdate{}
	return &this
}

// GetAuthEmail returns the AuthEmail field value if set, zero value otherwise.
func (o *CloudflareGlobalAPITokenUpdate) GetAuthEmail() string {
	if o == nil || o.AuthEmail == nil {
		var ret string
		return ret
	}
	return *o.AuthEmail
}

// GetAuthEmailOk returns a tuple with the AuthEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPITokenUpdate) GetAuthEmailOk() (*string, bool) {
	if o == nil || o.AuthEmail == nil {
		return nil, false
	}
	return o.AuthEmail, true
}

// HasAuthEmail returns a boolean if a field has been set.
func (o *CloudflareGlobalAPITokenUpdate) HasAuthEmail() bool {
	return o != nil && o.AuthEmail != nil
}

// SetAuthEmail gets a reference to the given string and assigns it to the AuthEmail field.
func (o *CloudflareGlobalAPITokenUpdate) SetAuthEmail(v string) {
	o.AuthEmail = &v
}

// GetGlobalApiKey returns the GlobalApiKey field value if set, zero value otherwise.
func (o *CloudflareGlobalAPITokenUpdate) GetGlobalApiKey() string {
	if o == nil || o.GlobalApiKey == nil {
		var ret string
		return ret
	}
	return *o.GlobalApiKey
}

// GetGlobalApiKeyOk returns a tuple with the GlobalApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPITokenUpdate) GetGlobalApiKeyOk() (*string, bool) {
	if o == nil || o.GlobalApiKey == nil {
		return nil, false
	}
	return o.GlobalApiKey, true
}

// HasGlobalApiKey returns a boolean if a field has been set.
func (o *CloudflareGlobalAPITokenUpdate) HasGlobalApiKey() bool {
	return o != nil && o.GlobalApiKey != nil
}

// SetGlobalApiKey gets a reference to the given string and assigns it to the GlobalApiKey field.
func (o *CloudflareGlobalAPITokenUpdate) SetGlobalApiKey(v string) {
	o.GlobalApiKey = &v
}

// GetType returns the Type field value.
func (o *CloudflareGlobalAPITokenUpdate) GetType() CloudflareGlobalAPITokenType {
	if o == nil {
		var ret CloudflareGlobalAPITokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPITokenUpdate) GetTypeOk() (*CloudflareGlobalAPITokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CloudflareGlobalAPITokenUpdate) SetType(v CloudflareGlobalAPITokenType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudflareGlobalAPITokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthEmail != nil {
		toSerialize["auth_email"] = o.AuthEmail
	}
	if o.GlobalApiKey != nil {
		toSerialize["global_api_key"] = o.GlobalApiKey
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudflareGlobalAPITokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthEmail    *string                       `json:"auth_email,omitempty"`
		GlobalApiKey *string                       `json:"global_api_key,omitempty"`
		Type         *CloudflareGlobalAPITokenType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_email", "global_api_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AuthEmail = all.AuthEmail
	o.GlobalApiKey = all.GlobalApiKey
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
