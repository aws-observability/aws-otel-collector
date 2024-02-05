// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardType The type of the dashboard.
type DashboardType string

// List of DashboardType.
const (
	DASHBOARDTYPE_CUSTOM_TIMEBOARD        DashboardType = "custom_timeboard"
	DASHBOARDTYPE_CUSTOM_SCREENBOARD      DashboardType = "custom_screenboard"
	DASHBOARDTYPE_INTEGRATION_SCREENBOARD DashboardType = "integration_screenboard"
	DASHBOARDTYPE_INTEGRATION_TIMEBOARD   DashboardType = "integration_timeboard"
	DASHBOARDTYPE_HOST_TIMEBOARD          DashboardType = "host_timeboard"
)

var allowedDashboardTypeEnumValues = []DashboardType{
	DASHBOARDTYPE_CUSTOM_TIMEBOARD,
	DASHBOARDTYPE_CUSTOM_SCREENBOARD,
	DASHBOARDTYPE_INTEGRATION_SCREENBOARD,
	DASHBOARDTYPE_INTEGRATION_TIMEBOARD,
	DASHBOARDTYPE_HOST_TIMEBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardType) GetAllowedValues() []DashboardType {
	return allowedDashboardTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardType(value)
	return nil
}

// NewDashboardTypeFromValue returns a pointer to a valid DashboardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardTypeFromValue(v string) (*DashboardType, error) {
	ev := DashboardType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardType: valid values are %v", v, allowedDashboardTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardType) IsValid() bool {
	for _, existing := range allowedDashboardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardType value.
func (v DashboardType) Ptr() *DashboardType {
	return &v
}
