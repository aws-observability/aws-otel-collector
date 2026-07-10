// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsCreateAttributes OAuth2 client credentials attributes for a create request.
type WebhooksOAuth2ClientCredentialsCreateAttributes struct {
	// URL of the OAuth2 access token endpoint.
	AccessTokenUrl string `json:"access_token_url"`
	// The intended audience for the OAuth2 access token.
	Audience datadog.NullableString `json:"audience,omitempty"`
	// The OAuth2 client ID issued by the authorization server.
	ClientId string `json:"client_id"`
	// The OAuth2 client secret issued by the authorization server.
	// Write-only; never returned by the API.
	ClientSecret string `json:"client_secret"`
	// Human-readable name for this auth method. Must be unique within your organization.
	Name string `json:"name"`
	// Space-separated list of OAuth2 scopes to request.
	Scope datadog.NullableString `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsCreateAttributes instantiates a new WebhooksOAuth2ClientCredentialsCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsCreateAttributes(accessTokenUrl string, clientId string, clientSecret string, name string) *WebhooksOAuth2ClientCredentialsCreateAttributes {
	this := WebhooksOAuth2ClientCredentialsCreateAttributes{}
	this.AccessTokenUrl = accessTokenUrl
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Name = name
	return &this
}

// NewWebhooksOAuth2ClientCredentialsCreateAttributesWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsCreateAttributesWithDefaults() *WebhooksOAuth2ClientCredentialsCreateAttributes {
	this := WebhooksOAuth2ClientCredentialsCreateAttributes{}
	return &this
}

// GetAccessTokenUrl returns the AccessTokenUrl field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetAccessTokenUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessTokenUrl
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenUrl, true
}

// SetAccessTokenUrl sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl = v
}

// GetAudience returns the Audience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetAudience() string {
	if o == nil || o.Audience.Get() == nil {
		var ret string
		return ret
	}
	return *o.Audience.Get()
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audience.Get(), o.Audience.IsSet()
}

// HasAudience returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) HasAudience() bool {
	return o != nil && o.Audience.IsSet()
}

// SetAudience gets a reference to the given datadog.NullableString and assigns it to the Audience field.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetAudience(v string) {
	o.Audience.Set(&v)
}

// SetAudienceNil sets the value for Audience to be an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetAudienceNil() {
	o.Audience.Set(nil)
}

// UnsetAudience ensures that no value is present for Audience, not even an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) UnsetAudience() {
	o.Audience.Unset()
}

// GetClientId returns the ClientId field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetName returns the Name field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) HasScope() bool {
	return o != nil && o.Scope.IsSet()
}

// SetScope gets a reference to the given datadog.NullableString and assigns it to the Scope field.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetScope(v string) {
	o.Scope.Set(&v)
}

// SetScopeNil sets the value for Scope to be an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) UnsetScope() {
	o.Scope.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_token_url"] = o.AccessTokenUrl
	if o.Audience.IsSet() {
		toSerialize["audience"] = o.Audience.Get()
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["name"] = o.Name
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksOAuth2ClientCredentialsCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessTokenUrl *string                `json:"access_token_url"`
		Audience       datadog.NullableString `json:"audience,omitempty"`
		ClientId       *string                `json:"client_id"`
		ClientSecret   *string                `json:"client_secret"`
		Name           *string                `json:"name"`
		Scope          datadog.NullableString `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessTokenUrl == nil {
		return fmt.Errorf("required field access_token_url missing")
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.ClientSecret == nil {
		return fmt.Errorf("required field client_secret missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token_url", "audience", "client_id", "client_secret", "name", "scope"})
	} else {
		return err
	}
	o.AccessTokenUrl = *all.AccessTokenUrl
	o.Audience = all.Audience
	o.ClientId = *all.ClientId
	o.ClientSecret = *all.ClientSecret
	o.Name = *all.Name
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
