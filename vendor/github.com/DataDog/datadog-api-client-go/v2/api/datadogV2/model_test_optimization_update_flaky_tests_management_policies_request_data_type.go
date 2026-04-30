// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType JSON:API type for update Flaky Tests Management policies request.
// The value must always be `test_optimization_update_flaky_tests_management_policies_request`.
type TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType string

// List of TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType.
const (
	TESTOPTIMIZATIONUPDATEFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType = "test_optimization_update_flaky_tests_management_policies_request"
)

var allowedTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeEnumValues = []TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType{
	TESTOPTIMIZATIONUPDATEFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType) GetAllowedValues() []TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType {
	return allowedTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType(value)
	return nil
}

// NewTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeFromValue returns a pointer to a valid TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeFromValue(v string) (*TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType, error) {
	ev := TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType: valid values are %v", v, allowedTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType) IsValid() bool {
	for _, existing := range allowedTestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType value.
func (v TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType) Ptr() *TestOptimizationUpdateFlakyTestsManagementPoliciesRequestDataType {
	return &v
}
