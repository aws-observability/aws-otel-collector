// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyAccountType The JSON:API type for this API. Should always be `fastly-accounts`.
type FastlyAccountType string

// List of FastlyAccountType.
const (
	FASTLYACCOUNTTYPE_FASTLY_ACCOUNTS FastlyAccountType = "fastly-accounts"
)

var allowedFastlyAccountTypeEnumValues = []FastlyAccountType{
	FASTLYACCOUNTTYPE_FASTLY_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FastlyAccountType) GetAllowedValues() []FastlyAccountType {
	return allowedFastlyAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FastlyAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FastlyAccountType(value)
	return nil
}

// NewFastlyAccountTypeFromValue returns a pointer to a valid FastlyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFastlyAccountTypeFromValue(v string) (*FastlyAccountType, error) {
	ev := FastlyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FastlyAccountType: valid values are %v", v, allowedFastlyAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FastlyAccountType) IsValid() bool {
	for _, existing := range allowedFastlyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FastlyAccountType value.
func (v FastlyAccountType) Ptr() *FastlyAccountType {
	return &v
}
