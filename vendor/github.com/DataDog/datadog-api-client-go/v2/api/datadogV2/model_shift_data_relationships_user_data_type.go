// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ShiftDataRelationshipsUserDataType Indicates that the related resource is of type 'users'.
type ShiftDataRelationshipsUserDataType string

// List of ShiftDataRelationshipsUserDataType.
const (
	SHIFTDATARELATIONSHIPSUSERDATATYPE_USERS ShiftDataRelationshipsUserDataType = "users"
)

var allowedShiftDataRelationshipsUserDataTypeEnumValues = []ShiftDataRelationshipsUserDataType{
	SHIFTDATARELATIONSHIPSUSERDATATYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ShiftDataRelationshipsUserDataType) GetAllowedValues() []ShiftDataRelationshipsUserDataType {
	return allowedShiftDataRelationshipsUserDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ShiftDataRelationshipsUserDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ShiftDataRelationshipsUserDataType(value)
	return nil
}

// NewShiftDataRelationshipsUserDataTypeFromValue returns a pointer to a valid ShiftDataRelationshipsUserDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewShiftDataRelationshipsUserDataTypeFromValue(v string) (*ShiftDataRelationshipsUserDataType, error) {
	ev := ShiftDataRelationshipsUserDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ShiftDataRelationshipsUserDataType: valid values are %v", v, allowedShiftDataRelationshipsUserDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ShiftDataRelationshipsUserDataType) IsValid() bool {
	for _, existing := range allowedShiftDataRelationshipsUserDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShiftDataRelationshipsUserDataType value.
func (v ShiftDataRelationshipsUserDataType) Ptr() *ShiftDataRelationshipsUserDataType {
	return &v
}
