// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardUsageType The type of the resource. Always `dashboards-usages`.
type DashboardUsageType string

// List of DashboardUsageType.
const (
	DASHBOARDUSAGETYPE_DASHBOARDS_USAGES DashboardUsageType = "dashboards-usages"
)

var allowedDashboardUsageTypeEnumValues = []DashboardUsageType{
	DASHBOARDUSAGETYPE_DASHBOARDS_USAGES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardUsageType) GetAllowedValues() []DashboardUsageType {
	return allowedDashboardUsageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardUsageType(value)
	return nil
}

// NewDashboardUsageTypeFromValue returns a pointer to a valid DashboardUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardUsageTypeFromValue(v string) (*DashboardUsageType, error) {
	ev := DashboardUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardUsageType: valid values are %v", v, allowedDashboardUsageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardUsageType) IsValid() bool {
	for _, existing := range allowedDashboardUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardUsageType value.
func (v DashboardUsageType) Ptr() *DashboardUsageType {
	return &v
}
