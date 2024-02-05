// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SortDirection The direction to sort by.
type SortDirection string

// List of SortDirection.
const (
	SORTDIRECTION_DESC SortDirection = "desc"
	SORTDIRECTION_ASC  SortDirection = "asc"
)

var allowedSortDirectionEnumValues = []SortDirection{
	SORTDIRECTION_DESC,
	SORTDIRECTION_ASC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SortDirection) GetAllowedValues() []SortDirection {
	return allowedSortDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SortDirection(value)
	return nil
}

// NewSortDirectionFromValue returns a pointer to a valid SortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSortDirectionFromValue(v string) (*SortDirection, error) {
	ev := SortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SortDirection: valid values are %v", v, allowedSortDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SortDirection) IsValid() bool {
	for _, existing := range allowedSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortDirection value.
func (v SortDirection) Ptr() *SortDirection {
	return &v
}
