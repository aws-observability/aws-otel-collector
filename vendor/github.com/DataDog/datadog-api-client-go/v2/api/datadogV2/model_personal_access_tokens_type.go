// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokensType Personal access tokens resource type.
type PersonalAccessTokensType string

// List of PersonalAccessTokensType.
const (
	PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS PersonalAccessTokensType = "personal_access_tokens"
)

var allowedPersonalAccessTokensTypeEnumValues = []PersonalAccessTokensType{
	PERSONALACCESSTOKENSTYPE_PERSONAL_ACCESS_TOKENS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PersonalAccessTokensType) GetAllowedValues() []PersonalAccessTokensType {
	return allowedPersonalAccessTokensTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PersonalAccessTokensType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PersonalAccessTokensType(value)
	return nil
}

// NewPersonalAccessTokensTypeFromValue returns a pointer to a valid PersonalAccessTokensType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPersonalAccessTokensTypeFromValue(v string) (*PersonalAccessTokensType, error) {
	ev := PersonalAccessTokensType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PersonalAccessTokensType: valid values are %v", v, allowedPersonalAccessTokensTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PersonalAccessTokensType) IsValid() bool {
	for _, existing := range allowedPersonalAccessTokensTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalAccessTokensType value.
func (v PersonalAccessTokensType) Ptr() *PersonalAccessTokensType {
	return &v
}
