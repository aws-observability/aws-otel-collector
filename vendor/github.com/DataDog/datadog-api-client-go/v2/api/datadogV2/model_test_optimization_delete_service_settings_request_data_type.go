// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationDeleteServiceSettingsRequestDataType JSON:API type for delete service settings request.
// The value must always be `test_optimization_delete_service_settings_request`.
type TestOptimizationDeleteServiceSettingsRequestDataType string

// List of TestOptimizationDeleteServiceSettingsRequestDataType.
const (
	TESTOPTIMIZATIONDELETESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_DELETE_SERVICE_SETTINGS_REQUEST TestOptimizationDeleteServiceSettingsRequestDataType = "test_optimization_delete_service_settings_request"
)

var allowedTestOptimizationDeleteServiceSettingsRequestDataTypeEnumValues = []TestOptimizationDeleteServiceSettingsRequestDataType{
	TESTOPTIMIZATIONDELETESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_DELETE_SERVICE_SETTINGS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationDeleteServiceSettingsRequestDataType) GetAllowedValues() []TestOptimizationDeleteServiceSettingsRequestDataType {
	return allowedTestOptimizationDeleteServiceSettingsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationDeleteServiceSettingsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationDeleteServiceSettingsRequestDataType(value)
	return nil
}

// NewTestOptimizationDeleteServiceSettingsRequestDataTypeFromValue returns a pointer to a valid TestOptimizationDeleteServiceSettingsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationDeleteServiceSettingsRequestDataTypeFromValue(v string) (*TestOptimizationDeleteServiceSettingsRequestDataType, error) {
	ev := TestOptimizationDeleteServiceSettingsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationDeleteServiceSettingsRequestDataType: valid values are %v", v, allowedTestOptimizationDeleteServiceSettingsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationDeleteServiceSettingsRequestDataType) IsValid() bool {
	for _, existing := range allowedTestOptimizationDeleteServiceSettingsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationDeleteServiceSettingsRequestDataType value.
func (v TestOptimizationDeleteServiceSettingsRequestDataType) Ptr() *TestOptimizationDeleteServiceSettingsRequestDataType {
	return &v
}
