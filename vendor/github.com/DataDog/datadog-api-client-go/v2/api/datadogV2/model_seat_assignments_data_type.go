// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeatAssignmentsDataType Seat assignments resource type.
type SeatAssignmentsDataType string

// List of SeatAssignmentsDataType.
const (
	SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS SeatAssignmentsDataType = "seat-assignments"
)

var allowedSeatAssignmentsDataTypeEnumValues = []SeatAssignmentsDataType{
	SEATASSIGNMENTSDATATYPE_SEAT_ASSIGNMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeatAssignmentsDataType) GetAllowedValues() []SeatAssignmentsDataType {
	return allowedSeatAssignmentsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeatAssignmentsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeatAssignmentsDataType(value)
	return nil
}

// NewSeatAssignmentsDataTypeFromValue returns a pointer to a valid SeatAssignmentsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeatAssignmentsDataTypeFromValue(v string) (*SeatAssignmentsDataType, error) {
	ev := SeatAssignmentsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeatAssignmentsDataType: valid values are %v", v, allowedSeatAssignmentsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeatAssignmentsDataType) IsValid() bool {
	for _, existing := range allowedSeatAssignmentsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeatAssignmentsDataType value.
func (v SeatAssignmentsDataType) Ptr() *SeatAssignmentsDataType {
	return &v
}
