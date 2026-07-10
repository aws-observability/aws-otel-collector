// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccessTokensType Resource type returned by the access tokens list endpoint. Includes both personal and service access tokens.
type AccessTokensType string

// List of AccessTokensType.
const (
	ACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS AccessTokensType = "personal_access_tokens"
	ACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS  AccessTokensType = "service_access_tokens"
)

var allowedAccessTokensTypeEnumValues = []AccessTokensType{
	ACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
	ACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccessTokensType) GetAllowedValues() []AccessTokensType {
	return allowedAccessTokensTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccessTokensType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccessTokensType(value)
	return nil
}

// NewAccessTokensTypeFromValue returns a pointer to a valid AccessTokensType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccessTokensTypeFromValue(v string) (*AccessTokensType, error) {
	ev := AccessTokensType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccessTokensType: valid values are %v", v, allowedAccessTokensTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccessTokensType) IsValid() bool {
	for _, existing := range allowedAccessTokensTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessTokensType value.
func (v AccessTokensType) Ptr() *AccessTokensType {
	return &v
}
