// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardIncludedDashboardType Included dashboard resource type.
type SharedDashboardIncludedDashboardType string

// List of SharedDashboardIncludedDashboardType.
const (
	SHAREDDASHBOARDINCLUDEDDASHBOARDTYPE_DASHBOARD SharedDashboardIncludedDashboardType = "dashboard"
)

var allowedSharedDashboardIncludedDashboardTypeEnumValues = []SharedDashboardIncludedDashboardType{
	SHAREDDASHBOARDINCLUDEDDASHBOARDTYPE_DASHBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SharedDashboardIncludedDashboardType) GetAllowedValues() []SharedDashboardIncludedDashboardType {
	return allowedSharedDashboardIncludedDashboardTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SharedDashboardIncludedDashboardType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SharedDashboardIncludedDashboardType(value)
	return nil
}

// NewSharedDashboardIncludedDashboardTypeFromValue returns a pointer to a valid SharedDashboardIncludedDashboardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSharedDashboardIncludedDashboardTypeFromValue(v string) (*SharedDashboardIncludedDashboardType, error) {
	ev := SharedDashboardIncludedDashboardType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SharedDashboardIncludedDashboardType: valid values are %v", v, allowedSharedDashboardIncludedDashboardTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SharedDashboardIncludedDashboardType) IsValid() bool {
	for _, existing := range allowedSharedDashboardIncludedDashboardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDashboardIncludedDashboardType value.
func (v SharedDashboardIncludedDashboardType) Ptr() *SharedDashboardIncludedDashboardType {
	return &v
}
