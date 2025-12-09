// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestsSearchRequestDataType The definition of `FlakyTestsSearchRequestDataType` object.
type FlakyTestsSearchRequestDataType string

// List of FlakyTestsSearchRequestDataType.
const (
	FLAKYTESTSSEARCHREQUESTDATATYPE_SEARCH_FLAKY_TESTS_REQUEST FlakyTestsSearchRequestDataType = "search_flaky_tests_request"
)

var allowedFlakyTestsSearchRequestDataTypeEnumValues = []FlakyTestsSearchRequestDataType{
	FLAKYTESTSSEARCHREQUESTDATATYPE_SEARCH_FLAKY_TESTS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestsSearchRequestDataType) GetAllowedValues() []FlakyTestsSearchRequestDataType {
	return allowedFlakyTestsSearchRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestsSearchRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestsSearchRequestDataType(value)
	return nil
}

// NewFlakyTestsSearchRequestDataTypeFromValue returns a pointer to a valid FlakyTestsSearchRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestsSearchRequestDataTypeFromValue(v string) (*FlakyTestsSearchRequestDataType, error) {
	ev := FlakyTestsSearchRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestsSearchRequestDataType: valid values are %v", v, allowedFlakyTestsSearchRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestsSearchRequestDataType) IsValid() bool {
	for _, existing := range allowedFlakyTestsSearchRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestsSearchRequestDataType value.
func (v FlakyTestsSearchRequestDataType) Ptr() *FlakyTestsSearchRequestDataType {
	return &v
}
