// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuthClientRegistrationResponseType OAuth 2.0 response type that a registered client may use.
type OAuthClientRegistrationResponseType string

// List of OAuthClientRegistrationResponseType.
const (
	OAUTHCLIENTREGISTRATIONRESPONSETYPE_CODE OAuthClientRegistrationResponseType = "code"
)

var allowedOAuthClientRegistrationResponseTypeEnumValues = []OAuthClientRegistrationResponseType{
	OAUTHCLIENTREGISTRATIONRESPONSETYPE_CODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OAuthClientRegistrationResponseType) GetAllowedValues() []OAuthClientRegistrationResponseType {
	return allowedOAuthClientRegistrationResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OAuthClientRegistrationResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OAuthClientRegistrationResponseType(value)
	return nil
}

// NewOAuthClientRegistrationResponseTypeFromValue returns a pointer to a valid OAuthClientRegistrationResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOAuthClientRegistrationResponseTypeFromValue(v string) (*OAuthClientRegistrationResponseType, error) {
	ev := OAuthClientRegistrationResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OAuthClientRegistrationResponseType: valid values are %v", v, allowedOAuthClientRegistrationResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OAuthClientRegistrationResponseType) IsValid() bool {
	for _, existing := range allowedOAuthClientRegistrationResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuthClientRegistrationResponseType value.
func (v OAuthClientRegistrationResponseType) Ptr() *OAuthClientRegistrationResponseType {
	return &v
}
