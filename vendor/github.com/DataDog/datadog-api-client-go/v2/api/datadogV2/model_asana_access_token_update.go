// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AsanaAccessTokenUpdate The definition of the `AsanaAccessToken` object.
type AsanaAccessTokenUpdate struct {
	// The `AsanaAccessTokenUpdate` `access_token`.
	AccessToken *string `json:"access_token,omitempty"`
	// The definition of the `AsanaAccessToken` object.
	Type AsanaAccessTokenType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAsanaAccessTokenUpdate instantiates a new AsanaAccessTokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAsanaAccessTokenUpdate(typeVar AsanaAccessTokenType) *AsanaAccessTokenUpdate {
	this := AsanaAccessTokenUpdate{}
	this.Type = typeVar
	return &this
}

// NewAsanaAccessTokenUpdateWithDefaults instantiates a new AsanaAccessTokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAsanaAccessTokenUpdateWithDefaults() *AsanaAccessTokenUpdate {
	this := AsanaAccessTokenUpdate{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AsanaAccessTokenUpdate) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsanaAccessTokenUpdate) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AsanaAccessTokenUpdate) HasAccessToken() bool {
	return o != nil && o.AccessToken != nil
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AsanaAccessTokenUpdate) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetType returns the Type field value.
func (o *AsanaAccessTokenUpdate) GetType() AsanaAccessTokenType {
	if o == nil {
		var ret AsanaAccessTokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AsanaAccessTokenUpdate) GetTypeOk() (*AsanaAccessTokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AsanaAccessTokenUpdate) SetType(v AsanaAccessTokenType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AsanaAccessTokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AsanaAccessTokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessToken *string               `json:"access_token,omitempty"`
		Type        *AsanaAccessTokenType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessToken = all.AccessToken
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
