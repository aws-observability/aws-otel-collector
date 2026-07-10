// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthScopesRestrictionType JSON:API resource type for an OAuth2 client scopes restriction.
type OAuthScopesRestrictionType string

// List of OAuthScopesRestrictionType.
const (
	OAUTHSCOPESRESTRICTIONTYPE_SCOPES_RESTRICTION OAuthScopesRestrictionType = "scopes_restriction"
)

var allowedOAuthScopesRestrictionTypeEnumValues = []OAuthScopesRestrictionType{
	OAUTHSCOPESRESTRICTIONTYPE_SCOPES_RESTRICTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OAuthScopesRestrictionType) GetAllowedValues() []OAuthScopesRestrictionType {
	return allowedOAuthScopesRestrictionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OAuthScopesRestrictionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OAuthScopesRestrictionType(value)
	return nil
}

// NewOAuthScopesRestrictionTypeFromValue returns a pointer to a valid OAuthScopesRestrictionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOAuthScopesRestrictionTypeFromValue(v string) (*OAuthScopesRestrictionType, error) {
	ev := OAuthScopesRestrictionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OAuthScopesRestrictionType: valid values are %v", v, allowedOAuthScopesRestrictionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OAuthScopesRestrictionType) IsValid() bool {
	for _, existing := range allowedOAuthScopesRestrictionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuthScopesRestrictionType value.
func (v OAuthScopesRestrictionType) Ptr() *OAuthScopesRestrictionType {
	return &v
}
