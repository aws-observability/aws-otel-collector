// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedTestsRequestType Type for the bulk delete Synthetic tests request, `delete_tests_request`.
type DeletedTestsRequestType string

// List of DeletedTestsRequestType.
const (
	DELETEDTESTSREQUESTTYPE_DELETE_TESTS_REQUEST DeletedTestsRequestType = "delete_tests_request"
)

var allowedDeletedTestsRequestTypeEnumValues = []DeletedTestsRequestType{
	DELETEDTESTSREQUESTTYPE_DELETE_TESTS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeletedTestsRequestType) GetAllowedValues() []DeletedTestsRequestType {
	return allowedDeletedTestsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeletedTestsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeletedTestsRequestType(value)
	return nil
}

// NewDeletedTestsRequestTypeFromValue returns a pointer to a valid DeletedTestsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeletedTestsRequestTypeFromValue(v string) (*DeletedTestsRequestType, error) {
	ev := DeletedTestsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeletedTestsRequestType: valid values are %v", v, allowedDeletedTestsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeletedTestsRequestType) IsValid() bool {
	for _, existing := range allowedDeletedTestsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeletedTestsRequestType value.
func (v DeletedTestsRequestType) Ptr() *DeletedTestsRequestType {
	return &v
}
