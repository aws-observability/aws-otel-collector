// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValidateV2Attributes Attributes of the API key validation response.
type ValidateV2Attributes struct {
	// The UUID of the API key.
	ApiKeyId string `json:"api_key_id"`
	// List of scope names associated with the API key.
	ApiKeyScopes []string `json:"api_key_scopes"`
	// Whether the API key is valid.
	Valid bool `json:"valid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewValidateV2Attributes instantiates a new ValidateV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewValidateV2Attributes(apiKeyId string, apiKeyScopes []string, valid bool) *ValidateV2Attributes {
	this := ValidateV2Attributes{}
	this.ApiKeyId = apiKeyId
	this.ApiKeyScopes = apiKeyScopes
	this.Valid = valid
	return &this
}

// NewValidateV2AttributesWithDefaults instantiates a new ValidateV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewValidateV2AttributesWithDefaults() *ValidateV2Attributes {
	this := ValidateV2Attributes{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value.
func (o *ValidateV2Attributes) GetApiKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value
// and a boolean to check if the value has been set.
func (o *ValidateV2Attributes) GetApiKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyId, true
}

// SetApiKeyId sets field value.
func (o *ValidateV2Attributes) SetApiKeyId(v string) {
	o.ApiKeyId = v
}

// GetApiKeyScopes returns the ApiKeyScopes field value.
func (o *ValidateV2Attributes) GetApiKeyScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ApiKeyScopes
}

// GetApiKeyScopesOk returns a tuple with the ApiKeyScopes field value
// and a boolean to check if the value has been set.
func (o *ValidateV2Attributes) GetApiKeyScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyScopes, true
}

// SetApiKeyScopes sets field value.
func (o *ValidateV2Attributes) SetApiKeyScopes(v []string) {
	o.ApiKeyScopes = v
}

// GetValid returns the Valid field value.
func (o *ValidateV2Attributes) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *ValidateV2Attributes) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value.
func (o *ValidateV2Attributes) SetValid(v bool) {
	o.Valid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ValidateV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key_id"] = o.ApiKeyId
	toSerialize["api_key_scopes"] = o.ApiKeyScopes
	toSerialize["valid"] = o.Valid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ValidateV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKeyId     *string   `json:"api_key_id"`
		ApiKeyScopes *[]string `json:"api_key_scopes"`
		Valid        *bool     `json:"valid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKeyId == nil {
		return fmt.Errorf("required field api_key_id missing")
	}
	if all.ApiKeyScopes == nil {
		return fmt.Errorf("required field api_key_scopes missing")
	}
	if all.Valid == nil {
		return fmt.Errorf("required field valid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key_id", "api_key_scopes", "valid"})
	} else {
		return err
	}
	o.ApiKeyId = *all.ApiKeyId
	o.ApiKeyScopes = *all.ApiKeyScopes
	o.Valid = *all.Valid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
