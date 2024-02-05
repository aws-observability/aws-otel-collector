// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingsSort Sorting options for AuthN Mappings.
type AuthNMappingsSort string

// List of AuthNMappingsSort.
const (
	AUTHNMAPPINGSSORT_CREATED_AT_ASCENDING                      AuthNMappingsSort = "created_at"
	AUTHNMAPPINGSSORT_CREATED_AT_DESCENDING                     AuthNMappingsSort = "-created_at"
	AUTHNMAPPINGSSORT_ROLE_ID_ASCENDING                         AuthNMappingsSort = "role_id"
	AUTHNMAPPINGSSORT_ROLE_ID_DESCENDING                        AuthNMappingsSort = "-role_id"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_ID_ASCENDING     AuthNMappingsSort = "saml_assertion_attribute_id"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_ID_DESCENDING    AuthNMappingsSort = "-saml_assertion_attribute_id"
	AUTHNMAPPINGSSORT_ROLE_NAME_ASCENDING                       AuthNMappingsSort = "role.name"
	AUTHNMAPPINGSSORT_ROLE_NAME_DESCENDING                      AuthNMappingsSort = "-role.name"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_KEY_ASCENDING    AuthNMappingsSort = "saml_assertion_attribute.attribute_key"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_KEY_DESCENDING   AuthNMappingsSort = "-saml_assertion_attribute.attribute_key"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_VALUE_ASCENDING  AuthNMappingsSort = "saml_assertion_attribute.attribute_value"
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_VALUE_DESCENDING AuthNMappingsSort = "-saml_assertion_attribute.attribute_value"
)

var allowedAuthNMappingsSortEnumValues = []AuthNMappingsSort{
	AUTHNMAPPINGSSORT_CREATED_AT_ASCENDING,
	AUTHNMAPPINGSSORT_CREATED_AT_DESCENDING,
	AUTHNMAPPINGSSORT_ROLE_ID_ASCENDING,
	AUTHNMAPPINGSSORT_ROLE_ID_DESCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_ID_ASCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_ID_DESCENDING,
	AUTHNMAPPINGSSORT_ROLE_NAME_ASCENDING,
	AUTHNMAPPINGSSORT_ROLE_NAME_DESCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_KEY_ASCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_KEY_DESCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_VALUE_ASCENDING,
	AUTHNMAPPINGSSORT_SAML_ASSERTION_ATTRIBUTE_VALUE_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AuthNMappingsSort) GetAllowedValues() []AuthNMappingsSort {
	return allowedAuthNMappingsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuthNMappingsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuthNMappingsSort(value)
	return nil
}

// NewAuthNMappingsSortFromValue returns a pointer to a valid AuthNMappingsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuthNMappingsSortFromValue(v string) (*AuthNMappingsSort, error) {
	ev := AuthNMappingsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuthNMappingsSort: valid values are %v", v, allowedAuthNMappingsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuthNMappingsSort) IsValid() bool {
	for _, existing := range allowedAuthNMappingsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthNMappingsSort value.
func (v AuthNMappingsSort) Ptr() *AuthNMappingsSort {
	return &v
}
