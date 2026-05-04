// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationFlakyTestsManagementPoliciesType JSON:API type for Flaky Tests Management policies response.
// The value must always be `test_optimization_flaky_tests_management_policies`.
type TestOptimizationFlakyTestsManagementPoliciesType string

// List of TestOptimizationFlakyTestsManagementPoliciesType.
const (
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESTYPE_TEST_OPTIMIZATION_FLAKY_TESTS_MANAGEMENT_POLICIES TestOptimizationFlakyTestsManagementPoliciesType = "test_optimization_flaky_tests_management_policies"
)

var allowedTestOptimizationFlakyTestsManagementPoliciesTypeEnumValues = []TestOptimizationFlakyTestsManagementPoliciesType{
	TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESTYPE_TEST_OPTIMIZATION_FLAKY_TESTS_MANAGEMENT_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationFlakyTestsManagementPoliciesType) GetAllowedValues() []TestOptimizationFlakyTestsManagementPoliciesType {
	return allowedTestOptimizationFlakyTestsManagementPoliciesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationFlakyTestsManagementPoliciesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationFlakyTestsManagementPoliciesType(value)
	return nil
}

// NewTestOptimizationFlakyTestsManagementPoliciesTypeFromValue returns a pointer to a valid TestOptimizationFlakyTestsManagementPoliciesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationFlakyTestsManagementPoliciesTypeFromValue(v string) (*TestOptimizationFlakyTestsManagementPoliciesType, error) {
	ev := TestOptimizationFlakyTestsManagementPoliciesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationFlakyTestsManagementPoliciesType: valid values are %v", v, allowedTestOptimizationFlakyTestsManagementPoliciesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationFlakyTestsManagementPoliciesType) IsValid() bool {
	for _, existing := range allowedTestOptimizationFlakyTestsManagementPoliciesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationFlakyTestsManagementPoliciesType value.
func (v TestOptimizationFlakyTestsManagementPoliciesType) Ptr() *TestOptimizationFlakyTestsManagementPoliciesType {
	return &v
}
