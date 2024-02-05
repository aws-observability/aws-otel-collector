// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardLayoutType Layout type of the dashboard.
type DashboardLayoutType string

// List of DashboardLayoutType.
const (
	DASHBOARDLAYOUTTYPE_ORDERED DashboardLayoutType = "ordered"
	DASHBOARDLAYOUTTYPE_FREE    DashboardLayoutType = "free"
)

var allowedDashboardLayoutTypeEnumValues = []DashboardLayoutType{
	DASHBOARDLAYOUTTYPE_ORDERED,
	DASHBOARDLAYOUTTYPE_FREE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardLayoutType) GetAllowedValues() []DashboardLayoutType {
	return allowedDashboardLayoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardLayoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardLayoutType(value)
	return nil
}

// NewDashboardLayoutTypeFromValue returns a pointer to a valid DashboardLayoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardLayoutTypeFromValue(v string) (*DashboardLayoutType, error) {
	ev := DashboardLayoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardLayoutType: valid values are %v", v, allowedDashboardLayoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardLayoutType) IsValid() bool {
	for _, existing := range allowedDashboardLayoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardLayoutType value.
func (v DashboardLayoutType) Ptr() *DashboardLayoutType {
	return &v
}
