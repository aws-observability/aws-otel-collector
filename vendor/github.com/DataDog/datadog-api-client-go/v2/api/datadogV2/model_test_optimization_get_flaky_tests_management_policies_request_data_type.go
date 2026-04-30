// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType JSON:API type for get Flaky Tests Management policies request.
// The value must always be `test_optimization_get_flaky_tests_management_policies_request`.
type TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType string

// List of TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType.
const (
	TESTOPTIMIZATIONGETFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_GET_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType = "test_optimization_get_flaky_tests_management_policies_request"
)

var allowedTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeEnumValues = []TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType{
	TESTOPTIMIZATIONGETFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_GET_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType) GetAllowedValues() []TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType {
	return allowedTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType(value)
	return nil
}

// NewTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeFromValue returns a pointer to a valid TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeFromValue(v string) (*TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType, error) {
	ev := TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType: valid values are %v", v, allowedTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType) IsValid() bool {
	for _, existing := range allowedTestOptimizationGetFlakyTestsManagementPoliciesRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType value.
func (v TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType) Ptr() *TestOptimizationGetFlakyTestsManagementPoliciesRequestDataType {
	return &v
}
