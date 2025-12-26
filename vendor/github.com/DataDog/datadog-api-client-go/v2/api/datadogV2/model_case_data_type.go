// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseDataType Cases resource type.
type CaseDataType string

// List of CaseDataType.
const (
	CASEDATATYPE_CASES CaseDataType = "cases"
)

var allowedCaseDataTypeEnumValues = []CaseDataType{
	CASEDATATYPE_CASES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseDataType) GetAllowedValues() []CaseDataType {
	return allowedCaseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseDataType(value)
	return nil
}

// NewCaseDataTypeFromValue returns a pointer to a valid CaseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseDataTypeFromValue(v string) (*CaseDataType, error) {
	ev := CaseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseDataType: valid values are %v", v, allowedCaseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseDataType) IsValid() bool {
	for _, existing := range allowedCaseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseDataType value.
func (v CaseDataType) Ptr() *CaseDataType {
	return &v
}
