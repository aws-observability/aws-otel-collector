// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QuerySortOrder Direction of sort.
type QuerySortOrder string

// List of QuerySortOrder.
const (
	QUERYSORTORDER_ASC  QuerySortOrder = "asc"
	QUERYSORTORDER_DESC QuerySortOrder = "desc"
)

var allowedQuerySortOrderEnumValues = []QuerySortOrder{
	QUERYSORTORDER_ASC,
	QUERYSORTORDER_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QuerySortOrder) GetAllowedValues() []QuerySortOrder {
	return allowedQuerySortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QuerySortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QuerySortOrder(value)
	return nil
}

// NewQuerySortOrderFromValue returns a pointer to a valid QuerySortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQuerySortOrderFromValue(v string) (*QuerySortOrder, error) {
	ev := QuerySortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QuerySortOrder: valid values are %v", v, allowedQuerySortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QuerySortOrder) IsValid() bool {
	for _, existing := range allowedQuerySortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuerySortOrder value.
func (v QuerySortOrder) Ptr() *QuerySortOrder {
	return &v
}
