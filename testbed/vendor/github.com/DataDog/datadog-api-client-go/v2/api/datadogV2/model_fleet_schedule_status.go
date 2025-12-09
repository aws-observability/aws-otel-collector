// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleStatus The status of the schedule.
// - `active`: The schedule is active and will create deployments according to its recurrence rule.
// - `inactive`: The schedule is inactive and will not create any deployments.
type FleetScheduleStatus string

// List of FleetScheduleStatus.
const (
	FLEETSCHEDULESTATUS_ACTIVE   FleetScheduleStatus = "active"
	FLEETSCHEDULESTATUS_INACTIVE FleetScheduleStatus = "inactive"
)

var allowedFleetScheduleStatusEnumValues = []FleetScheduleStatus{
	FLEETSCHEDULESTATUS_ACTIVE,
	FLEETSCHEDULESTATUS_INACTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetScheduleStatus) GetAllowedValues() []FleetScheduleStatus {
	return allowedFleetScheduleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetScheduleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetScheduleStatus(value)
	return nil
}

// NewFleetScheduleStatusFromValue returns a pointer to a valid FleetScheduleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetScheduleStatusFromValue(v string) (*FleetScheduleStatus, error) {
	ev := FleetScheduleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetScheduleStatus: valid values are %v", v, allowedFleetScheduleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetScheduleStatus) IsValid() bool {
	for _, existing := range allowedFleetScheduleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetScheduleStatus value.
func (v FleetScheduleStatus) Ptr() *FleetScheduleStatus {
	return &v
}
