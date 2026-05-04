// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMaintenanceRequestDataType Maintenances resource type.
type PatchMaintenanceRequestDataType string

// List of PatchMaintenanceRequestDataType.
const (
	PATCHMAINTENANCEREQUESTDATATYPE_MAINTENANCES PatchMaintenanceRequestDataType = "maintenances"
)

var allowedPatchMaintenanceRequestDataTypeEnumValues = []PatchMaintenanceRequestDataType{
	PATCHMAINTENANCEREQUESTDATATYPE_MAINTENANCES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchMaintenanceRequestDataType) GetAllowedValues() []PatchMaintenanceRequestDataType {
	return allowedPatchMaintenanceRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchMaintenanceRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchMaintenanceRequestDataType(value)
	return nil
}

// NewPatchMaintenanceRequestDataTypeFromValue returns a pointer to a valid PatchMaintenanceRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchMaintenanceRequestDataTypeFromValue(v string) (*PatchMaintenanceRequestDataType, error) {
	ev := PatchMaintenanceRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchMaintenanceRequestDataType: valid values are %v", v, allowedPatchMaintenanceRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchMaintenanceRequestDataType) IsValid() bool {
	for _, existing := range allowedPatchMaintenanceRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchMaintenanceRequestDataType value.
func (v PatchMaintenanceRequestDataType) Ptr() *PatchMaintenanceRequestDataType {
	return &v
}
