// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseStatus Case status
type CaseStatus string

// List of CaseStatus.
const (
	CASESTATUS_OPEN        CaseStatus = "OPEN"
	CASESTATUS_IN_PROGRESS CaseStatus = "IN_PROGRESS"
	CASESTATUS_CLOSED      CaseStatus = "CLOSED"
)

var allowedCaseStatusEnumValues = []CaseStatus{
	CASESTATUS_OPEN,
	CASESTATUS_IN_PROGRESS,
	CASESTATUS_CLOSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseStatus) GetAllowedValues() []CaseStatus {
	return allowedCaseStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseStatus(value)
	return nil
}

// NewCaseStatusFromValue returns a pointer to a valid CaseStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseStatusFromValue(v string) (*CaseStatus, error) {
	ev := CaseStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseStatus: valid values are %v", v, allowedCaseStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseStatus) IsValid() bool {
	for _, existing := range allowedCaseStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseStatus value.
func (v CaseStatus) Ptr() *CaseStatus {
	return &v
}
