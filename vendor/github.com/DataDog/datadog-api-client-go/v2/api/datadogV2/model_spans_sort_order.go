// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansSortOrder The order to use, ascending or descending.
type SpansSortOrder string

// List of SpansSortOrder.
const (
	SPANSSORTORDER_ASCENDING  SpansSortOrder = "asc"
	SPANSSORTORDER_DESCENDING SpansSortOrder = "desc"
)

var allowedSpansSortOrderEnumValues = []SpansSortOrder{
	SPANSSORTORDER_ASCENDING,
	SPANSSORTORDER_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansSortOrder) GetAllowedValues() []SpansSortOrder {
	return allowedSpansSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansSortOrder(value)
	return nil
}

// NewSpansSortOrderFromValue returns a pointer to a valid SpansSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansSortOrderFromValue(v string) (*SpansSortOrder, error) {
	ev := SpansSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansSortOrder: valid values are %v", v, allowedSpansSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansSortOrder) IsValid() bool {
	for _, existing := range allowedSpansSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansSortOrder value.
func (v SpansSortOrder) Ptr() *SpansSortOrder {
	return &v
}
