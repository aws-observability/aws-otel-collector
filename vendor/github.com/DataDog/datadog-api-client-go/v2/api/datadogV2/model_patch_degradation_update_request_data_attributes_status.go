// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationUpdateRequestDataAttributesStatus The status of the degradation update.
type PatchDegradationUpdateRequestDataAttributesStatus string

// List of PatchDegradationUpdateRequestDataAttributesStatus.
const (
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING PatchDegradationUpdateRequestDataAttributesStatus = "investigating"
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED    PatchDegradationUpdateRequestDataAttributesStatus = "identified"
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_MONITORING    PatchDegradationUpdateRequestDataAttributesStatus = "monitoring"
)

var allowedPatchDegradationUpdateRequestDataAttributesStatusEnumValues = []PatchDegradationUpdateRequestDataAttributesStatus{
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED,
	PATCHDEGRADATIONUPDATEREQUESTDATAATTRIBUTESSTATUS_MONITORING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchDegradationUpdateRequestDataAttributesStatus) GetAllowedValues() []PatchDegradationUpdateRequestDataAttributesStatus {
	return allowedPatchDegradationUpdateRequestDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchDegradationUpdateRequestDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchDegradationUpdateRequestDataAttributesStatus(value)
	return nil
}

// NewPatchDegradationUpdateRequestDataAttributesStatusFromValue returns a pointer to a valid PatchDegradationUpdateRequestDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchDegradationUpdateRequestDataAttributesStatusFromValue(v string) (*PatchDegradationUpdateRequestDataAttributesStatus, error) {
	ev := PatchDegradationUpdateRequestDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchDegradationUpdateRequestDataAttributesStatus: valid values are %v", v, allowedPatchDegradationUpdateRequestDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchDegradationUpdateRequestDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedPatchDegradationUpdateRequestDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchDegradationUpdateRequestDataAttributesStatus value.
func (v PatchDegradationUpdateRequestDataAttributesStatus) Ptr() *PatchDegradationUpdateRequestDataAttributesStatus {
	return &v
}
