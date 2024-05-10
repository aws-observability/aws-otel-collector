// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppSort Sort parameters when querying events.
type CIAppSort string

// List of CIAppSort.
const (
	CIAPPSORT_TIMESTAMP_ASCENDING  CIAppSort = "timestamp"
	CIAPPSORT_TIMESTAMP_DESCENDING CIAppSort = "-timestamp"
)

var allowedCIAppSortEnumValues = []CIAppSort{
	CIAPPSORT_TIMESTAMP_ASCENDING,
	CIAPPSORT_TIMESTAMP_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppSort) GetAllowedValues() []CIAppSort {
	return allowedCIAppSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppSort(value)
	return nil
}

// NewCIAppSortFromValue returns a pointer to a valid CIAppSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppSortFromValue(v string) (*CIAppSort, error) {
	ev := CIAppSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppSort: valid values are %v", v, allowedCIAppSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppSort) IsValid() bool {
	for _, existing := range allowedCIAppSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppSort value.
func (v CIAppSort) Ptr() *CIAppSort {
	return &v
}
