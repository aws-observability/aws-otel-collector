// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipSortOption Field to sort memberships by.
type OrgGroupMembershipSortOption string

// List of OrgGroupMembershipSortOption.
const (
	ORGGROUPMEMBERSHIPSORTOPTION_NAME       OrgGroupMembershipSortOption = "name"
	ORGGROUPMEMBERSHIPSORTOPTION_MINUS_NAME OrgGroupMembershipSortOption = "-name"
	ORGGROUPMEMBERSHIPSORTOPTION_UUID       OrgGroupMembershipSortOption = "uuid"
	ORGGROUPMEMBERSHIPSORTOPTION_MINUS_UUID OrgGroupMembershipSortOption = "-uuid"
)

var allowedOrgGroupMembershipSortOptionEnumValues = []OrgGroupMembershipSortOption{
	ORGGROUPMEMBERSHIPSORTOPTION_NAME,
	ORGGROUPMEMBERSHIPSORTOPTION_MINUS_NAME,
	ORGGROUPMEMBERSHIPSORTOPTION_UUID,
	ORGGROUPMEMBERSHIPSORTOPTION_MINUS_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupMembershipSortOption) GetAllowedValues() []OrgGroupMembershipSortOption {
	return allowedOrgGroupMembershipSortOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupMembershipSortOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupMembershipSortOption(value)
	return nil
}

// NewOrgGroupMembershipSortOptionFromValue returns a pointer to a valid OrgGroupMembershipSortOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupMembershipSortOptionFromValue(v string) (*OrgGroupMembershipSortOption, error) {
	ev := OrgGroupMembershipSortOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupMembershipSortOption: valid values are %v", v, allowedOrgGroupMembershipSortOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupMembershipSortOption) IsValid() bool {
	for _, existing := range allowedOrgGroupMembershipSortOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupMembershipSortOption value.
func (v OrgGroupMembershipSortOption) Ptr() *OrgGroupMembershipSortOption {
	return &v
}
