// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthOidcScope OIDC scope a client may be restricted to.
type OAuthOidcScope string

// List of OAuthOidcScope.
const (
	OAUTHOIDCSCOPE_OPENID         OAuthOidcScope = "openid"
	OAUTHOIDCSCOPE_PROFILE        OAuthOidcScope = "profile"
	OAUTHOIDCSCOPE_EMAIL          OAuthOidcScope = "email"
	OAUTHOIDCSCOPE_OFFLINE_ACCESS OAuthOidcScope = "offline_access"
)

var allowedOAuthOidcScopeEnumValues = []OAuthOidcScope{
	OAUTHOIDCSCOPE_OPENID,
	OAUTHOIDCSCOPE_PROFILE,
	OAUTHOIDCSCOPE_EMAIL,
	OAUTHOIDCSCOPE_OFFLINE_ACCESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OAuthOidcScope) GetAllowedValues() []OAuthOidcScope {
	return allowedOAuthOidcScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OAuthOidcScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OAuthOidcScope(value)
	return nil
}

// NewOAuthOidcScopeFromValue returns a pointer to a valid OAuthOidcScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOAuthOidcScopeFromValue(v string) (*OAuthOidcScope, error) {
	ev := OAuthOidcScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OAuthOidcScope: valid values are %v", v, allowedOAuthOidcScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OAuthOidcScope) IsValid() bool {
	for _, existing := range allowedOAuthOidcScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuthOidcScope value.
func (v OAuthOidcScope) Ptr() *OAuthOidcScope {
	return &v
}
