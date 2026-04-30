// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationGetServiceSettingsRequestDataType JSON:API type for get service settings request.
// The value must always be `test_optimization_get_service_settings_request`.
type TestOptimizationGetServiceSettingsRequestDataType string

// List of TestOptimizationGetServiceSettingsRequestDataType.
const (
	TESTOPTIMIZATIONGETSERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_GET_SERVICE_SETTINGS_REQUEST TestOptimizationGetServiceSettingsRequestDataType = "test_optimization_get_service_settings_request"
)

var allowedTestOptimizationGetServiceSettingsRequestDataTypeEnumValues = []TestOptimizationGetServiceSettingsRequestDataType{
	TESTOPTIMIZATIONGETSERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_GET_SERVICE_SETTINGS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationGetServiceSettingsRequestDataType) GetAllowedValues() []TestOptimizationGetServiceSettingsRequestDataType {
	return allowedTestOptimizationGetServiceSettingsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationGetServiceSettingsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationGetServiceSettingsRequestDataType(value)
	return nil
}

// NewTestOptimizationGetServiceSettingsRequestDataTypeFromValue returns a pointer to a valid TestOptimizationGetServiceSettingsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationGetServiceSettingsRequestDataTypeFromValue(v string) (*TestOptimizationGetServiceSettingsRequestDataType, error) {
	ev := TestOptimizationGetServiceSettingsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationGetServiceSettingsRequestDataType: valid values are %v", v, allowedTestOptimizationGetServiceSettingsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationGetServiceSettingsRequestDataType) IsValid() bool {
	for _, existing := range allowedTestOptimizationGetServiceSettingsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationGetServiceSettingsRequestDataType value.
func (v TestOptimizationGetServiceSettingsRequestDataType) Ptr() *TestOptimizationGetServiceSettingsRequestDataType {
	return &v
}
