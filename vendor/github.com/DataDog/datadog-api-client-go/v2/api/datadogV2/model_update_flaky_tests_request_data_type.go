// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFlakyTestsRequestDataType The definition of `UpdateFlakyTestsRequestDataType` object.
type UpdateFlakyTestsRequestDataType string

// List of UpdateFlakyTestsRequestDataType.
const (
	UPDATEFLAKYTESTSREQUESTDATATYPE_UPDATE_FLAKY_TEST_STATE_REQUEST UpdateFlakyTestsRequestDataType = "update_flaky_test_state_request"
)

var allowedUpdateFlakyTestsRequestDataTypeEnumValues = []UpdateFlakyTestsRequestDataType{
	UPDATEFLAKYTESTSREQUESTDATATYPE_UPDATE_FLAKY_TEST_STATE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateFlakyTestsRequestDataType) GetAllowedValues() []UpdateFlakyTestsRequestDataType {
	return allowedUpdateFlakyTestsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateFlakyTestsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateFlakyTestsRequestDataType(value)
	return nil
}

// NewUpdateFlakyTestsRequestDataTypeFromValue returns a pointer to a valid UpdateFlakyTestsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateFlakyTestsRequestDataTypeFromValue(v string) (*UpdateFlakyTestsRequestDataType, error) {
	ev := UpdateFlakyTestsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateFlakyTestsRequestDataType: valid values are %v", v, allowedUpdateFlakyTestsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateFlakyTestsRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateFlakyTestsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateFlakyTestsRequestDataType value.
func (v UpdateFlakyTestsRequestDataType) Ptr() *UpdateFlakyTestsRequestDataType {
	return &v
}
