// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowAssignmentGroupType Type identifier for ServiceNow assignment group resources
type ServiceNowAssignmentGroupType string

// List of ServiceNowAssignmentGroupType.
const (
	SERVICENOWASSIGNMENTGROUPTYPE_ASSIGNMENT_GROUPS ServiceNowAssignmentGroupType = "assignment_groups"
)

var allowedServiceNowAssignmentGroupTypeEnumValues = []ServiceNowAssignmentGroupType{
	SERVICENOWASSIGNMENTGROUPTYPE_ASSIGNMENT_GROUPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceNowAssignmentGroupType) GetAllowedValues() []ServiceNowAssignmentGroupType {
	return allowedServiceNowAssignmentGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceNowAssignmentGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceNowAssignmentGroupType(value)
	return nil
}

// NewServiceNowAssignmentGroupTypeFromValue returns a pointer to a valid ServiceNowAssignmentGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceNowAssignmentGroupTypeFromValue(v string) (*ServiceNowAssignmentGroupType, error) {
	ev := ServiceNowAssignmentGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceNowAssignmentGroupType: valid values are %v", v, allowedServiceNowAssignmentGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceNowAssignmentGroupType) IsValid() bool {
	for _, existing := range allowedServiceNowAssignmentGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceNowAssignmentGroupType value.
func (v ServiceNowAssignmentGroupType) Ptr() *ServiceNowAssignmentGroupType {
	return &v
}
