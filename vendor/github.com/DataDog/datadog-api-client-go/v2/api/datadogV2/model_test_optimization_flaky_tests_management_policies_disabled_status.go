// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesDisabledStatus Test status that the disable policy applies to.
// Must be either `active` or `quarantined`.
type TestOptimizationFlakyTestsManagementPoliciesDisabledStatus string

// List of TestOptimizationFlakyTestsManagementPoliciesDisabledStatus.
const (
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_ACTIVE      TestOptimizationFlakyTestsManagementPoliciesDisabledStatus = "active"
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_QUARANTINED TestOptimizationFlakyTestsManagementPoliciesDisabledStatus = "quarantined"
)

var allowedTestOptimizationFlakyTestsManagementPoliciesDisabledStatusEnumValues = []TestOptimizationFlakyTestsManagementPoliciesDisabledStatus{
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_ACTIVE,
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_QUARANTINED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationFlakyTestsManagementPoliciesDisabledStatus) GetAllowedValues() []TestOptimizationFlakyTestsManagementPoliciesDisabledStatus {
	return allowedTestOptimizationFlakyTestsManagementPoliciesDisabledStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationFlakyTestsManagementPoliciesDisabledStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationFlakyTestsManagementPoliciesDisabledStatus(value)
	return nil
}

// NewTestOptimizationFlakyTestsManagementPoliciesDisabledStatusFromValue returns a pointer to a valid TestOptimizationFlakyTestsManagementPoliciesDisabledStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationFlakyTestsManagementPoliciesDisabledStatusFromValue(v string) (*TestOptimizationFlakyTestsManagementPoliciesDisabledStatus, error) {
	ev := TestOptimizationFlakyTestsManagementPoliciesDisabledStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationFlakyTestsManagementPoliciesDisabledStatus: valid values are %v", v, allowedTestOptimizationFlakyTestsManagementPoliciesDisabledStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationFlakyTestsManagementPoliciesDisabledStatus) IsValid() bool {
	for _, existing := range allowedTestOptimizationFlakyTestsManagementPoliciesDisabledStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationFlakyTestsManagementPoliciesDisabledStatus value.
func (v TestOptimizationFlakyTestsManagementPoliciesDisabledStatus) Ptr() *TestOptimizationFlakyTestsManagementPoliciesDisabledStatus {
	return &v
}
