// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareGlobalAPIToken The definition of the `CloudflareGlobalAPIToken` object.
type CloudflareGlobalAPIToken struct {
	// The `CloudflareGlobalAPIToken` `auth_email`.
	AuthEmail string `json:"auth_email"`
	// The `CloudflareGlobalAPIToken` `global_api_key`.
	GlobalApiKey string `json:"global_api_key"`
	// The definition of the `CloudflareGlobalAPIToken` object.
	Type CloudflareGlobalAPITokenType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudflareGlobalAPIToken instantiates a new CloudflareGlobalAPIToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudflareGlobalAPIToken(authEmail string, globalApiKey string, typeVar CloudflareGlobalAPITokenType) *CloudflareGlobalAPIToken {
	this := CloudflareGlobalAPIToken{}
	this.AuthEmail = authEmail
	this.GlobalApiKey = globalApiKey
	this.Type = typeVar
	return &this
}

// NewCloudflareGlobalAPITokenWithDefaults instantiates a new CloudflareGlobalAPIToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudflareGlobalAPITokenWithDefaults() *CloudflareGlobalAPIToken {
	this := CloudflareGlobalAPIToken{}
	return &this
}

// GetAuthEmail returns the AuthEmail field value.
func (o *CloudflareGlobalAPIToken) GetAuthEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthEmail
}

// GetAuthEmailOk returns a tuple with the AuthEmail field value
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPIToken) GetAuthEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthEmail, true
}

// SetAuthEmail sets field value.
func (o *CloudflareGlobalAPIToken) SetAuthEmail(v string) {
	o.AuthEmail = v
}

// GetGlobalApiKey returns the GlobalApiKey field value.
func (o *CloudflareGlobalAPIToken) GetGlobalApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GlobalApiKey
}

// GetGlobalApiKeyOk returns a tuple with the GlobalApiKey field value
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPIToken) GetGlobalApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalApiKey, true
}

// SetGlobalApiKey sets field value.
func (o *CloudflareGlobalAPIToken) SetGlobalApiKey(v string) {
	o.GlobalApiKey = v
}

// GetType returns the Type field value.
func (o *CloudflareGlobalAPIToken) GetType() CloudflareGlobalAPITokenType {
	if o == nil {
		var ret CloudflareGlobalAPITokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudflareGlobalAPIToken) GetTypeOk() (*CloudflareGlobalAPITokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CloudflareGlobalAPIToken) SetType(v CloudflareGlobalAPITokenType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudflareGlobalAPIToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_email"] = o.AuthEmail
	toSerialize["global_api_key"] = o.GlobalApiKey
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudflareGlobalAPIToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthEmail    *string                       `json:"auth_email"`
		GlobalApiKey *string                       `json:"global_api_key"`
		Type         *CloudflareGlobalAPITokenType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthEmail == nil {
		return fmt.Errorf("required field auth_email missing")
	}
	if all.GlobalApiKey == nil {
		return fmt.Errorf("required field global_api_key missing")
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
	o.AuthEmail = *all.AuthEmail
	o.GlobalApiKey = *all.GlobalApiKey
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
