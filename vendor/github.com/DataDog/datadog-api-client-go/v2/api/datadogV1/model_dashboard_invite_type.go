// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardInviteType Type for shared dashboard invitation request body.
type DashboardInviteType string

// List of DashboardInviteType.
const (
	DASHBOARDINVITETYPE_PUBLIC_DASHBOARD_INVITATION DashboardInviteType = "public_dashboard_invitation"
)

var allowedDashboardInviteTypeEnumValues = []DashboardInviteType{
	DASHBOARDINVITETYPE_PUBLIC_DASHBOARD_INVITATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardInviteType) GetAllowedValues() []DashboardInviteType {
	return allowedDashboardInviteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardInviteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardInviteType(value)
	return nil
}

// NewDashboardInviteTypeFromValue returns a pointer to a valid DashboardInviteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardInviteTypeFromValue(v string) (*DashboardInviteType, error) {
	ev := DashboardInviteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardInviteType: valid values are %v", v, allowedDashboardInviteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardInviteType) IsValid() bool {
	for _, existing := range allowedDashboardInviteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardInviteType value.
func (v DashboardInviteType) Ptr() *DashboardInviteType {
	return &v
}
