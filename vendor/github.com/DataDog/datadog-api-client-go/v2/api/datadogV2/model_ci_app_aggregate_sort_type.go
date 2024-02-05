// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppAggregateSortType The type of sorting algorithm.
type CIAppAggregateSortType string

// List of CIAppAggregateSortType.
const (
	CIAPPAGGREGATESORTTYPE_ALPHABETICAL CIAppAggregateSortType = "alphabetical"
	CIAPPAGGREGATESORTTYPE_MEASURE      CIAppAggregateSortType = "measure"
)

var allowedCIAppAggregateSortTypeEnumValues = []CIAppAggregateSortType{
	CIAPPAGGREGATESORTTYPE_ALPHABETICAL,
	CIAPPAGGREGATESORTTYPE_MEASURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppAggregateSortType) GetAllowedValues() []CIAppAggregateSortType {
	return allowedCIAppAggregateSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppAggregateSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppAggregateSortType(value)
	return nil
}

// NewCIAppAggregateSortTypeFromValue returns a pointer to a valid CIAppAggregateSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppAggregateSortTypeFromValue(v string) (*CIAppAggregateSortType, error) {
	ev := CIAppAggregateSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppAggregateSortType: valid values are %v", v, allowedCIAppAggregateSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppAggregateSortType) IsValid() bool {
	for _, existing := range allowedCIAppAggregateSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppAggregateSortType value.
func (v CIAppAggregateSortType) Ptr() *CIAppAggregateSortType {
	return &v
}
