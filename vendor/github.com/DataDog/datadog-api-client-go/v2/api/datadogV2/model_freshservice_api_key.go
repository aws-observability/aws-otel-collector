// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceAPIKey The definition of the `FreshserviceAPIKey` object.
type FreshserviceAPIKey struct {
	// The `FreshserviceAPIKey` `api_key`.
	ApiKey string `json:"api_key"`
	// The `FreshserviceAPIKey` `domain`.
	Domain string `json:"domain"`
	// The definition of the `FreshserviceAPIKey` object.
	Type FreshserviceAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFreshserviceAPIKey instantiates a new FreshserviceAPIKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFreshserviceAPIKey(apiKey string, domain string, typeVar FreshserviceAPIKeyType) *FreshserviceAPIKey {
	this := FreshserviceAPIKey{}
	this.ApiKey = apiKey
	this.Domain = domain
	this.Type = typeVar
	return &this
}

// NewFreshserviceAPIKeyWithDefaults instantiates a new FreshserviceAPIKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFreshserviceAPIKeyWithDefaults() *FreshserviceAPIKey {
	this := FreshserviceAPIKey{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *FreshserviceAPIKey) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKey) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *FreshserviceAPIKey) SetApiKey(v string) {
	o.ApiKey = v
}

// GetDomain returns the Domain field value.
func (o *FreshserviceAPIKey) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKey) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *FreshserviceAPIKey) SetDomain(v string) {
	o.Domain = v
}

// GetType returns the Type field value.
func (o *FreshserviceAPIKey) GetType() FreshserviceAPIKeyType {
	if o == nil {
		var ret FreshserviceAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreshserviceAPIKey) GetTypeOk() (*FreshserviceAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FreshserviceAPIKey) SetType(v FreshserviceAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FreshserviceAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["domain"] = o.Domain
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FreshserviceAPIKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string                 `json:"api_key"`
		Domain *string                 `json:"domain"`
		Type   *FreshserviceAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
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
	o.ApiKey = *all.ApiKey
	o.Domain = *all.Domain
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
