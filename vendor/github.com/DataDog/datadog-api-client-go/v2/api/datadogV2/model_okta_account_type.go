// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaAccountType Account type for an Okta account.
type OktaAccountType string

// List of OktaAccountType.
const (
	OKTAACCOUNTTYPE_OKTA_ACCOUNTS OktaAccountType = "okta-accounts"
)

var allowedOktaAccountTypeEnumValues = []OktaAccountType{
	OKTAACCOUNTTYPE_OKTA_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OktaAccountType) GetAllowedValues() []OktaAccountType {
	return allowedOktaAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OktaAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OktaAccountType(value)
	return nil
}

// NewOktaAccountTypeFromValue returns a pointer to a valid OktaAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOktaAccountTypeFromValue(v string) (*OktaAccountType, error) {
	ev := OktaAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OktaAccountType: valid values are %v", v, allowedOktaAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OktaAccountType) IsValid() bool {
	for _, existing := range allowedOktaAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OktaAccountType value.
func (v OktaAccountType) Ptr() *OktaAccountType {
	return &v
}
