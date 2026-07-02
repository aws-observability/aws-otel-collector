// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthClientRegistrationGrantType OAuth 2.0 grant type that a registered client may use.
type OAuthClientRegistrationGrantType string

// List of OAuthClientRegistrationGrantType.
const (
	OAUTHCLIENTREGISTRATIONGRANTTYPE_AUTHORIZATION_CODE OAuthClientRegistrationGrantType = "authorization_code"
	OAUTHCLIENTREGISTRATIONGRANTTYPE_REFRESH_TOKEN      OAuthClientRegistrationGrantType = "refresh_token"
)

var allowedOAuthClientRegistrationGrantTypeEnumValues = []OAuthClientRegistrationGrantType{
	OAUTHCLIENTREGISTRATIONGRANTTYPE_AUTHORIZATION_CODE,
	OAUTHCLIENTREGISTRATIONGRANTTYPE_REFRESH_TOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OAuthClientRegistrationGrantType) GetAllowedValues() []OAuthClientRegistrationGrantType {
	return allowedOAuthClientRegistrationGrantTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OAuthClientRegistrationGrantType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OAuthClientRegistrationGrantType(value)
	return nil
}

// NewOAuthClientRegistrationGrantTypeFromValue returns a pointer to a valid OAuthClientRegistrationGrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOAuthClientRegistrationGrantTypeFromValue(v string) (*OAuthClientRegistrationGrantType, error) {
	ev := OAuthClientRegistrationGrantType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OAuthClientRegistrationGrantType: valid values are %v", v, allowedOAuthClientRegistrationGrantTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OAuthClientRegistrationGrantType) IsValid() bool {
	for _, existing := range allowedOAuthClientRegistrationGrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuthClientRegistrationGrantType value.
func (v OAuthClientRegistrationGrantType) Ptr() *OAuthClientRegistrationGrantType {
	return &v
}
