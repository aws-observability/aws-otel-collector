// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryEventFilteredUsersRequestDataType Query event filtered users request resource type.
type QueryEventFilteredUsersRequestDataType string

// List of QueryEventFilteredUsersRequestDataType.
const (
	QUERYEVENTFILTEREDUSERSREQUESTDATATYPE_QUERY_EVENT_FILTERED_USERS_REQUEST QueryEventFilteredUsersRequestDataType = "query_event_filtered_users_request"
)

var allowedQueryEventFilteredUsersRequestDataTypeEnumValues = []QueryEventFilteredUsersRequestDataType{
	QUERYEVENTFILTEREDUSERSREQUESTDATATYPE_QUERY_EVENT_FILTERED_USERS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryEventFilteredUsersRequestDataType) GetAllowedValues() []QueryEventFilteredUsersRequestDataType {
	return allowedQueryEventFilteredUsersRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryEventFilteredUsersRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryEventFilteredUsersRequestDataType(value)
	return nil
}

// NewQueryEventFilteredUsersRequestDataTypeFromValue returns a pointer to a valid QueryEventFilteredUsersRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryEventFilteredUsersRequestDataTypeFromValue(v string) (*QueryEventFilteredUsersRequestDataType, error) {
	ev := QueryEventFilteredUsersRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryEventFilteredUsersRequestDataType: valid values are %v", v, allowedQueryEventFilteredUsersRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryEventFilteredUsersRequestDataType) IsValid() bool {
	for _, existing := range allowedQueryEventFilteredUsersRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryEventFilteredUsersRequestDataType value.
func (v QueryEventFilteredUsersRequestDataType) Ptr() *QueryEventFilteredUsersRequestDataType {
	return &v
}
