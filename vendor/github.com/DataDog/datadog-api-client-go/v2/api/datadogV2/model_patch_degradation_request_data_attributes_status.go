// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationRequestDataAttributesStatus The status of the degradation.
type PatchDegradationRequestDataAttributesStatus string

// List of PatchDegradationRequestDataAttributesStatus.
const (
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING PatchDegradationRequestDataAttributesStatus = "investigating"
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED    PatchDegradationRequestDataAttributesStatus = "identified"
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_MONITORING    PatchDegradationRequestDataAttributesStatus = "monitoring"
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_RESOLVED      PatchDegradationRequestDataAttributesStatus = "resolved"
)

var allowedPatchDegradationRequestDataAttributesStatusEnumValues = []PatchDegradationRequestDataAttributesStatus{
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_INVESTIGATING,
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_IDENTIFIED,
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_MONITORING,
	PATCHDEGRADATIONREQUESTDATAATTRIBUTESSTATUS_RESOLVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchDegradationRequestDataAttributesStatus) GetAllowedValues() []PatchDegradationRequestDataAttributesStatus {
	return allowedPatchDegradationRequestDataAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchDegradationRequestDataAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchDegradationRequestDataAttributesStatus(value)
	return nil
}

// NewPatchDegradationRequestDataAttributesStatusFromValue returns a pointer to a valid PatchDegradationRequestDataAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchDegradationRequestDataAttributesStatusFromValue(v string) (*PatchDegradationRequestDataAttributesStatus, error) {
	ev := PatchDegradationRequestDataAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchDegradationRequestDataAttributesStatus: valid values are %v", v, allowedPatchDegradationRequestDataAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchDegradationRequestDataAttributesStatus) IsValid() bool {
	for _, existing := range allowedPatchDegradationRequestDataAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchDegradationRequestDataAttributesStatus value.
func (v PatchDegradationRequestDataAttributesStatus) Ptr() *PatchDegradationRequestDataAttributesStatus {
	return &v
}
