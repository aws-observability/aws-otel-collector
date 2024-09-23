// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaAccountUpdateRequestAttributes Attributes object for updating an Okta account.
type OktaAccountUpdateRequestAttributes struct {
	// The API key of the Okta account.
	ApiKey *string `json:"api_key,omitempty"`
	// The authorization method for an Okta account.
	AuthMethod string `json:"auth_method"`
	// The Client ID of an Okta app integration.
	ClientId *string `json:"client_id,omitempty"`
	// The client secret of an Okta app integration.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The domain associated with an Okta account.
	Domain string `json:"domain"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOktaAccountUpdateRequestAttributes instantiates a new OktaAccountUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOktaAccountUpdateRequestAttributes(authMethod string, domain string) *OktaAccountUpdateRequestAttributes {
	this := OktaAccountUpdateRequestAttributes{}
	this.AuthMethod = authMethod
	this.Domain = domain
	return &this
}

// NewOktaAccountUpdateRequestAttributesWithDefaults instantiates a new OktaAccountUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOktaAccountUpdateRequestAttributesWithDefaults() *OktaAccountUpdateRequestAttributes {
	this := OktaAccountUpdateRequestAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *OktaAccountUpdateRequestAttributes) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountUpdateRequestAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *OktaAccountUpdateRequestAttributes) HasApiKey() bool {
	return o != nil && o.ApiKey != nil
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *OktaAccountUpdateRequestAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthMethod returns the AuthMethod field value.
func (o *OktaAccountUpdateRequestAttributes) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *OktaAccountUpdateRequestAttributes) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value.
func (o *OktaAccountUpdateRequestAttributes) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OktaAccountUpdateRequestAttributes) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountUpdateRequestAttributes) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OktaAccountUpdateRequestAttributes) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OktaAccountUpdateRequestAttributes) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OktaAccountUpdateRequestAttributes) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountUpdateRequestAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OktaAccountUpdateRequestAttributes) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OktaAccountUpdateRequestAttributes) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetDomain returns the Domain field value.
func (o *OktaAccountUpdateRequestAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *OktaAccountUpdateRequestAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *OktaAccountUpdateRequestAttributes) SetDomain(v string) {
	o.Domain = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OktaAccountUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	toSerialize["auth_method"] = o.AuthMethod
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	toSerialize["domain"] = o.Domain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OktaAccountUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey       *string `json:"api_key,omitempty"`
		AuthMethod   *string `json:"auth_method"`
		ClientId     *string `json:"client_id,omitempty"`
		ClientSecret *string `json:"client_secret,omitempty"`
		Domain       *string `json:"domain"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthMethod == nil {
		return fmt.Errorf("required field auth_method missing")
	}
	if all.Domain == nil {
		return fmt.Errorf("required field domain missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "auth_method", "client_id", "client_secret", "domain"})
	} else {
		return err
	}
	o.ApiKey = all.ApiKey
	o.AuthMethod = *all.AuthMethod
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret
	o.Domain = *all.Domain

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
