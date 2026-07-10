// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateMaintenanceRequestDataAttributesUpdatesItemsStatus The status of a maintenance update.
type CreateMaintenanceRequestDataAttributesUpdatesItemsStatus string

// List of CreateMaintenanceRequestDataAttributesUpdatesItemsStatus.
const (
	CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_IN_PROGRESS CreateMaintenanceRequestDataAttributesUpdatesItemsStatus = "in_progress"
	CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_COMPLETED   CreateMaintenanceRequestDataAttributesUpdatesItemsStatus = "completed"
)

var allowedCreateMaintenanceRequestDataAttributesUpdatesItemsStatusEnumValues = []CreateMaintenanceRequestDataAttributesUpdatesItemsStatus{
	CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_IN_PROGRESS,
	CREATEMAINTENANCEREQUESTDATAATTRIBUTESUPDATESITEMSSTATUS_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) GetAllowedValues() []CreateMaintenanceRequestDataAttributesUpdatesItemsStatus {
	return allowedCreateMaintenanceRequestDataAttributesUpdatesItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateMaintenanceRequestDataAttributesUpdatesItemsStatus(value)
	return nil
}

// NewCreateMaintenanceRequestDataAttributesUpdatesItemsStatusFromValue returns a pointer to a valid CreateMaintenanceRequestDataAttributesUpdatesItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateMaintenanceRequestDataAttributesUpdatesItemsStatusFromValue(v string) (*CreateMaintenanceRequestDataAttributesUpdatesItemsStatus, error) {
	ev := CreateMaintenanceRequestDataAttributesUpdatesItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateMaintenanceRequestDataAttributesUpdatesItemsStatus: valid values are %v", v, allowedCreateMaintenanceRequestDataAttributesUpdatesItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) IsValid() bool {
	for _, existing := range allowedCreateMaintenanceRequestDataAttributesUpdatesItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateMaintenanceRequestDataAttributesUpdatesItemsStatus value.
func (v CreateMaintenanceRequestDataAttributesUpdatesItemsStatus) Ptr() *CreateMaintenanceRequestDataAttributesUpdatesItemsStatus {
	return &v
}
