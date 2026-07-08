// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardType Shared dashboard resource type.
type SharedDashboardType string

// List of SharedDashboardType.
const (
	SHAREDDASHBOARDTYPE_SHARED_DASHBOARD SharedDashboardType = "shared_dashboard"
)

var allowedSharedDashboardTypeEnumValues = []SharedDashboardType{
	SHAREDDASHBOARDTYPE_SHARED_DASHBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SharedDashboardType) GetAllowedValues() []SharedDashboardType {
	return allowedSharedDashboardTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SharedDashboardType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SharedDashboardType(value)
	return nil
}

// NewSharedDashboardTypeFromValue returns a pointer to a valid SharedDashboardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSharedDashboardTypeFromValue(v string) (*SharedDashboardType, error) {
	ev := SharedDashboardType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SharedDashboardType: valid values are %v", v, allowedSharedDashboardTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SharedDashboardType) IsValid() bool {
	for _, existing := range allowedSharedDashboardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDashboardType value.
func (v SharedDashboardType) Ptr() *SharedDashboardType {
	return &v
}
