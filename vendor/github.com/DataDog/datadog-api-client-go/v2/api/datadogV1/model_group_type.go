// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupType Set the sort type to group.
type GroupType string

// List of GroupType.
const (
	GROUPTYPE_GROUP GroupType = "group"
)

var allowedGroupTypeEnumValues = []GroupType{
	GROUPTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GroupType) GetAllowedValues() []GroupType {
	return allowedGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GroupType(value)
	return nil
}

// NewGroupTypeFromValue returns a pointer to a valid GroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGroupTypeFromValue(v string) (*GroupType, error) {
	ev := GroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GroupType: valid values are %v", v, allowedGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GroupType) IsValid() bool {
	for _, existing := range allowedGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupType value.
func (v GroupType) Ptr() *GroupType {
	return &v
}
