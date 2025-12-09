// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansSort Sort parameters when querying spans.
type SpansSort string

// List of SpansSort.
const (
	SPANSSORT_TIMESTAMP_ASCENDING  SpansSort = "timestamp"
	SPANSSORT_TIMESTAMP_DESCENDING SpansSort = "-timestamp"
)

var allowedSpansSortEnumValues = []SpansSort{
	SPANSSORT_TIMESTAMP_ASCENDING,
	SPANSSORT_TIMESTAMP_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansSort) GetAllowedValues() []SpansSort {
	return allowedSpansSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansSort(value)
	return nil
}

// NewSpansSortFromValue returns a pointer to a valid SpansSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansSortFromValue(v string) (*SpansSort, error) {
	ev := SpansSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansSort: valid values are %v", v, allowedSpansSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansSort) IsValid() bool {
	for _, existing := range allowedSpansSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansSort value.
func (v SpansSort) Ptr() *SpansSort {
	return &v
}
