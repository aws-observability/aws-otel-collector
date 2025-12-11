// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicAPIKeyUpdate The definition of the `AnthropicAPIKey` object.
type AnthropicAPIKeyUpdate struct {
	// The `AnthropicAPIKeyUpdate` `api_token`.
	ApiToken *string `json:"api_token,omitempty"`
	// The definition of the `AnthropicAPIKey` object.
	Type AnthropicAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnthropicAPIKeyUpdate instantiates a new AnthropicAPIKeyUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnthropicAPIKeyUpdate(typeVar AnthropicAPIKeyType) *AnthropicAPIKeyUpdate {
	this := AnthropicAPIKeyUpdate{}
	this.Type = typeVar
	return &this
}

// NewAnthropicAPIKeyUpdateWithDefaults instantiates a new AnthropicAPIKeyUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnthropicAPIKeyUpdateWithDefaults() *AnthropicAPIKeyUpdate {
	this := AnthropicAPIKeyUpdate{}
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *AnthropicAPIKeyUpdate) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnthropicAPIKeyUpdate) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *AnthropicAPIKeyUpdate) HasApiToken() bool {
	return o != nil && o.ApiToken != nil
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *AnthropicAPIKeyUpdate) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetType returns the Type field value.
func (o *AnthropicAPIKeyUpdate) GetType() AnthropicAPIKeyType {
	if o == nil {
		var ret AnthropicAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnthropicAPIKeyUpdate) GetTypeOk() (*AnthropicAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnthropicAPIKeyUpdate) SetType(v AnthropicAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnthropicAPIKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiToken != nil {
		toSerialize["api_token"] = o.ApiToken
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnthropicAPIKeyUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string              `json:"api_token,omitempty"`
		Type     *AnthropicAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.ApiToken = all.ApiToken
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
