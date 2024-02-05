// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardResourceType Dashboard resource type.
type DashboardResourceType string

// List of DashboardResourceType.
const (
	DASHBOARDRESOURCETYPE_DASHBOARD DashboardResourceType = "dashboard"
)

var allowedDashboardResourceTypeEnumValues = []DashboardResourceType{
	DASHBOARDRESOURCETYPE_DASHBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardResourceType) GetAllowedValues() []DashboardResourceType {
	return allowedDashboardResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardResourceType(value)
	return nil
}

// NewDashboardResourceTypeFromValue returns a pointer to a valid DashboardResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardResourceTypeFromValue(v string) (*DashboardResourceType, error) {
	ev := DashboardResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardResourceType: valid values are %v", v, allowedDashboardResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardResourceType) IsValid() bool {
	for _, existing := range allowedDashboardResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardResourceType value.
func (v DashboardResourceType) Ptr() *DashboardResourceType {
	return &v
}
