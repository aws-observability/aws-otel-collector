// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgSAMLPreferencesType SAML preferences resource type.
type OrgSAMLPreferencesType string

// List of OrgSAMLPreferencesType.
const (
	ORGSAMLPREFERENCESTYPE_SAML_PREFERENCES OrgSAMLPreferencesType = "saml_preferences"
)

var allowedOrgSAMLPreferencesTypeEnumValues = []OrgSAMLPreferencesType{
	ORGSAMLPREFERENCESTYPE_SAML_PREFERENCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgSAMLPreferencesType) GetAllowedValues() []OrgSAMLPreferencesType {
	return allowedOrgSAMLPreferencesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgSAMLPreferencesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgSAMLPreferencesType(value)
	return nil
}

// NewOrgSAMLPreferencesTypeFromValue returns a pointer to a valid OrgSAMLPreferencesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgSAMLPreferencesTypeFromValue(v string) (*OrgSAMLPreferencesType, error) {
	ev := OrgSAMLPreferencesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgSAMLPreferencesType: valid values are %v", v, allowedOrgSAMLPreferencesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgSAMLPreferencesType) IsValid() bool {
	for _, existing := range allowedOrgSAMLPreferencesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgSAMLPreferencesType value.
func (v OrgSAMLPreferencesType) Ptr() *OrgSAMLPreferencesType {
	return &v
}
