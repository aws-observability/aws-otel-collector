// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateSortType The type of sorting algorithm.
type SpansAggregateSortType string

// List of SpansAggregateSortType.
const (
	SPANSAGGREGATESORTTYPE_ALPHABETICAL SpansAggregateSortType = "alphabetical"
	SPANSAGGREGATESORTTYPE_MEASURE      SpansAggregateSortType = "measure"
)

var allowedSpansAggregateSortTypeEnumValues = []SpansAggregateSortType{
	SPANSAGGREGATESORTTYPE_ALPHABETICAL,
	SPANSAGGREGATESORTTYPE_MEASURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansAggregateSortType) GetAllowedValues() []SpansAggregateSortType {
	return allowedSpansAggregateSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansAggregateSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansAggregateSortType(value)
	return nil
}

// NewSpansAggregateSortTypeFromValue returns a pointer to a valid SpansAggregateSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansAggregateSortTypeFromValue(v string) (*SpansAggregateSortType, error) {
	ev := SpansAggregateSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansAggregateSortType: valid values are %v", v, allowedSpansAggregateSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansAggregateSortType) IsValid() bool {
	for _, existing := range allowedSpansAggregateSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansAggregateSortType value.
func (v SpansAggregateSortType) Ptr() *SpansAggregateSortType {
	return &v
}
