// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ShiftDataType Indicates that the resource is of type 'shifts'.
type ShiftDataType string

// List of ShiftDataType.
const (
	SHIFTDATATYPE_SHIFTS ShiftDataType = "shifts"
)

var allowedShiftDataTypeEnumValues = []ShiftDataType{
	SHIFTDATATYPE_SHIFTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ShiftDataType) GetAllowedValues() []ShiftDataType {
	return allowedShiftDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ShiftDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ShiftDataType(value)
	return nil
}

// NewShiftDataTypeFromValue returns a pointer to a valid ShiftDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewShiftDataTypeFromValue(v string) (*ShiftDataType, error) {
	ev := ShiftDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ShiftDataType: valid values are %v", v, allowedShiftDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ShiftDataType) IsValid() bool {
	for _, existing := range allowedShiftDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShiftDataType value.
func (v ShiftDataType) Ptr() *ShiftDataType {
	return &v
}
