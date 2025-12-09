// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryUsersRequestDataType Query users request resource type.
type QueryUsersRequestDataType string

// List of QueryUsersRequestDataType.
const (
	QUERYUSERSREQUESTDATATYPE_QUERY_USERS_REQUEST QueryUsersRequestDataType = "query_users_request"
)

var allowedQueryUsersRequestDataTypeEnumValues = []QueryUsersRequestDataType{
	QUERYUSERSREQUESTDATATYPE_QUERY_USERS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryUsersRequestDataType) GetAllowedValues() []QueryUsersRequestDataType {
	return allowedQueryUsersRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryUsersRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryUsersRequestDataType(value)
	return nil
}

// NewQueryUsersRequestDataTypeFromValue returns a pointer to a valid QueryUsersRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryUsersRequestDataTypeFromValue(v string) (*QueryUsersRequestDataType, error) {
	ev := QueryUsersRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryUsersRequestDataType: valid values are %v", v, allowedQueryUsersRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryUsersRequestDataType) IsValid() bool {
	for _, existing := range allowedQueryUsersRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryUsersRequestDataType value.
func (v QueryUsersRequestDataType) Ptr() *QueryUsersRequestDataType {
	return &v
}
