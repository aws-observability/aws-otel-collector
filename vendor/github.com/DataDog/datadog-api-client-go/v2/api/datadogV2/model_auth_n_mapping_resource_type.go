// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingResourceType The type of resource being mapped to.
type AuthNMappingResourceType string

// List of AuthNMappingResourceType.
const (
	AUTHNMAPPINGRESOURCETYPE_ROLE AuthNMappingResourceType = "role"
	AUTHNMAPPINGRESOURCETYPE_TEAM AuthNMappingResourceType = "team"
)

var allowedAuthNMappingResourceTypeEnumValues = []AuthNMappingResourceType{
	AUTHNMAPPINGRESOURCETYPE_ROLE,
	AUTHNMAPPINGRESOURCETYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AuthNMappingResourceType) GetAllowedValues() []AuthNMappingResourceType {
	return allowedAuthNMappingResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AuthNMappingResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AuthNMappingResourceType(value)
	return nil
}

// NewAuthNMappingResourceTypeFromValue returns a pointer to a valid AuthNMappingResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAuthNMappingResourceTypeFromValue(v string) (*AuthNMappingResourceType, error) {
	ev := AuthNMappingResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AuthNMappingResourceType: valid values are %v", v, allowedAuthNMappingResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AuthNMappingResourceType) IsValid() bool {
	for _, existing := range allowedAuthNMappingResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthNMappingResourceType value.
func (v AuthNMappingResourceType) Ptr() *AuthNMappingResourceType {
	return &v
}
