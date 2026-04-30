// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TestOptimizationServiceSettingsType JSON:API type for service settings response.
// The value must always be `test_optimization_service_settings`.
type TestOptimizationServiceSettingsType string

// List of TestOptimizationServiceSettingsType.
const (
	TESTOPTIMIZATIONSERVICESETTINGSTYPE_TEST_OPTIMIZATION_SERVICE_SETTINGS TestOptimizationServiceSettingsType = "test_optimization_service_settings"
)

var allowedTestOptimizationServiceSettingsTypeEnumValues = []TestOptimizationServiceSettingsType{
	TESTOPTIMIZATIONSERVICESETTINGSTYPE_TEST_OPTIMIZATION_SERVICE_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TestOptimizationServiceSettingsType) GetAllowedValues() []TestOptimizationServiceSettingsType {
	return allowedTestOptimizationServiceSettingsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TestOptimizationServiceSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TestOptimizationServiceSettingsType(value)
	return nil
}

// NewTestOptimizationServiceSettingsTypeFromValue returns a pointer to a valid TestOptimizationServiceSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTestOptimizationServiceSettingsTypeFromValue(v string) (*TestOptimizationServiceSettingsType, error) {
	ev := TestOptimizationServiceSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TestOptimizationServiceSettingsType: valid values are %v", v, allowedTestOptimizationServiceSettingsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TestOptimizationServiceSettingsType) IsValid() bool {
	for _, existing := range allowedTestOptimizationServiceSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestOptimizationServiceSettingsType value.
func (v TestOptimizationServiceSettingsType) Ptr() *TestOptimizationServiceSettingsType {
	return &v
}
