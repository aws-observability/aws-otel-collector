// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceAPIKeyUpdate The definition of the `FreshserviceAPIKey` object.
type FreshserviceAPIKeyUpdate struct {
	// The `FreshserviceAPIKeyUpdate` `api_key`.
	ApiKey *string `json:"api_key,omitempty"`
	// The `FreshserviceAPIKeyUpdate` `domain`.
	Domain *string `json:"domain,omitempty"`
	// The definition of the `FreshserviceAPIKey` object.
	Type FreshserviceAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFreshserviceAPIKeyUpdate instantiates a new FreshserviceAPIKeyUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFreshserviceAPIKeyUpdate(typeVar FreshserviceAPIKeyType) *FreshserviceAPIKeyUpdate {
	this := FreshserviceAPIKeyUpdate{}
	this.Type = typeVar
	return &this
}

// NewFreshserviceAPIKeyUpdateWithDefaults instantiates a new FreshserviceAPIKeyUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFreshserviceAPIKeyUpdateWithDefaults() *FreshserviceAPIKeyUpdate {
	this := FreshserviceAPIKeyUpdate{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *FreshserviceAPIKeyUpdate) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKeyUpdate) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *FreshserviceAPIKeyUpdate) HasApiKey() bool {
	return o != nil && o.ApiKey != nil
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *FreshserviceAPIKeyUpdate) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *FreshserviceAPIKeyUpdate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKeyUpdate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *FreshserviceAPIKeyUpdate) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *FreshserviceAPIKeyUpdate) SetDomain(v string) {
	o.Domain = &v
}

// GetType returns the Type field value.
func (o *FreshserviceAPIKeyUpdate) GetType() FreshserviceAPIKeyType {
	if o == nil {
		var ret FreshserviceAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKeyUpdate) GetTypeOk() (*FreshserviceAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FreshserviceAPIKeyUpdate) SetType(v FreshserviceAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FreshserviceAPIKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
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
func (o *FreshserviceAPIKeyUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string                 `json:"api_key,omitempty"`
		Domain *string                 `json:"domain,omitempty"`
		Type   *FreshserviceAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "domain", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = all.ApiKey
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
