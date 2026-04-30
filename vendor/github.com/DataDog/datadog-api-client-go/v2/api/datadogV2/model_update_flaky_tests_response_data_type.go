// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsResponseDataType The definition of `UpdateFlakyTestsResponseDataType` object.
type UpdateFlakyTestsResponseDataType string

// List of UpdateFlakyTestsResponseDataType.
const (
	UPDATEFLAKYTESTSRESPONSEDATATYPE_UPDATE_FLAKY_TEST_STATE_RESPONSE UpdateFlakyTestsResponseDataType = "update_flaky_test_state_response"
)

var allowedUpdateFlakyTestsResponseDataTypeEnumValues = []UpdateFlakyTestsResponseDataType{
	UPDATEFLAKYTESTSRESPONSEDATATYPE_UPDATE_FLAKY_TEST_STATE_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateFlakyTestsResponseDataType) GetAllowedValues() []UpdateFlakyTestsResponseDataType {
	return allowedUpdateFlakyTestsResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateFlakyTestsResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateFlakyTestsResponseDataType(value)
	return nil
}

// NewUpdateFlakyTestsResponseDataTypeFromValue returns a pointer to a valid UpdateFlakyTestsResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateFlakyTestsResponseDataTypeFromValue(v string) (*UpdateFlakyTestsResponseDataType, error) {
	ev := UpdateFlakyTestsResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateFlakyTestsResponseDataType: valid values are %v", v, allowedUpdateFlakyTestsResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateFlakyTestsResponseDataType) IsValid() bool {
	for _, existing := range allowedUpdateFlakyTestsResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateFlakyTestsResponseDataType value.
func (v UpdateFlakyTestsResponseDataType) Ptr() *UpdateFlakyTestsResponseDataType {
	return &v
}
