// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryResponseDataType Query response resource type.
type QueryResponseDataType string

// List of QueryResponseDataType.
const (
	QUERYRESPONSEDATATYPE_QUERY_RESPONSE QueryResponseDataType = "query_response"
)

var allowedQueryResponseDataTypeEnumValues = []QueryResponseDataType{
	QUERYRESPONSEDATATYPE_QUERY_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryResponseDataType) GetAllowedValues() []QueryResponseDataType {
	return allowedQueryResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryResponseDataType(value)
	return nil
}

// NewQueryResponseDataTypeFromValue returns a pointer to a valid QueryResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryResponseDataTypeFromValue(v string) (*QueryResponseDataType, error) {
	ev := QueryResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryResponseDataType: valid values are %v", v, allowedQueryResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryResponseDataType) IsValid() bool {
	for _, existing := range allowedQueryResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryResponseDataType value.
func (v QueryResponseDataType) Ptr() *QueryResponseDataType {
	return &v
}
