// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardStatus Active means the dashboard is publicly available. Paused means the dashboard is not publicly available.
type SharedDashboardStatus string

// List of SharedDashboardStatus.
const (
	SHAREDDASHBOARDSTATUS_ACTIVE SharedDashboardStatus = "active"
	SHAREDDASHBOARDSTATUS_PAUSED SharedDashboardStatus = "paused"
)

var allowedSharedDashboardStatusEnumValues = []SharedDashboardStatus{
	SHAREDDASHBOARDSTATUS_ACTIVE,
	SHAREDDASHBOARDSTATUS_PAUSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SharedDashboardStatus) GetAllowedValues() []SharedDashboardStatus {
	return allowedSharedDashboardStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SharedDashboardStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SharedDashboardStatus(value)
	return nil
}

// NewSharedDashboardStatusFromValue returns a pointer to a valid SharedDashboardStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSharedDashboardStatusFromValue(v string) (*SharedDashboardStatus, error) {
	ev := SharedDashboardStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SharedDashboardStatus: valid values are %v", v, allowedSharedDashboardStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SharedDashboardStatus) IsValid() bool {
	for _, existing := range allowedSharedDashboardStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDashboardStatus value.
func (v SharedDashboardStatus) Ptr() *SharedDashboardStatus {
	return &v
}
