// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestType The type of the flaky test from Flaky Test Management.
type FlakyTestType string

// List of FlakyTestType.
const (
	FLAKYTESTTYPE_FLAKY_TEST FlakyTestType = "flaky_test"
)

var allowedFlakyTestTypeEnumValues = []FlakyTestType{
	FLAKYTESTTYPE_FLAKY_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestType) GetAllowedValues() []FlakyTestType {
	return allowedFlakyTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestType(value)
	return nil
}

// NewFlakyTestTypeFromValue returns a pointer to a valid FlakyTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestTypeFromValue(v string) (*FlakyTestType, error) {
	ev := FlakyTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestType: valid values are %v", v, allowedFlakyTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestType) IsValid() bool {
	for _, existing := range allowedFlakyTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestType value.
func (v FlakyTestType) Ptr() *FlakyTestType {
	return &v
}
