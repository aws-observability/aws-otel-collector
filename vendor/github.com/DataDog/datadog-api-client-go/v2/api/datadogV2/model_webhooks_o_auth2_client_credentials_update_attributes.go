// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsUpdateAttributes OAuth2 client credentials attributes for an update request.
type WebhooksOAuth2ClientCredentialsUpdateAttributes struct {
	// URL of the OAuth2 access token endpoint.
	AccessTokenUrl *string `json:"access_token_url,omitempty"`
	// The intended audience for the OAuth2 access token.
	Audience datadog.NullableString `json:"audience,omitempty"`
	// The OAuth2 client ID issued by the authorization server.
	ClientId *string `json:"client_id,omitempty"`
	// The OAuth2 client secret issued by the authorization server.
	// Write-only; never returned by the API.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Human-readable name for this auth method.
	Name *string `json:"name,omitempty"`
	// Space-separated list of OAuth2 scopes to request.
	Scope datadog.NullableString `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsUpdateAttributes instantiates a new WebhooksOAuth2ClientCredentialsUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsUpdateAttributes() *WebhooksOAuth2ClientCredentialsUpdateAttributes {
	this := WebhooksOAuth2ClientCredentialsUpdateAttributes{}
	return &this
}

// NewWebhooksOAuth2ClientCredentialsUpdateAttributesWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsUpdateAttributesWithDefaults() *WebhooksOAuth2ClientCredentialsUpdateAttributes {
	this := WebhooksOAuth2ClientCredentialsUpdateAttributes{}
	return &this
}

// GetAccessTokenUrl returns the AccessTokenUrl field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetAccessTokenUrl() string {
	if o == nil || o.AccessTokenUrl == nil {
		var ret string
		return ret
	}
	return *o.AccessTokenUrl
}

// GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetAccessTokenUrlOk() (*string, bool) {
	if o == nil || o.AccessTokenUrl == nil {
		return nil, false
	}
	return o.AccessTokenUrl, true
}

// HasAccessTokenUrl returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasAccessTokenUrl() bool {
	return o != nil && o.AccessTokenUrl != nil
}

// SetAccessTokenUrl gets a reference to the given string and assigns it to the AccessTokenUrl field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetAccessTokenUrl(v string) {
	o.AccessTokenUrl = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetAudience() string {
	if o == nil || o.Audience.Get() == nil {
		var ret string
		return ret
	}
	return *o.Audience.Get()
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audience.Get(), o.Audience.IsSet()
}

// HasAudience returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasAudience() bool {
	return o != nil && o.Audience.IsSet()
}

// SetAudience gets a reference to the given datadog.NullableString and assigns it to the Audience field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetAudience(v string) {
	o.Audience.Set(&v)
}

// SetAudienceNil sets the value for Audience to be an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetAudienceNil() {
	o.Audience.Set(nil)
}

// UnsetAudience ensures that no value is present for Audience, not even an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) UnsetAudience() {
	o.Audience.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasClientId() bool {
	return o != nil && o.ClientId != nil
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasClientSecret() bool {
	return o != nil && o.ClientSecret != nil
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) HasScope() bool {
	return o != nil && o.Scope.IsSet()
}

// SetScope gets a reference to the given datadog.NullableString and assigns it to the Scope field.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetScope(v string) {
	o.Scope.Set(&v)
}

// SetScopeNil sets the value for Scope to be an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) UnsetScope() {
	o.Scope.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessTokenUrl != nil {
		toSerialize["access_token_url"] = o.AccessTokenUrl
	}
	if o.Audience.IsSet() {
		toSerialize["audience"] = o.Audience.Get()
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksOAuth2ClientCredentialsUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessTokenUrl *string                `json:"access_token_url,omitempty"`
		Audience       datadog.NullableString `json:"audience,omitempty"`
		ClientId       *string                `json:"client_id,omitempty"`
		ClientSecret   *string                `json:"client_secret,omitempty"`
		Name           *string                `json:"name,omitempty"`
		Scope          datadog.NullableString `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token_url", "audience", "client_id", "client_secret", "name", "scope"})
	} else {
		return err
	}
	o.AccessTokenUrl = all.AccessTokenUrl
	o.Audience = all.Audience
	o.ClientId = all.ClientId
	o.ClientSecret = all.ClientSecret
	o.Name = all.Name
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
