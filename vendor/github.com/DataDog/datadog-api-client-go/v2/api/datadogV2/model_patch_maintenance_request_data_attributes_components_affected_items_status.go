// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus The status of the component.
type PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus string

// List of PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus.
const (
	PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus = "operational"
	PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_MAINTENANCE PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus = "maintenance"
)

var allowedPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusEnumValues = []PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus{
	PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_OPERATIONAL,
	PATCHMAINTENANCEREQUESTDATAATTRIBUTESCOMPONENTSAFFECTEDITEMSSTATUS_MAINTENANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus) GetAllowedValues() []PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus {
	return allowedPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus(value)
	return nil
}

// NewPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusFromValue returns a pointer to a valid PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusFromValue(v string) (*PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus, error) {
	ev := PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus: valid values are %v", v, allowedPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus) IsValid() bool {
	for _, existing := range allowedPatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus value.
func (v PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus) Ptr() *PatchMaintenanceRequestDataAttributesComponentsAffectedItemsStatus {
	return &v
}
