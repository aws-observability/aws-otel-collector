// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseViewResourceType JSON:API resource type for case views.
type CaseViewResourceType string

// List of CaseViewResourceType.
const (
	CASEVIEWRESOURCETYPE_VIEW CaseViewResourceType = "view"
)

var allowedCaseViewResourceTypeEnumValues = []CaseViewResourceType{
	CASEVIEWRESOURCETYPE_VIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseViewResourceType) GetAllowedValues() []CaseViewResourceType {
	return allowedCaseViewResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseViewResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseViewResourceType(value)
	return nil
}

// NewCaseViewResourceTypeFromValue returns a pointer to a valid CaseViewResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseViewResourceTypeFromValue(v string) (*CaseViewResourceType, error) {
	ev := CaseViewResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseViewResourceType: valid values are %v", v, allowedCaseViewResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseViewResourceType) IsValid() bool {
	for _, existing := range allowedCaseViewResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseViewResourceType value.
func (v CaseViewResourceType) Ptr() *CaseViewResourceType {
	return &v
}
