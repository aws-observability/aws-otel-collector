// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetTeamMembershipsSort Specifies the order of returned team memberships
type GetTeamMembershipsSort string

// List of GetTeamMembershipsSort.
const (
	GETTEAMMEMBERSHIPSSORT_MANAGER_NAME  GetTeamMembershipsSort = "manager_name"
	GETTEAMMEMBERSHIPSSORT__MANAGER_NAME GetTeamMembershipsSort = "-manager_name"
	GETTEAMMEMBERSHIPSSORT_NAME          GetTeamMembershipsSort = "name"
	GETTEAMMEMBERSHIPSSORT__NAME         GetTeamMembershipsSort = "-name"
	GETTEAMMEMBERSHIPSSORT_HANDLE        GetTeamMembershipsSort = "handle"
	GETTEAMMEMBERSHIPSSORT__HANDLE       GetTeamMembershipsSort = "-handle"
	GETTEAMMEMBERSHIPSSORT_EMAIL         GetTeamMembershipsSort = "email"
	GETTEAMMEMBERSHIPSSORT__EMAIL        GetTeamMembershipsSort = "-email"
)

var allowedGetTeamMembershipsSortEnumValues = []GetTeamMembershipsSort{
	GETTEAMMEMBERSHIPSSORT_MANAGER_NAME,
	GETTEAMMEMBERSHIPSSORT__MANAGER_NAME,
	GETTEAMMEMBERSHIPSSORT_NAME,
	GETTEAMMEMBERSHIPSSORT__NAME,
	GETTEAMMEMBERSHIPSSORT_HANDLE,
	GETTEAMMEMBERSHIPSSORT__HANDLE,
	GETTEAMMEMBERSHIPSSORT_EMAIL,
	GETTEAMMEMBERSHIPSSORT__EMAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetTeamMembershipsSort) GetAllowedValues() []GetTeamMembershipsSort {
	return allowedGetTeamMembershipsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetTeamMembershipsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetTeamMembershipsSort(value)
	return nil
}

// NewGetTeamMembershipsSortFromValue returns a pointer to a valid GetTeamMembershipsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetTeamMembershipsSortFromValue(v string) (*GetTeamMembershipsSort, error) {
	ev := GetTeamMembershipsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetTeamMembershipsSort: valid values are %v", v, allowedGetTeamMembershipsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetTeamMembershipsSort) IsValid() bool {
	for _, existing := range allowedGetTeamMembershipsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetTeamMembershipsSort value.
func (v GetTeamMembershipsSort) Ptr() *GetTeamMembershipsSort {
	return &v
}
