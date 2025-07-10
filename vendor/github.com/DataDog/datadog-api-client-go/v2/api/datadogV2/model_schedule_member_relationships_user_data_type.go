// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleMemberRelationshipsUserDataType Users resource type.
type ScheduleMemberRelationshipsUserDataType string

// List of ScheduleMemberRelationshipsUserDataType.
const (
	SCHEDULEMEMBERRELATIONSHIPSUSERDATATYPE_USERS ScheduleMemberRelationshipsUserDataType = "users"
)

var allowedScheduleMemberRelationshipsUserDataTypeEnumValues = []ScheduleMemberRelationshipsUserDataType{
	SCHEDULEMEMBERRELATIONSHIPSUSERDATATYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleMemberRelationshipsUserDataType) GetAllowedValues() []ScheduleMemberRelationshipsUserDataType {
	return allowedScheduleMemberRelationshipsUserDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleMemberRelationshipsUserDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleMemberRelationshipsUserDataType(value)
	return nil
}

// NewScheduleMemberRelationshipsUserDataTypeFromValue returns a pointer to a valid ScheduleMemberRelationshipsUserDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleMemberRelationshipsUserDataTypeFromValue(v string) (*ScheduleMemberRelationshipsUserDataType, error) {
	ev := ScheduleMemberRelationshipsUserDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleMemberRelationshipsUserDataType: valid values are %v", v, allowedScheduleMemberRelationshipsUserDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleMemberRelationshipsUserDataType) IsValid() bool {
	for _, existing := range allowedScheduleMemberRelationshipsUserDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleMemberRelationshipsUserDataType value.
func (v ScheduleMemberRelationshipsUserDataType) Ptr() *ScheduleMemberRelationshipsUserDataType {
	return &v
}
