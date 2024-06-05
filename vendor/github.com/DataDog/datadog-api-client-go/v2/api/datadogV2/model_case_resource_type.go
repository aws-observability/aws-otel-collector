// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseResourceType Case resource type
type CaseResourceType string

// List of CaseResourceType.
const (
	CASERESOURCETYPE_CASE CaseResourceType = "case"
)

var allowedCaseResourceTypeEnumValues = []CaseResourceType{
	CASERESOURCETYPE_CASE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseResourceType) GetAllowedValues() []CaseResourceType {
	return allowedCaseResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseResourceType(value)
	return nil
}

// NewCaseResourceTypeFromValue returns a pointer to a valid CaseResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseResourceTypeFromValue(v string) (*CaseResourceType, error) {
	ev := CaseResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseResourceType: valid values are %v", v, allowedCaseResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseResourceType) IsValid() bool {
	for _, existing := range allowedCaseResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseResourceType value.
func (v CaseResourceType) Ptr() *CaseResourceType {
	return &v
}
