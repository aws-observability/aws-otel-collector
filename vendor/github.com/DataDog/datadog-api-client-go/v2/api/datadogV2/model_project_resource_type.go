// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectResourceType Project resource type
type ProjectResourceType string

// List of ProjectResourceType.
const (
	PROJECTRESOURCETYPE_PROJECT ProjectResourceType = "project"
)

var allowedProjectResourceTypeEnumValues = []ProjectResourceType{
	PROJECTRESOURCETYPE_PROJECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProjectResourceType) GetAllowedValues() []ProjectResourceType {
	return allowedProjectResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProjectResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProjectResourceType(value)
	return nil
}

// NewProjectResourceTypeFromValue returns a pointer to a valid ProjectResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProjectResourceTypeFromValue(v string) (*ProjectResourceType, error) {
	ev := ProjectResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProjectResourceType: valid values are %v", v, allowedProjectResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProjectResourceType) IsValid() bool {
	for _, existing := range allowedProjectResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectResourceType value.
func (v ProjectResourceType) Ptr() *ProjectResourceType {
	return &v
}
