// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthClientRegistrationRequest Request payload for OAuth2 dynamic client registration as defined by RFC 7591.
type OAuthClientRegistrationRequest struct {
	// Human-readable name of the client. Control characters are rejected.
	ClientName string `json:"client_name"`
	// URL of the home page of the client.
	ClientUri *string `json:"client_uri,omitempty"`
	// OAuth 2.0 grant types the client may use.
	// Defaults to `authorization_code` and `refresh_token` when omitted.
	GrantTypes []OAuthClientRegistrationGrantType `json:"grant_types,omitempty"`
	// URL referencing the client's JSON Web Key Set.
	JwksUri *string `json:"jwks_uri,omitempty"`
	// URL referencing a logo for the client.
	LogoUri *string `json:"logo_uri,omitempty"`
	// URL pointing to the client's privacy policy.
	PolicyUri *string `json:"policy_uri,omitempty"`
	// Array of redirection URI strings used by the client in redirect-based flows.
	RedirectUris []string `json:"redirect_uris"`
	// OAuth 2.0 response types the client may use. Only `code` is supported.
	ResponseTypes []OAuthClientRegistrationResponseType `json:"response_types,omitempty"`
	// Space-separated list of scope values the client may request.
	Scope *string `json:"scope,omitempty"`
	// Requested authentication method for the token endpoint. Only `none` is supported.
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`
	// URL pointing to the client's terms of service.
	TosUri *string `json:"tos_uri,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuthClientRegistrationRequest instantiates a new OAuthClientRegistrationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuthClientRegistrationRequest(clientName string, redirectUris []string) *OAuthClientRegistrationRequest {
	this := OAuthClientRegistrationRequest{}
	this.ClientName = clientName
	this.RedirectUris = redirectUris
	return &this
}

// NewOAuthClientRegistrationRequestWithDefaults instantiates a new OAuthClientRegistrationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuthClientRegistrationRequestWithDefaults() *OAuthClientRegistrationRequest {
	this := OAuthClientRegistrationRequest{}
	return &this
}

// GetClientName returns the ClientName field value.
func (o *OAuthClientRegistrationRequest) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value.
func (o *OAuthClientRegistrationRequest) SetClientName(v string) {
	o.ClientName = v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetClientUri() string {
	if o == nil || o.ClientUri == nil {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetClientUriOk() (*string, bool) {
	if o == nil || o.ClientUri == nil {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasClientUri() bool {
	return o != nil && o.ClientUri != nil
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *OAuthClientRegistrationRequest) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetGrantTypes() []OAuthClientRegistrationGrantType {
	if o == nil || o.GrantTypes == nil {
		var ret []OAuthClientRegistrationGrantType
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetGrantTypesOk() (*[]OAuthClientRegistrationGrantType, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return &o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasGrantTypes() bool {
	return o != nil && o.GrantTypes != nil
}

// SetGrantTypes gets a reference to the given []OAuthClientRegistrationGrantType and assigns it to the GrantTypes field.
func (o *OAuthClientRegistrationRequest) SetGrantTypes(v []OAuthClientRegistrationGrantType) {
	o.GrantTypes = v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasJwksUri() bool {
	return o != nil && o.JwksUri != nil
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *OAuthClientRegistrationRequest) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetLogoUri() string {
	if o == nil || o.LogoUri == nil {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetLogoUriOk() (*string, bool) {
	if o == nil || o.LogoUri == nil {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasLogoUri() bool {
	return o != nil && o.LogoUri != nil
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *OAuthClientRegistrationRequest) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetPolicyUri() string {
	if o == nil || o.PolicyUri == nil {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetPolicyUriOk() (*string, bool) {
	if o == nil || o.PolicyUri == nil {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasPolicyUri() bool {
	return o != nil && o.PolicyUri != nil
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *OAuthClientRegistrationRequest) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetRedirectUris returns the RedirectUris field value.
func (o *OAuthClientRegistrationRequest) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUris, true
}

// SetRedirectUris sets field value.
func (o *OAuthClientRegistrationRequest) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetResponseTypes() []OAuthClientRegistrationResponseType {
	if o == nil || o.ResponseTypes == nil {
		var ret []OAuthClientRegistrationResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetResponseTypesOk() (*[]OAuthClientRegistrationResponseType, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return &o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasResponseTypes() bool {
	return o != nil && o.ResponseTypes != nil
}

// SetResponseTypes gets a reference to the given []OAuthClientRegistrationResponseType and assigns it to the ResponseTypes field.
func (o *OAuthClientRegistrationRequest) SetResponseTypes(v []OAuthClientRegistrationResponseType) {
	o.ResponseTypes = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OAuthClientRegistrationRequest) SetScope(v string) {
	o.Scope = &v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetTokenEndpointAuthMethod() string {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasTokenEndpointAuthMethod() bool {
	return o != nil && o.TokenEndpointAuthMethod != nil
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *OAuthClientRegistrationRequest) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *OAuthClientRegistrationRequest) GetTosUri() string {
	if o == nil || o.TosUri == nil {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationRequest) GetTosUriOk() (*string, bool) {
	if o == nil || o.TosUri == nil {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *OAuthClientRegistrationRequest) HasTosUri() bool {
	return o != nil && o.TosUri != nil
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *OAuthClientRegistrationRequest) SetTosUri(v string) {
	o.TosUri = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuthClientRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_name"] = o.ClientName
	if o.ClientUri != nil {
		toSerialize["client_uri"] = o.ClientUri
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.LogoUri != nil {
		toSerialize["logo_uri"] = o.LogoUri
	}
	if o.PolicyUri != nil {
		toSerialize["policy_uri"] = o.PolicyUri
	}
	toSerialize["redirect_uris"] = o.RedirectUris
	if o.ResponseTypes != nil {
		toSerialize["response_types"] = o.ResponseTypes
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenEndpointAuthMethod != nil {
		toSerialize["token_endpoint_auth_method"] = o.TokenEndpointAuthMethod
	}
	if o.TosUri != nil {
		toSerialize["tos_uri"] = o.TosUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuthClientRegistrationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientName              *string                               `json:"client_name"`
		ClientUri               *string                               `json:"client_uri,omitempty"`
		GrantTypes              []OAuthClientRegistrationGrantType    `json:"grant_types,omitempty"`
		JwksUri                 *string                               `json:"jwks_uri,omitempty"`
		LogoUri                 *string                               `json:"logo_uri,omitempty"`
		PolicyUri               *string                               `json:"policy_uri,omitempty"`
		RedirectUris            *[]string                             `json:"redirect_uris"`
		ResponseTypes           []OAuthClientRegistrationResponseType `json:"response_types,omitempty"`
		Scope                   *string                               `json:"scope,omitempty"`
		TokenEndpointAuthMethod *string                               `json:"token_endpoint_auth_method,omitempty"`
		TosUri                  *string                               `json:"tos_uri,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientName == nil {
		return fmt.Errorf("required field client_name missing")
	}
	if all.RedirectUris == nil {
		return fmt.Errorf("required field redirect_uris missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_name", "client_uri", "grant_types", "jwks_uri", "logo_uri", "policy_uri", "redirect_uris", "response_types", "scope", "token_endpoint_auth_method", "tos_uri"})
	} else {
		return err
	}
	o.ClientName = *all.ClientName
	o.ClientUri = all.ClientUri
	o.GrantTypes = all.GrantTypes
	o.JwksUri = all.JwksUri
	o.LogoUri = all.LogoUri
	o.PolicyUri = all.PolicyUri
	o.RedirectUris = *all.RedirectUris
	o.ResponseTypes = all.ResponseTypes
	o.Scope = all.Scope
	o.TokenEndpointAuthMethod = all.TokenEndpointAuthMethod
	o.TosUri = all.TosUri

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
