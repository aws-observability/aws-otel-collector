// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseManagementProjectDataType Projects resource type.
type CaseManagementProjectDataType string

// List of CaseManagementProjectDataType.
const (
	CASEMANAGEMENTPROJECTDATATYPE_PROJECTS CaseManagementProjectDataType = "projects"
)

var allowedCaseManagementProjectDataTypeEnumValues = []CaseManagementProjectDataType{
	CASEMANAGEMENTPROJECTDATATYPE_PROJECTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseManagementProjectDataType) GetAllowedValues() []CaseManagementProjectDataType {
	return allowedCaseManagementProjectDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseManagementProjectDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseManagementProjectDataType(value)
	return nil
}

// NewCaseManagementProjectDataTypeFromValue returns a pointer to a valid CaseManagementProjectDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseManagementProjectDataTypeFromValue(v string) (*CaseManagementProjectDataType, error) {
	ev := CaseManagementProjectDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseManagementProjectDataType: valid values are %v", v, allowedCaseManagementProjectDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseManagementProjectDataType) IsValid() bool {
	for _, existing := range allowedCaseManagementProjectDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseManagementProjectDataType value.
func (v CaseManagementProjectDataType) Ptr() *CaseManagementProjectDataType {
	return &v
}
