// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaAPITokenUpdate The definition of the `OktaAPIToken` object.
type OktaAPITokenUpdate struct {
	// The `OktaAPITokenUpdate` `api_token`.
	ApiToken *string `json:"api_token,omitempty"`
	// The `OktaAPITokenUpdate` `domain`.
	Domain *string `json:"domain,omitempty"`
	// The definition of the `OktaAPIToken` object.
	Type OktaAPITokenType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOktaAPITokenUpdate instantiates a new OktaAPITokenUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOktaAPITokenUpdate(typeVar OktaAPITokenType) *OktaAPITokenUpdate {
	this := OktaAPITokenUpdate{}
	this.Type = typeVar
	return &this
}

// NewOktaAPITokenUpdateWithDefaults instantiates a new OktaAPITokenUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOktaAPITokenUpdateWithDefaults() *OktaAPITokenUpdate {
	this := OktaAPITokenUpdate{}
	return &this
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *OktaAPITokenUpdate) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAPITokenUpdate) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *OktaAPITokenUpdate) HasApiToken() bool {
	return o != nil && o.ApiToken != nil
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *OktaAPITokenUpdate) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *OktaAPITokenUpdate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAPITokenUpdate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *OktaAPITokenUpdate) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *OktaAPITokenUpdate) SetDomain(v string) {
	o.Domain = &v
}

// GetType returns the Type field value.
func (o *OktaAPITokenUpdate) GetType() OktaAPITokenType {
	if o == nil {
		var ret OktaAPITokenType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OktaAPITokenUpdate) GetTypeOk() (*OktaAPITokenType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OktaAPITokenUpdate) SetType(v OktaAPITokenType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OktaAPITokenUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiToken != nil {
		toSerialize["api_token"] = o.ApiToken
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OktaAPITokenUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiToken *string           `json:"api_token,omitempty"`
		Domain   *string           `json:"domain,omitempty"`
		Type     *OktaAPITokenType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_token", "domain", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiToken = all.ApiToken
	o.Domain = all.Domain
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
