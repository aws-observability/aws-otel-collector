// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthScopesRestrictionResponseAttributes Attributes of an OAuth2 client scopes restriction.
type OAuthScopesRestrictionResponseAttributes struct {
	// Permission scopes automatically required for this client (for example, mobile-app permission scopes).
	// Returns `null` when no scopes are required.
	RequiredPermissionScopes datadog.NullableList[string] `json:"required_permission_scopes"`
	// Allowlist of OIDC and permission scopes enforced for the OAuth2 client.
	ScopesRestriction NullableOAuthScopesRestriction `json:"scopes_restriction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuthScopesRestrictionResponseAttributes instantiates a new OAuthScopesRestrictionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuthScopesRestrictionResponseAttributes(requiredPermissionScopes datadog.NullableList[string], scopesRestriction NullableOAuthScopesRestriction) *OAuthScopesRestrictionResponseAttributes {
	this := OAuthScopesRestrictionResponseAttributes{}
	this.RequiredPermissionScopes = requiredPermissionScopes
	this.ScopesRestriction = scopesRestriction
	return &this
}

// NewOAuthScopesRestrictionResponseAttributesWithDefaults instantiates a new OAuthScopesRestrictionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuthScopesRestrictionResponseAttributesWithDefaults() *OAuthScopesRestrictionResponseAttributes {
	this := OAuthScopesRestrictionResponseAttributes{}
	return &this
}

// GetRequiredPermissionScopes returns the RequiredPermissionScopes field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *OAuthScopesRestrictionResponseAttributes) GetRequiredPermissionScopes() []string {
	if o == nil || o.RequiredPermissionScopes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.RequiredPermissionScopes.Get()
}

// GetRequiredPermissionScopesOk returns a tuple with the RequiredPermissionScopes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OAuthScopesRestrictionResponseAttributes) GetRequiredPermissionScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredPermissionScopes.Get(), o.RequiredPermissionScopes.IsSet()
}

// SetRequiredPermissionScopes sets field value.
func (o *OAuthScopesRestrictionResponseAttributes) SetRequiredPermissionScopes(v []string) {
	o.RequiredPermissionScopes.Set(&v)
}

// GetScopesRestriction returns the ScopesRestriction field value.
// If the value is explicit nil, the zero value for OAuthScopesRestriction will be returned.
func (o *OAuthScopesRestrictionResponseAttributes) GetScopesRestriction() OAuthScopesRestriction {
	if o == nil || o.ScopesRestriction.Get() == nil {
		var ret OAuthScopesRestriction
		return ret
	}
	return *o.ScopesRestriction.Get()
}

// GetScopesRestrictionOk returns a tuple with the ScopesRestriction field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OAuthScopesRestrictionResponseAttributes) GetScopesRestrictionOk() (*OAuthScopesRestriction, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopesRestriction.Get(), o.ScopesRestriction.IsSet()
}

// SetScopesRestriction sets field value.
func (o *OAuthScopesRestrictionResponseAttributes) SetScopesRestriction(v OAuthScopesRestriction) {
	o.ScopesRestriction.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuthScopesRestrictionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["required_permission_scopes"] = o.RequiredPermissionScopes.Get()
	toSerialize["scopes_restriction"] = o.ScopesRestriction.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuthScopesRestrictionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequiredPermissionScopes datadog.NullableList[string]   `json:"required_permission_scopes"`
		ScopesRestriction        NullableOAuthScopesRestriction `json:"scopes_restriction"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.RequiredPermissionScopes.IsSet() {
		return fmt.Errorf("required field required_permission_scopes missing")
	}
	if !all.ScopesRestriction.IsSet() {
		return fmt.Errorf("required field scopes_restriction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"required_permission_scopes", "scopes_restriction"})
	} else {
		return err
	}
	o.RequiredPermissionScopes = all.RequiredPermissionScopes
	o.ScopesRestriction = all.ScopesRestriction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
