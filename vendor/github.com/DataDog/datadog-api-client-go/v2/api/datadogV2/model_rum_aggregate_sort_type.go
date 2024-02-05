// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMAggregateSortType The type of sorting algorithm.
type RUMAggregateSortType string

// List of RUMAggregateSortType.
const (
	RUMAGGREGATESORTTYPE_ALPHABETICAL RUMAggregateSortType = "alphabetical"
	RUMAGGREGATESORTTYPE_MEASURE      RUMAggregateSortType = "measure"
)

var allowedRUMAggregateSortTypeEnumValues = []RUMAggregateSortType{
	RUMAGGREGATESORTTYPE_ALPHABETICAL,
	RUMAGGREGATESORTTYPE_MEASURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMAggregateSortType) GetAllowedValues() []RUMAggregateSortType {
	return allowedRUMAggregateSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMAggregateSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMAggregateSortType(value)
	return nil
}

// NewRUMAggregateSortTypeFromValue returns a pointer to a valid RUMAggregateSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMAggregateSortTypeFromValue(v string) (*RUMAggregateSortType, error) {
	ev := RUMAggregateSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMAggregateSortType: valid values are %v", v, allowedRUMAggregateSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMAggregateSortType) IsValid() bool {
	for _, existing := range allowedRUMAggregateSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMAggregateSortType value.
func (v RUMAggregateSortType) Ptr() *RUMAggregateSortType {
	return &v
}
