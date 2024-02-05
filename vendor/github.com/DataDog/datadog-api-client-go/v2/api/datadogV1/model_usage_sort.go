// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSort The field to sort by.
type UsageSort string

// List of UsageSort.
const (
	USAGESORT_COMPUTED_ON UsageSort = "computed_on"
	USAGESORT_SIZE        UsageSort = "size"
	USAGESORT_START_DATE  UsageSort = "start_date"
	USAGESORT_END_DATE    UsageSort = "end_date"
)

var allowedUsageSortEnumValues = []UsageSort{
	USAGESORT_COMPUTED_ON,
	USAGESORT_SIZE,
	USAGESORT_START_DATE,
	USAGESORT_END_DATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageSort) GetAllowedValues() []UsageSort {
	return allowedUsageSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageSort(value)
	return nil
}

// NewUsageSortFromValue returns a pointer to a valid UsageSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageSortFromValue(v string) (*UsageSort, error) {
	ev := UsageSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageSort: valid values are %v", v, allowedUsageSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageSort) IsValid() bool {
	for _, existing := range allowedUsageSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageSort value.
func (v UsageSort) Ptr() *UsageSort {
	return &v
}
