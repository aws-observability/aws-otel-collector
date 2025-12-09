// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataType List connections response resource type.
type ListConnectionsResponseDataType string

// List of ListConnectionsResponseDataType.
const (
	LISTCONNECTIONSRESPONSEDATATYPE_LIST_CONNECTIONS_RESPONSE ListConnectionsResponseDataType = "list_connections_response"
)

var allowedListConnectionsResponseDataTypeEnumValues = []ListConnectionsResponseDataType{
	LISTCONNECTIONSRESPONSEDATATYPE_LIST_CONNECTIONS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListConnectionsResponseDataType) GetAllowedValues() []ListConnectionsResponseDataType {
	return allowedListConnectionsResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListConnectionsResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListConnectionsResponseDataType(value)
	return nil
}

// NewListConnectionsResponseDataTypeFromValue returns a pointer to a valid ListConnectionsResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListConnectionsResponseDataTypeFromValue(v string) (*ListConnectionsResponseDataType, error) {
	ev := ListConnectionsResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListConnectionsResponseDataType: valid values are %v", v, allowedListConnectionsResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListConnectionsResponseDataType) IsValid() bool {
	for _, existing := range allowedListConnectionsResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListConnectionsResponseDataType value.
func (v ListConnectionsResponseDataType) Ptr() *ListConnectionsResponseDataType {
	return &v
}
