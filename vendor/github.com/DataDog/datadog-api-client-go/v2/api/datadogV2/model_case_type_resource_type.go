// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseTypeResourceType Case type resource type
type CaseTypeResourceType string

// List of CaseTypeResourceType.
const (
	CASETYPERESOURCETYPE_CASE_TYPE CaseTypeResourceType = "case_type"
)

var allowedCaseTypeResourceTypeEnumValues = []CaseTypeResourceType{
	CASETYPERESOURCETYPE_CASE_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseTypeResourceType) GetAllowedValues() []CaseTypeResourceType {
	return allowedCaseTypeResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseTypeResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseTypeResourceType(value)
	return nil
}

// NewCaseTypeResourceTypeFromValue returns a pointer to a valid CaseTypeResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseTypeResourceTypeFromValue(v string) (*CaseTypeResourceType, error) {
	ev := CaseTypeResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseTypeResourceType: valid values are %v", v, allowedCaseTypeResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseTypeResourceType) IsValid() bool {
	for _, existing := range allowedCaseTypeResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseTypeResourceType value.
func (v CaseTypeResourceType) Ptr() *CaseTypeResourceType {
	return &v
}
