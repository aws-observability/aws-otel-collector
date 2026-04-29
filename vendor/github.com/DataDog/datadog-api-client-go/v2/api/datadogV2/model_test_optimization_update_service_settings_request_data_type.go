// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationUpdateServiceSettingsRequestDataType JSON:API type for update service settings request.
// The value must always be `test_optimization_update_service_settings_request`.
type TestOptimizationUpdateServiceSettingsRequestDataType string

// List of TestOptimizationUpdateServiceSettingsRequestDataType.
const (
	TESTOPTIMIZATIONUPDATESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_SERVICE_SETTINGS_REQUEST TestOptimizationUpdateServiceSettingsRequestDataType = "test_optimization_update_service_settings_request"
)

var allowedTestOptimizationUpdateServiceSettingsRequestDataTypeEnumValues = []TestOptimizationUpdateServiceSettingsRequestDataType{
	TESTOPTIMIZATIONUPDATESERVICESETTINGSREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_SERVICE_SETTINGS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationUpdateServiceSettingsRequestDataType) GetAllowedValues() []TestOptimizationUpdateServiceSettingsRequestDataType {
	return allowedTestOptimizationUpdateServiceSettingsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationUpdateServiceSettingsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationUpdateServiceSettingsRequestDataType(value)
	return nil
}

// NewTestOptimizationUpdateServiceSettingsRequestDataTypeFromValue returns a pointer to a valid TestOptimizationUpdateServiceSettingsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationUpdateServiceSettingsRequestDataTypeFromValue(v string) (*TestOptimizationUpdateServiceSettingsRequestDataType, error) {
	ev := TestOptimizationUpdateServiceSettingsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationUpdateServiceSettingsRequestDataType: valid values are %v", v, allowedTestOptimizationUpdateServiceSettingsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationUpdateServiceSettingsRequestDataType) IsValid() bool {
	for _, existing := range allowedTestOptimizationUpdateServiceSettingsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationUpdateServiceSettingsRequestDataType value.
func (v TestOptimizationUpdateServiceSettingsRequestDataType) Ptr() *TestOptimizationUpdateServiceSettingsRequestDataType {
	return &v
}
