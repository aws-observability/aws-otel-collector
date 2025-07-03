// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleMemberType Schedule Members resource type.
type ScheduleMemberType string

// List of ScheduleMemberType.
const (
	SCHEDULEMEMBERTYPE_MEMBERS ScheduleMemberType = "members"
)

var allowedScheduleMemberTypeEnumValues = []ScheduleMemberType{
	SCHEDULEMEMBERTYPE_MEMBERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleMemberType) GetAllowedValues() []ScheduleMemberType {
	return allowedScheduleMemberTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleMemberType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleMemberType(value)
	return nil
}

// NewScheduleMemberTypeFromValue returns a pointer to a valid ScheduleMemberType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleMemberTypeFromValue(v string) (*ScheduleMemberType, error) {
	ev := ScheduleMemberType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleMemberType: valid values are %v", v, allowedScheduleMemberTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleMemberType) IsValid() bool {
	for _, existing := range allowedScheduleMemberTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleMemberType value.
func (v ScheduleMemberType) Ptr() *ScheduleMemberType {
	return &v
}
