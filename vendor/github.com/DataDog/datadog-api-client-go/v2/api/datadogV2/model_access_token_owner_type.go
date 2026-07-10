// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccessTokenOwnerType Owner resource type. Either a user or a service account.
type AccessTokenOwnerType string

// List of AccessTokenOwnerType.
const (
	ACCESSTOKENOWNERTYPE_USERS           AccessTokenOwnerType = "users"
	ACCESSTOKENOWNERTYPE_SERVICE_ACCOUNT AccessTokenOwnerType = "service_account"
)

var allowedAccessTokenOwnerTypeEnumValues = []AccessTokenOwnerType{
	ACCESSTOKENOWNERTYPE_USERS,
	ACCESSTOKENOWNERTYPE_SERVICE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AccessTokenOwnerType) GetAllowedValues() []AccessTokenOwnerType {
	return allowedAccessTokenOwnerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AccessTokenOwnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AccessTokenOwnerType(value)
	return nil
}

// NewAccessTokenOwnerTypeFromValue returns a pointer to a valid AccessTokenOwnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAccessTokenOwnerTypeFromValue(v string) (*AccessTokenOwnerType, error) {
	ev := AccessTokenOwnerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AccessTokenOwnerType: valid values are %v", v, allowedAccessTokenOwnerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AccessTokenOwnerType) IsValid() bool {
	for _, existing := range allowedAccessTokenOwnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessTokenOwnerType value.
func (v AccessTokenOwnerType) Ptr() *AccessTokenOwnerType {
	return &v
}
