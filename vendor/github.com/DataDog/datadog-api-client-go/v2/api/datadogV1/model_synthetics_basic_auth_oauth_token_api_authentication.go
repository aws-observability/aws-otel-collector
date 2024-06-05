// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthOauthTokenApiAuthentication Type of token to use when performing the authentication.
type SyntheticsBasicAuthOauthTokenApiAuthentication string

// List of SyntheticsBasicAuthOauthTokenApiAuthentication.
const (
	SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_HEADER SyntheticsBasicAuthOauthTokenApiAuthentication = "header"
	SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_BODY   SyntheticsBasicAuthOauthTokenApiAuthentication = "body"
)

var allowedSyntheticsBasicAuthOauthTokenApiAuthenticationEnumValues = []SyntheticsBasicAuthOauthTokenApiAuthentication{
	SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_HEADER,
	SYNTHETICSBASICAUTHOAUTHTOKENAPIAUTHENTICATION_BODY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthOauthTokenApiAuthentication) GetAllowedValues() []SyntheticsBasicAuthOauthTokenApiAuthentication {
	return allowedSyntheticsBasicAuthOauthTokenApiAuthenticationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthOauthTokenApiAuthentication) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthOauthTokenApiAuthentication(value)
	return nil
}

// NewSyntheticsBasicAuthOauthTokenApiAuthenticationFromValue returns a pointer to a valid SyntheticsBasicAuthOauthTokenApiAuthentication
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthOauthTokenApiAuthenticationFromValue(v string) (*SyntheticsBasicAuthOauthTokenApiAuthentication, error) {
	ev := SyntheticsBasicAuthOauthTokenApiAuthentication(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthOauthTokenApiAuthentication: valid values are %v", v, allowedSyntheticsBasicAuthOauthTokenApiAuthenticationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthOauthTokenApiAuthentication) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthOauthTokenApiAuthenticationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthOauthTokenApiAuthentication value.
func (v SyntheticsBasicAuthOauthTokenApiAuthentication) Ptr() *SyntheticsBasicAuthOauthTokenApiAuthentication {
	return &v
}
