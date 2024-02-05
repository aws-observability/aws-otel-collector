// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthOauthROPType The type of basic authentication to use when performing the test.
type SyntheticsBasicAuthOauthROPType string

// List of SyntheticsBasicAuthOauthROPType.
const (
	SYNTHETICSBASICAUTHOAUTHROPTYPE_OAUTH_ROP SyntheticsBasicAuthOauthROPType = "oauth-rop"
)

var allowedSyntheticsBasicAuthOauthROPTypeEnumValues = []SyntheticsBasicAuthOauthROPType{
	SYNTHETICSBASICAUTHOAUTHROPTYPE_OAUTH_ROP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthOauthROPType) GetAllowedValues() []SyntheticsBasicAuthOauthROPType {
	return allowedSyntheticsBasicAuthOauthROPTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthOauthROPType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthOauthROPType(value)
	return nil
}

// NewSyntheticsBasicAuthOauthROPTypeFromValue returns a pointer to a valid SyntheticsBasicAuthOauthROPType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthOauthROPTypeFromValue(v string) (*SyntheticsBasicAuthOauthROPType, error) {
	ev := SyntheticsBasicAuthOauthROPType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthOauthROPType: valid values are %v", v, allowedSyntheticsBasicAuthOauthROPTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthOauthROPType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthOauthROPTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthOauthROPType value.
func (v SyntheticsBasicAuthOauthROPType) Ptr() *SyntheticsBasicAuthOauthROPType {
	return &v
}
