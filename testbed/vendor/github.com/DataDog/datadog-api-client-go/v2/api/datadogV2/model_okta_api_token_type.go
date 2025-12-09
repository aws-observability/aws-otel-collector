// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaAPITokenType The definition of the `OktaAPIToken` object.
type OktaAPITokenType string

// List of OktaAPITokenType.
const (
	OKTAAPITOKENTYPE_OKTAAPITOKEN OktaAPITokenType = "OktaAPIToken"
)

var allowedOktaAPITokenTypeEnumValues = []OktaAPITokenType{
	OKTAAPITOKENTYPE_OKTAAPITOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OktaAPITokenType) GetAllowedValues() []OktaAPITokenType {
	return allowedOktaAPITokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OktaAPITokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OktaAPITokenType(value)
	return nil
}

// NewOktaAPITokenTypeFromValue returns a pointer to a valid OktaAPITokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOktaAPITokenTypeFromValue(v string) (*OktaAPITokenType, error) {
	ev := OktaAPITokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OktaAPITokenType: valid values are %v", v, allowedOktaAPITokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OktaAPITokenType) IsValid() bool {
	for _, existing := range allowedOktaAPITokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OktaAPITokenType value.
func (v OktaAPITokenType) Ptr() *OktaAPITokenType {
	return &v
}
