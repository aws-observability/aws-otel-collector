// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppSortOrder The order to use, ascending or descending.
type CIAppSortOrder string

// List of CIAppSortOrder.
const (
	CIAPPSORTORDER_ASCENDING  CIAppSortOrder = "asc"
	CIAPPSORTORDER_DESCENDING CIAppSortOrder = "desc"
)

var allowedCIAppSortOrderEnumValues = []CIAppSortOrder{
	CIAPPSORTORDER_ASCENDING,
	CIAPPSORTORDER_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppSortOrder) GetAllowedValues() []CIAppSortOrder {
	return allowedCIAppSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppSortOrder(value)
	return nil
}

// NewCIAppSortOrderFromValue returns a pointer to a valid CIAppSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppSortOrderFromValue(v string) (*CIAppSortOrder, error) {
	ev := CIAppSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppSortOrder: valid values are %v", v, allowedCIAppSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppSortOrder) IsValid() bool {
	for _, existing := range allowedCIAppSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppSortOrder value.
func (v CIAppSortOrder) Ptr() *CIAppSortOrder {
	return &v
}
