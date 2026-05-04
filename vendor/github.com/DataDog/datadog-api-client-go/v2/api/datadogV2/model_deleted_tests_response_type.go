// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedTestsResponseType Type for the bulk delete Synthetic tests response, `delete_tests`.
type DeletedTestsResponseType string

// List of DeletedTestsResponseType.
const (
	DELETEDTESTSRESPONSETYPE_DELETE_TESTS DeletedTestsResponseType = "delete_tests"
)

var allowedDeletedTestsResponseTypeEnumValues = []DeletedTestsResponseType{
	DELETEDTESTSRESPONSETYPE_DELETE_TESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeletedTestsResponseType) GetAllowedValues() []DeletedTestsResponseType {
	return allowedDeletedTestsResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeletedTestsResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeletedTestsResponseType(value)
	return nil
}

// NewDeletedTestsResponseTypeFromValue returns a pointer to a valid DeletedTestsResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeletedTestsResponseTypeFromValue(v string) (*DeletedTestsResponseType, error) {
	ev := DeletedTestsResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeletedTestsResponseType: valid values are %v", v, allowedDeletedTestsResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeletedTestsResponseType) IsValid() bool {
	for _, existing := range allowedDeletedTestsResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeletedTestsResponseType value.
func (v DeletedTestsResponseType) Ptr() *DeletedTestsResponseType {
	return &v
}
