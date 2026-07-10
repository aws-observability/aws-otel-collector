// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MaintenanceWindowResourceType JSON:API resource type for maintenance windows.
type MaintenanceWindowResourceType string

// List of MaintenanceWindowResourceType.
const (
	MAINTENANCEWINDOWRESOURCETYPE_MAINTENANCE_WINDOW MaintenanceWindowResourceType = "maintenance_window"
)

var allowedMaintenanceWindowResourceTypeEnumValues = []MaintenanceWindowResourceType{
	MAINTENANCEWINDOWRESOURCETYPE_MAINTENANCE_WINDOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MaintenanceWindowResourceType) GetAllowedValues() []MaintenanceWindowResourceType {
	return allowedMaintenanceWindowResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MaintenanceWindowResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MaintenanceWindowResourceType(value)
	return nil
}

// NewMaintenanceWindowResourceTypeFromValue returns a pointer to a valid MaintenanceWindowResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMaintenanceWindowResourceTypeFromValue(v string) (*MaintenanceWindowResourceType, error) {
	ev := MaintenanceWindowResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MaintenanceWindowResourceType: valid values are %v", v, allowedMaintenanceWindowResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MaintenanceWindowResourceType) IsValid() bool {
	for _, existing := range allowedMaintenanceWindowResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaintenanceWindowResourceType value.
func (v MaintenanceWindowResourceType) Ptr() *MaintenanceWindowResourceType {
	return &v
}
