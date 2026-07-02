// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuth2WellKnownSitesEnvType JSON:API resource type for OAuth2 well-known sites environment.
type OAuth2WellKnownSitesEnvType string

// List of OAuth2WellKnownSitesEnvType.
const (
	OAUTH2WELLKNOWNSITESENVTYPE_ENV OAuth2WellKnownSitesEnvType = "env"
)

var allowedOAuth2WellKnownSitesEnvTypeEnumValues = []OAuth2WellKnownSitesEnvType{
	OAUTH2WELLKNOWNSITESENVTYPE_ENV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OAuth2WellKnownSitesEnvType) GetAllowedValues() []OAuth2WellKnownSitesEnvType {
	return allowedOAuth2WellKnownSitesEnvTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OAuth2WellKnownSitesEnvType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OAuth2WellKnownSitesEnvType(value)
	return nil
}

// NewOAuth2WellKnownSitesEnvTypeFromValue returns a pointer to a valid OAuth2WellKnownSitesEnvType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOAuth2WellKnownSitesEnvTypeFromValue(v string) (*OAuth2WellKnownSitesEnvType, error) {
	ev := OAuth2WellKnownSitesEnvType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OAuth2WellKnownSitesEnvType: valid values are %v", v, allowedOAuth2WellKnownSitesEnvTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OAuth2WellKnownSitesEnvType) IsValid() bool {
	for _, existing := range allowedOAuth2WellKnownSitesEnvTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OAuth2WellKnownSitesEnvType value.
func (v OAuth2WellKnownSitesEnvType) Ptr() *OAuth2WellKnownSitesEnvType {
	return &v
}
