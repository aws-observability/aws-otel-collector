// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeletedSuitesRequestType Type for the bulk delete Synthetic suites request, `delete_suites_request`.
type DeletedSuitesRequestType string

// List of DeletedSuitesRequestType.
const (
	DELETEDSUITESREQUESTTYPE_DELETE_SUITES_REQUEST DeletedSuitesRequestType = "delete_suites_request"
)

var allowedDeletedSuitesRequestTypeEnumValues = []DeletedSuitesRequestType{
	DELETEDSUITESREQUESTTYPE_DELETE_SUITES_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeletedSuitesRequestType) GetAllowedValues() []DeletedSuitesRequestType {
	return allowedDeletedSuitesRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeletedSuitesRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeletedSuitesRequestType(value)
	return nil
}

// NewDeletedSuitesRequestTypeFromValue returns a pointer to a valid DeletedSuitesRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeletedSuitesRequestTypeFromValue(v string) (*DeletedSuitesRequestType, error) {
	ev := DeletedSuitesRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeletedSuitesRequestType: valid values are %v", v, allowedDeletedSuitesRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeletedSuitesRequestType) IsValid() bool {
	for _, existing := range allowedDeletedSuitesRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeletedSuitesRequestType value.
func (v DeletedSuitesRequestType) Ptr() *DeletedSuitesRequestType {
	return &v
}
