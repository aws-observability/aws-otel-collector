// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthClientRegistrationResponse Response payload for a successful OAuth2 dynamic client registration as defined by RFC 7591.
type OAuthClientRegistrationResponse struct {
	// Unique identifier assigned to the registered client.
	ClientId uuid.UUID `json:"client_id"`
	// Human-readable name of the client.
	ClientName string `json:"client_name"`
	// OAuth 2.0 grant types registered for the client.
	GrantTypes []OAuthClientRegistrationGrantType `json:"grant_types"`
	// Redirection URIs registered for the client.
	RedirectUris []string `json:"redirect_uris"`
	// OAuth 2.0 response types registered for the client.
	ResponseTypes []OAuthClientRegistrationResponseType `json:"response_types"`
	// Authentication method registered for the token endpoint. Always `none`.
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuthClientRegistrationResponse instantiates a new OAuthClientRegistrationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuthClientRegistrationResponse(clientId uuid.UUID, clientName string, grantTypes []OAuthClientRegistrationGrantType, redirectUris []string, responseTypes []OAuthClientRegistrationResponseType, tokenEndpointAuthMethod string) *OAuthClientRegistrationResponse {
	this := OAuthClientRegistrationResponse{}
	this.ClientId = clientId
	this.ClientName = clientName
	this.GrantTypes = grantTypes
	this.RedirectUris = redirectUris
	this.ResponseTypes = responseTypes
	this.TokenEndpointAuthMethod = tokenEndpointAuthMethod
	return &this
}

// NewOAuthClientRegistrationResponseWithDefaults instantiates a new OAuthClientRegistrationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuthClientRegistrationResponseWithDefaults() *OAuthClientRegistrationResponse {
	this := OAuthClientRegistrationResponse{}
	return &this
}

// GetClientId returns the ClientId field value.
func (o *OAuthClientRegistrationResponse) GetClientId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetClientIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *OAuthClientRegistrationResponse) SetClientId(v uuid.UUID) {
	o.ClientId = v
}

// GetClientName returns the ClientName field value.
func (o *OAuthClientRegistrationResponse) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value.
func (o *OAuthClientRegistrationResponse) SetClientName(v string) {
	o.ClientName = v
}

// GetGrantTypes returns the GrantTypes field value.
func (o *OAuthClientRegistrationResponse) GetGrantTypes() []OAuthClientRegistrationGrantType {
	if o == nil {
		var ret []OAuthClientRegistrationGrantType
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetGrantTypesOk() (*[]OAuthClientRegistrationGrantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantTypes, true
}

// SetGrantTypes sets field value.
func (o *OAuthClientRegistrationResponse) SetGrantTypes(v []OAuthClientRegistrationGrantType) {
	o.GrantTypes = v
}

// GetRedirectUris returns the RedirectUris field value.
func (o *OAuthClientRegistrationResponse) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUris, true
}

// SetRedirectUris sets field value.
func (o *OAuthClientRegistrationResponse) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetResponseTypes returns the ResponseTypes field value.
func (o *OAuthClientRegistrationResponse) GetResponseTypes() []OAuthClientRegistrationResponseType {
	if o == nil {
		var ret []OAuthClientRegistrationResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetResponseTypesOk() (*[]OAuthClientRegistrationResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTypes, true
}

// SetResponseTypes sets field value.
func (o *OAuthClientRegistrationResponse) SetResponseTypes(v []OAuthClientRegistrationResponseType) {
	o.ResponseTypes = v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value.
func (o *OAuthClientRegistrationResponse) GetTokenEndpointAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value
// and a boolean to check if the value has been set.
func (o *OAuthClientRegistrationResponse) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpointAuthMethod, true
}

// SetTokenEndpointAuthMethod sets field value.
func (o *OAuthClientRegistrationResponse) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuthClientRegistrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_name"] = o.ClientName
	toSerialize["grant_types"] = o.GrantTypes
	toSerialize["redirect_uris"] = o.RedirectUris
	toSerialize["response_types"] = o.ResponseTypes
	toSerialize["token_endpoint_auth_method"] = o.TokenEndpointAuthMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuthClientRegistrationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId                *uuid.UUID                             `json:"client_id"`
		ClientName              *string                                `json:"client_name"`
		GrantTypes              *[]OAuthClientRegistrationGrantType    `json:"grant_types"`
		RedirectUris            *[]string                              `json:"redirect_uris"`
		ResponseTypes           *[]OAuthClientRegistrationResponseType `json:"response_types"`
		TokenEndpointAuthMethod *string                                `json:"token_endpoint_auth_method"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.ClientName == nil {
		return fmt.Errorf("required field client_name missing")
	}
	if all.GrantTypes == nil {
		return fmt.Errorf("required field grant_types missing")
	}
	if all.RedirectUris == nil {
		return fmt.Errorf("required field redirect_uris missing")
	}
	if all.ResponseTypes == nil {
		return fmt.Errorf("required field response_types missing")
	}
	if all.TokenEndpointAuthMethod == nil {
		return fmt.Errorf("required field token_endpoint_auth_method missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "client_name", "grant_types", "redirect_uris", "response_types", "token_endpoint_auth_method"})
	} else {
		return err
	}
	o.ClientId = *all.ClientId
	o.ClientName = *all.ClientName
	o.GrantTypes = *all.GrantTypes
	o.RedirectUris = *all.RedirectUris
	o.ResponseTypes = *all.ResponseTypes
	o.TokenEndpointAuthMethod = *all.TokenEndpointAuthMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
