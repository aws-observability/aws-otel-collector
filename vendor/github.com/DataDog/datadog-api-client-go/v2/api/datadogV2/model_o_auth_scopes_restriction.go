// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthScopesRestriction Allowlist of OIDC and permission scopes enforced for the OAuth2 client.
type OAuthScopesRestriction struct {
	// OIDC scopes the client is restricted to.
	OidcScopes []OAuthOidcScope `json:"oidc_scopes"`
	// Datadog permission scopes the client is restricted to.
	PermissionScopes []string `json:"permission_scopes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuthScopesRestriction instantiates a new OAuthScopesRestriction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuthScopesRestriction(oidcScopes []OAuthOidcScope, permissionScopes []string) *OAuthScopesRestriction {
	this := OAuthScopesRestriction{}
	this.OidcScopes = oidcScopes
	this.PermissionScopes = permissionScopes
	return &this
}

// NewOAuthScopesRestrictionWithDefaults instantiates a new OAuthScopesRestriction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuthScopesRestrictionWithDefaults() *OAuthScopesRestriction {
	this := OAuthScopesRestriction{}
	return &this
}

// GetOidcScopes returns the OidcScopes field value.
func (o *OAuthScopesRestriction) GetOidcScopes() []OAuthOidcScope {
	if o == nil {
		var ret []OAuthOidcScope
		return ret
	}
	return o.OidcScopes
}

// GetOidcScopesOk returns a tuple with the OidcScopes field value
// and a boolean to check if the value has been set.
func (o *OAuthScopesRestriction) GetOidcScopesOk() (*[]OAuthOidcScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OidcScopes, true
}

// SetOidcScopes sets field value.
func (o *OAuthScopesRestriction) SetOidcScopes(v []OAuthOidcScope) {
	o.OidcScopes = v
}

// GetPermissionScopes returns the PermissionScopes field value.
func (o *OAuthScopesRestriction) GetPermissionScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PermissionScopes
}

// GetPermissionScopesOk returns a tuple with the PermissionScopes field value
// and a boolean to check if the value has been set.
func (o *OAuthScopesRestriction) GetPermissionScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionScopes, true
}

// SetPermissionScopes sets field value.
func (o *OAuthScopesRestriction) SetPermissionScopes(v []string) {
	o.PermissionScopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuthScopesRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["oidc_scopes"] = o.OidcScopes
	toSerialize["permission_scopes"] = o.PermissionScopes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuthScopesRestriction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OidcScopes       *[]OAuthOidcScope `json:"oidc_scopes"`
		PermissionScopes *[]string         `json:"permission_scopes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OidcScopes == nil {
		return fmt.Errorf("required field oidc_scopes missing")
	}
	if all.PermissionScopes == nil {
		return fmt.Errorf("required field permission_scopes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oidc_scopes", "permission_scopes"})
	} else {
		return err
	}
	o.OidcScopes = *all.OidcScopes
	o.PermissionScopes = *all.PermissionScopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableOAuthScopesRestriction handles when a null is used for OAuthScopesRestriction.
type NullableOAuthScopesRestriction struct {
	value *OAuthScopesRestriction
	isSet bool
}

// Get returns the associated value.
func (v NullableOAuthScopesRestriction) Get() *OAuthScopesRestriction {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableOAuthScopesRestriction) Set(val *OAuthScopesRestriction) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableOAuthScopesRestriction) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableOAuthScopesRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOAuthScopesRestriction initializes the struct as if Set has been called.
func NewNullableOAuthScopesRestriction(val *OAuthScopesRestriction) *NullableOAuthScopesRestriction {
	return &NullableOAuthScopesRestriction{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableOAuthScopesRestriction) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableOAuthScopesRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
