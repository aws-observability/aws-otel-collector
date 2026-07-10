// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfigurationsType SAML configurations resource type.
type SAMLConfigurationsType string

// List of SAMLConfigurationsType.
const (
	SAMLCONFIGURATIONSTYPE_SAML_CONFIGURATIONS SAMLConfigurationsType = "saml_configurations"
)

var allowedSAMLConfigurationsTypeEnumValues = []SAMLConfigurationsType{
	SAMLCONFIGURATIONSTYPE_SAML_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SAMLConfigurationsType) GetAllowedValues() []SAMLConfigurationsType {
	return allowedSAMLConfigurationsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SAMLConfigurationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SAMLConfigurationsType(value)
	return nil
}

// NewSAMLConfigurationsTypeFromValue returns a pointer to a valid SAMLConfigurationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSAMLConfigurationsTypeFromValue(v string) (*SAMLConfigurationsType, error) {
	ev := SAMLConfigurationsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SAMLConfigurationsType: valid values are %v", v, allowedSAMLConfigurationsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SAMLConfigurationsType) IsValid() bool {
	for _, existing := range allowedSAMLConfigurationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SAMLConfigurationsType value.
func (v SAMLConfigurationsType) Ptr() *SAMLConfigurationsType {
	return &v
}
