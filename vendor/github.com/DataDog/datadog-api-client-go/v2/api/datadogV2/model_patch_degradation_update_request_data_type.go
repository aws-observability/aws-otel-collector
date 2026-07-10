// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationUpdateRequestDataType Degradation updates resource type.
type PatchDegradationUpdateRequestDataType string

// List of PatchDegradationUpdateRequestDataType.
const (
	PATCHDEGRADATIONUPDATEREQUESTDATATYPE_DEGRADATION_UPDATES PatchDegradationUpdateRequestDataType = "degradation_updates"
)

var allowedPatchDegradationUpdateRequestDataTypeEnumValues = []PatchDegradationUpdateRequestDataType{
	PATCHDEGRADATIONUPDATEREQUESTDATATYPE_DEGRADATION_UPDATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PatchDegradationUpdateRequestDataType) GetAllowedValues() []PatchDegradationUpdateRequestDataType {
	return allowedPatchDegradationUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PatchDegradationUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PatchDegradationUpdateRequestDataType(value)
	return nil
}

// NewPatchDegradationUpdateRequestDataTypeFromValue returns a pointer to a valid PatchDegradationUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPatchDegradationUpdateRequestDataTypeFromValue(v string) (*PatchDegradationUpdateRequestDataType, error) {
	ev := PatchDegradationUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PatchDegradationUpdateRequestDataType: valid values are %v", v, allowedPatchDegradationUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PatchDegradationUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedPatchDegradationUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchDegradationUpdateRequestDataType value.
func (v PatchDegradationUpdateRequestDataType) Ptr() *PatchDegradationUpdateRequestDataType {
	return &v
}
