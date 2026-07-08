// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalOrgType The resource type for global user organizations.
type GlobalOrgType string

// List of GlobalOrgType.
const (
	GLOBALORGTYPE_GLOBAL_USER_ORGS GlobalOrgType = "global_user_orgs"
)

var allowedGlobalOrgTypeEnumValues = []GlobalOrgType{
	GLOBALORGTYPE_GLOBAL_USER_ORGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GlobalOrgType) GetAllowedValues() []GlobalOrgType {
	return allowedGlobalOrgTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GlobalOrgType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GlobalOrgType(value)
	return nil
}

// NewGlobalOrgTypeFromValue returns a pointer to a valid GlobalOrgType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGlobalOrgTypeFromValue(v string) (*GlobalOrgType, error) {
	ev := GlobalOrgType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GlobalOrgType: valid values are %v", v, allowedGlobalOrgTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GlobalOrgType) IsValid() bool {
	for _, existing := range allowedGlobalOrgTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GlobalOrgType value.
func (v GlobalOrgType) Ptr() *GlobalOrgType {
	return &v
}
