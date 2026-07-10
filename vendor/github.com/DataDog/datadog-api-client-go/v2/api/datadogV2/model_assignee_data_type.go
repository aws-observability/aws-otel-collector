// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssigneeDataType Assignee resource type.
type AssigneeDataType string

// List of AssigneeDataType.
const (
	ASSIGNEEDATATYPE_ASSIGNEE AssigneeDataType = "assignee"
)

var allowedAssigneeDataTypeEnumValues = []AssigneeDataType{
	ASSIGNEEDATATYPE_ASSIGNEE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AssigneeDataType) GetAllowedValues() []AssigneeDataType {
	return allowedAssigneeDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AssigneeDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AssigneeDataType(value)
	return nil
}

// NewAssigneeDataTypeFromValue returns a pointer to a valid AssigneeDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAssigneeDataTypeFromValue(v string) (*AssigneeDataType, error) {
	ev := AssigneeDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AssigneeDataType: valid values are %v", v, allowedAssigneeDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AssigneeDataType) IsValid() bool {
	for _, existing := range allowedAssigneeDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssigneeDataType value.
func (v AssigneeDataType) Ptr() *AssigneeDataType {
	return &v
}
