// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaAccountAttributes Attributes object for an Okta account.
type OktaAccountAttributes struct {
	// The API key of the Okta account.
	ApiKey *string `json:"api_key,omitempty"`
	// The authorization method for an Okta account.
	AuthMethod string `json:"auth_method"`
	// The Client ID of an Okta app integration.
	ClientId *string `json:"client_id,omitempty"`
	// The client secret of an Okta app integration.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The domain of the Okta account.
	Domain string `json:"domain"`
	// The name of the Okta account.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOktaAccountAttributes instantiates a new OktaAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOktaAccountAttributes(authMethod string, domain string, name string) *OktaAccountAttributes {
	this := OktaAccountAttributes{}
	this.AuthMethod = authMethod
	this.Domain = domain
	this.Name = name
	return &this
}

// NewOktaAccountAttributesWithDefaults instantiates a new OktaAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOktaAccountAttributesWithDefaults() *OktaAccountAttributes {
	this := OktaAccountAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *OktaAccountAttributes) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *OktaAccountAttributes) HasApiKey() bool {
	return o != nil && o.ApiKey != nil
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *OktaAccountAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAuthMethod returns the AuthMethod field value.
func (o *OktaAccountAttributes) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value.
func (o *OktaAccountAttributes) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OktaAccountAttributes) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OktaAccountAttributes) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OktaAccountAttributes) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OktaAccountAttributes) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OktaAccountAttributes) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OktaAccountAttributes) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetDomain returns the Domain field value.
func (o *OktaAccountAttributes) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value.
func (o *OktaAccountAttributes) SetDomain(v string) {
	o.Domain = v
}

// GetName returns the Name field value.
func (o *OktaAccountAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OktaAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OktaAccountAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OktaAccountAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OktaAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey       *string `json:"api_key,omitempty"`
		AuthMethod   *string `json:"auth_method"`
		ClientId     *string `json:"client_id,omitempty"`
		ClientSecret *string `json:"client_secret,omitempty"`
		Domain       *string `json:"domain"`
		Name         *string `json:"name"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "auth_method", "client_id", "client_secret", "domain", "name"})
	} else {
		return err
	}
	o.ApiKey = all.ApiKey
	o.AuthMethod = *all.AuthMethod
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret
	o.Domain = *all.Domain
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
