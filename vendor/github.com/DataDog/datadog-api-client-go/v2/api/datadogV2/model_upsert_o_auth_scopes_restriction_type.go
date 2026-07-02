// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertOAuthScopesRestrictionType JSON:API resource type for an upsert OAuth2 client scopes restriction request.
type UpsertOAuthScopesRestrictionType string

// List of UpsertOAuthScopesRestrictionType.
const (
	UPSERTOAUTHSCOPESRESTRICTIONTYPE_UPSERT_SCOPES_RESTRICTION UpsertOAuthScopesRestrictionType = "upsert_scopes_restriction"
)

var allowedUpsertOAuthScopesRestrictionTypeEnumValues = []UpsertOAuthScopesRestrictionType{
	UPSERTOAUTHSCOPESRESTRICTIONTYPE_UPSERT_SCOPES_RESTRICTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpsertOAuthScopesRestrictionType) GetAllowedValues() []UpsertOAuthScopesRestrictionType {
	return allowedUpsertOAuthScopesRestrictionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpsertOAuthScopesRestrictionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpsertOAuthScopesRestrictionType(value)
	return nil
}

// NewUpsertOAuthScopesRestrictionTypeFromValue returns a pointer to a valid UpsertOAuthScopesRestrictionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpsertOAuthScopesRestrictionTypeFromValue(v string) (*UpsertOAuthScopesRestrictionType, error) {
	ev := UpsertOAuthScopesRestrictionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpsertOAuthScopesRestrictionType: valid values are %v", v, allowedUpsertOAuthScopesRestrictionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpsertOAuthScopesRestrictionType) IsValid() bool {
	for _, existing := range allowedUpsertOAuthScopesRestrictionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpsertOAuthScopesRestrictionType value.
func (v UpsertOAuthScopesRestrictionType) Ptr() *UpsertOAuthScopesRestrictionType {
	return &v
}
