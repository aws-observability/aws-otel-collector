// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertOAuthScopesRestrictionDataAttributes Attributes of an upsert OAuth2 scopes restriction request.
type UpsertOAuthScopesRestrictionDataAttributes struct {
	// OIDC scopes the client is allowed to request.
	OidcScopes []OAuthOidcScope `json:"oidc_scopes,omitempty"`
	// Datadog permission scopes the client is allowed to request.
	// Each value must be a valid permission name.
	PermissionScopes []string `json:"permission_scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertOAuthScopesRestrictionDataAttributes instantiates a new UpsertOAuthScopesRestrictionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertOAuthScopesRestrictionDataAttributes() *UpsertOAuthScopesRestrictionDataAttributes {
	this := UpsertOAuthScopesRestrictionDataAttributes{}
	return &this
}

// NewUpsertOAuthScopesRestrictionDataAttributesWithDefaults instantiates a new UpsertOAuthScopesRestrictionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertOAuthScopesRestrictionDataAttributesWithDefaults() *UpsertOAuthScopesRestrictionDataAttributes {
	this := UpsertOAuthScopesRestrictionDataAttributes{}
	return &this
}

// GetOidcScopes returns the OidcScopes field value if set, zero value otherwise.
func (o *UpsertOAuthScopesRestrictionDataAttributes) GetOidcScopes() []OAuthOidcScope {
	if o == nil || o.OidcScopes == nil {
		var ret []OAuthOidcScope
		return ret
	}
	return o.OidcScopes
}

// GetOidcScopesOk returns a tuple with the OidcScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertOAuthScopesRestrictionDataAttributes) GetOidcScopesOk() (*[]OAuthOidcScope, bool) {
	if o == nil || o.OidcScopes == nil {
		return nil, false
	}
	return &o.OidcScopes, true
}

// HasOidcScopes returns a boolean if a field has been set.
func (o *UpsertOAuthScopesRestrictionDataAttributes) HasOidcScopes() bool {
	return o != nil && o.OidcScopes != nil
}

// SetOidcScopes gets a reference to the given []OAuthOidcScope and assigns it to the OidcScopes field.
func (o *UpsertOAuthScopesRestrictionDataAttributes) SetOidcScopes(v []OAuthOidcScope) {
	o.OidcScopes = v
}

// GetPermissionScopes returns the PermissionScopes field value if set, zero value otherwise.
func (o *UpsertOAuthScopesRestrictionDataAttributes) GetPermissionScopes() []string {
	if o == nil || o.PermissionScopes == nil {
		var ret []string
		return ret
	}
	return o.PermissionScopes
}

// GetPermissionScopesOk returns a tuple with the PermissionScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertOAuthScopesRestrictionDataAttributes) GetPermissionScopesOk() (*[]string, bool) {
	if o == nil || o.PermissionScopes == nil {
		return nil, false
	}
	return &o.PermissionScopes, true
}

// HasPermissionScopes returns a boolean if a field has been set.
func (o *UpsertOAuthScopesRestrictionDataAttributes) HasPermissionScopes() bool {
	return o != nil && o.PermissionScopes != nil
}

// SetPermissionScopes gets a reference to the given []string and assigns it to the PermissionScopes field.
func (o *UpsertOAuthScopesRestrictionDataAttributes) SetPermissionScopes(v []string) {
	o.PermissionScopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertOAuthScopesRestrictionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OidcScopes != nil {
		toSerialize["oidc_scopes"] = o.OidcScopes
	}
	if o.PermissionScopes != nil {
		toSerialize["permission_scopes"] = o.PermissionScopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertOAuthScopesRestrictionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OidcScopes       []OAuthOidcScope `json:"oidc_scopes,omitempty"`
		PermissionScopes []string         `json:"permission_scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oidc_scopes", "permission_scopes"})
	} else {
		return err
	}
	o.OidcScopes = all.OidcScopes
	o.PermissionScopes = all.PermissionScopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
