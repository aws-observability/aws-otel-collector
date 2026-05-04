// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceDataAttributesStatus The status of the maintenance.
type MaintenanceDataAttributesStatus string

// List of MaintenanceDataAttributesStatus.
const (
	MAINTENANCEDATAATTRIBUTESSTATUS_SCHEDULED   MaintenanceDataAttributesStatus = "scheduled"
	MAINTENANCEDATAATTRIBUTESSTATUS_IN_PROGRESS MaintenanceDataAttributesStatus = "in_progress"
	MAINTENANCEDATAATTRIBUTESSTATUS_COMPLETED   MaintenanceDataAttributesStatus = "completed"
)

var allowedMaintenanceDataAttributesStatusEnumValues = []MaintenanceDataAttributesStatus{
	MAINTENANCEDATAATTRIBUTESSTATUS_SCHEDULED,
	MAINTENANCEDATAATTRIBUTESSTATUS_IN_PROGRESS,
	MAINTENANCEDATAATTRIBUTESSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MaintenanceDataAttributesStatus) GetAllowedValues() []MaintenanceDataAttributesStatus {
	return allowedMaintenanceDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MaintenanceDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MaintenanceDataAttributesStatus(value)
	return nil
}

// NewMaintenanceDataAttributesStatusFromValue returns a pointer to a valid MaintenanceDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMaintenanceDataAttributesStatusFromValue(v string) (*MaintenanceDataAttributesStatus, error) {
	ev := MaintenanceDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MaintenanceDataAttributesStatus: valid values are %v", v, allowedMaintenanceDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MaintenanceDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedMaintenanceDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaintenanceDataAttributesStatus value.
func (v MaintenanceDataAttributesStatus) Ptr() *MaintenanceDataAttributesStatus {
	return &v
}
