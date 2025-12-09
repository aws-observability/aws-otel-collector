// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestsSearchSort Parameter for sorting flaky test results. The default sort is by ascending Fully Qualified Name (FQN). The FQN is the concatenation of the test module, suite, and name.
type FlakyTestsSearchSort string

// List of FlakyTestsSearchSort.
const (
	FLAKYTESTSSEARCHSORT_FQN_ASCENDING                      FlakyTestsSearchSort = "fqn"
	FLAKYTESTSSEARCHSORT_FQN_DESCENDING                     FlakyTestsSearchSort = "-fqn"
	FLAKYTESTSSEARCHSORT_FIRST_FLAKED_ASCENDING             FlakyTestsSearchSort = "first_flaked"
	FLAKYTESTSSEARCHSORT_FIRST_FLAKED_DESCENDING            FlakyTestsSearchSort = "-first_flaked"
	FLAKYTESTSSEARCHSORT_LAST_FLAKED_ASCENDING              FlakyTestsSearchSort = "last_flaked"
	FLAKYTESTSSEARCHSORT_LAST_FLAKED_DESCENDING             FlakyTestsSearchSort = "-last_flaked"
	FLAKYTESTSSEARCHSORT_FAILURE_RATE_ASCENDING             FlakyTestsSearchSort = "failure_rate"
	FLAKYTESTSSEARCHSORT_FAILURE_RATE_DESCENDING            FlakyTestsSearchSort = "-failure_rate"
	FLAKYTESTSSEARCHSORT_PIPELINES_FAILED_ASCENDING         FlakyTestsSearchSort = "pipelines_failed"
	FLAKYTESTSSEARCHSORT_PIPELINES_FAILED_DESCENDING        FlakyTestsSearchSort = "-pipelines_failed"
	FLAKYTESTSSEARCHSORT_PIPELINES_DURATION_LOST_ASCENDING  FlakyTestsSearchSort = "pipelines_duration_lost"
	FLAKYTESTSSEARCHSORT_PIPELINES_DURATION_LOST_DESCENDING FlakyTestsSearchSort = "-pipelines_duration_lost"
)

var allowedFlakyTestsSearchSortEnumValues = []FlakyTestsSearchSort{
	FLAKYTESTSSEARCHSORT_FQN_ASCENDING,
	FLAKYTESTSSEARCHSORT_FQN_DESCENDING,
	FLAKYTESTSSEARCHSORT_FIRST_FLAKED_ASCENDING,
	FLAKYTESTSSEARCHSORT_FIRST_FLAKED_DESCENDING,
	FLAKYTESTSSEARCHSORT_LAST_FLAKED_ASCENDING,
	FLAKYTESTSSEARCHSORT_LAST_FLAKED_DESCENDING,
	FLAKYTESTSSEARCHSORT_FAILURE_RATE_ASCENDING,
	FLAKYTESTSSEARCHSORT_FAILURE_RATE_DESCENDING,
	FLAKYTESTSSEARCHSORT_PIPELINES_FAILED_ASCENDING,
	FLAKYTESTSSEARCHSORT_PIPELINES_FAILED_DESCENDING,
	FLAKYTESTSSEARCHSORT_PIPELINES_DURATION_LOST_ASCENDING,
	FLAKYTESTSSEARCHSORT_PIPELINES_DURATION_LOST_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestsSearchSort) GetAllowedValues() []FlakyTestsSearchSort {
	return allowedFlakyTestsSearchSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestsSearchSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestsSearchSort(value)
	return nil
}

// NewFlakyTestsSearchSortFromValue returns a pointer to a valid FlakyTestsSearchSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestsSearchSortFromValue(v string) (*FlakyTestsSearchSort, error) {
	ev := FlakyTestsSearchSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestsSearchSort: valid values are %v", v, allowedFlakyTestsSearchSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestsSearchSort) IsValid() bool {
	for _, existing := range allowedFlakyTestsSearchSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestsSearchSort value.
func (v FlakyTestsSearchSort) Ptr() *FlakyTestsSearchSort {
	return &v
}
