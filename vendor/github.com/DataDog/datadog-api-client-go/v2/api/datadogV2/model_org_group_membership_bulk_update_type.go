// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipBulkUpdateType Org group membership bulk update resource type.
type OrgGroupMembershipBulkUpdateType string

// List of OrgGroupMembershipBulkUpdateType.
const (
	ORGGROUPMEMBERSHIPBULKUPDATETYPE_ORG_GROUP_MEMBERSHIP_BULK_UPDATES OrgGroupMembershipBulkUpdateType = "org_group_membership_bulk_updates"
)

var allowedOrgGroupMembershipBulkUpdateTypeEnumValues = []OrgGroupMembershipBulkUpdateType{
	ORGGROUPMEMBERSHIPBULKUPDATETYPE_ORG_GROUP_MEMBERSHIP_BULK_UPDATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupMembershipBulkUpdateType) GetAllowedValues() []OrgGroupMembershipBulkUpdateType {
	return allowedOrgGroupMembershipBulkUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupMembershipBulkUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupMembershipBulkUpdateType(value)
	return nil
}

// NewOrgGroupMembershipBulkUpdateTypeFromValue returns a pointer to a valid OrgGroupMembershipBulkUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupMembershipBulkUpdateTypeFromValue(v string) (*OrgGroupMembershipBulkUpdateType, error) {
	ev := OrgGroupMembershipBulkUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupMembershipBulkUpdateType: valid values are %v", v, allowedOrgGroupMembershipBulkUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupMembershipBulkUpdateType) IsValid() bool {
	for _, existing := range allowedOrgGroupMembershipBulkUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupMembershipBulkUpdateType value.
func (v OrgGroupMembershipBulkUpdateType) Ptr() *OrgGroupMembershipBulkUpdateType {
	return &v
}
