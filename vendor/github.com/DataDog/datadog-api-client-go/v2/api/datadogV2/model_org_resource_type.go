// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgResourceType The resource type for organizations.
type OrgResourceType string

// List of OrgResourceType.
const (
	ORGRESOURCETYPE_ORGS OrgResourceType = "orgs"
)

var allowedOrgResourceTypeEnumValues = []OrgResourceType{
	ORGRESOURCETYPE_ORGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgResourceType) GetAllowedValues() []OrgResourceType {
	return allowedOrgResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgResourceType(value)
	return nil
}

// NewOrgResourceTypeFromValue returns a pointer to a valid OrgResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgResourceTypeFromValue(v string) (*OrgResourceType, error) {
	ev := OrgResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgResourceType: valid values are %v", v, allowedOrgResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgResourceType) IsValid() bool {
	for _, existing := range allowedOrgResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgResourceType value.
func (v OrgResourceType) Ptr() *OrgResourceType {
	return &v
}
