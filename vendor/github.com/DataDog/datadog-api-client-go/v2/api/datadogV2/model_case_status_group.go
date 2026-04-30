// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseStatusGroup Status group of the case.
type CaseStatusGroup string

// List of CaseStatusGroup.
const (
	CASESTATUSGROUP_SG_OPEN        CaseStatusGroup = "SG_OPEN"
	CASESTATUSGROUP_SG_IN_PROGRESS CaseStatusGroup = "SG_IN_PROGRESS"
	CASESTATUSGROUP_SG_CLOSED      CaseStatusGroup = "SG_CLOSED"
)

var allowedCaseStatusGroupEnumValues = []CaseStatusGroup{
	CASESTATUSGROUP_SG_OPEN,
	CASESTATUSGROUP_SG_IN_PROGRESS,
	CASESTATUSGROUP_SG_CLOSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseStatusGroup) GetAllowedValues() []CaseStatusGroup {
	return allowedCaseStatusGroupEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseStatusGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseStatusGroup(value)
	return nil
}

// NewCaseStatusGroupFromValue returns a pointer to a valid CaseStatusGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseStatusGroupFromValue(v string) (*CaseStatusGroup, error) {
	ev := CaseStatusGroup(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseStatusGroup: valid values are %v", v, allowedCaseStatusGroupEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseStatusGroup) IsValid() bool {
	for _, existing := range allowedCaseStatusGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseStatusGroup value.
func (v CaseStatusGroup) Ptr() *CaseStatusGroup {
	return &v
}
