// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyAPIKey The definition of the `FastlyAPIKey` object.
type FastlyAPIKey struct {
	// The `FastlyAPIKey` `api_key`.
	ApiKey string `json:"api_key"`
	// The definition of the `FastlyAPIKey` object.
	Type FastlyAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFastlyAPIKey instantiates a new FastlyAPIKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFastlyAPIKey(apiKey string, typeVar FastlyAPIKeyType) *FastlyAPIKey {
	this := FastlyAPIKey{}
	this.ApiKey = apiKey
	this.Type = typeVar
	return &this
}

// NewFastlyAPIKeyWithDefaults instantiates a new FastlyAPIKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFastlyAPIKeyWithDefaults() *FastlyAPIKey {
	this := FastlyAPIKey{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *FastlyAPIKey) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *FastlyAPIKey) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *FastlyAPIKey) SetApiKey(v string) {
	o.ApiKey = v
}

// GetType returns the Type field value.
func (o *FastlyAPIKey) GetType() FastlyAPIKeyType {
	if o == nil {
		var ret FastlyAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FastlyAPIKey) GetTypeOk() (*FastlyAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FastlyAPIKey) SetType(v FastlyAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FastlyAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FastlyAPIKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string           `json:"api_key"`
		Type   *FastlyAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = *all.ApiKey
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
