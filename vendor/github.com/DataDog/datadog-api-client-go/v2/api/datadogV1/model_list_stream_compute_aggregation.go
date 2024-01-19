// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamComputeAggregation Aggregation value.
type ListStreamComputeAggregation string

// List of ListStreamComputeAggregation.
const (
	LISTSTREAMCOMPUTEAGGREGATION_COUNT         ListStreamComputeAggregation = "count"
	LISTSTREAMCOMPUTEAGGREGATION_CARDINALITY   ListStreamComputeAggregation = "cardinality"
	LISTSTREAMCOMPUTEAGGREGATION_MEDIAN        ListStreamComputeAggregation = "median"
	LISTSTREAMCOMPUTEAGGREGATION_PC75          ListStreamComputeAggregation = "pc75"
	LISTSTREAMCOMPUTEAGGREGATION_PC90          ListStreamComputeAggregation = "pc90"
	LISTSTREAMCOMPUTEAGGREGATION_PC95          ListStreamComputeAggregation = "pc95"
	LISTSTREAMCOMPUTEAGGREGATION_PC98          ListStreamComputeAggregation = "pc98"
	LISTSTREAMCOMPUTEAGGREGATION_PC99          ListStreamComputeAggregation = "pc99"
	LISTSTREAMCOMPUTEAGGREGATION_SUM           ListStreamComputeAggregation = "sum"
	LISTSTREAMCOMPUTEAGGREGATION_MIN           ListStreamComputeAggregation = "min"
	LISTSTREAMCOMPUTEAGGREGATION_MAX           ListStreamComputeAggregation = "max"
	LISTSTREAMCOMPUTEAGGREGATION_AVG           ListStreamComputeAggregation = "avg"
	LISTSTREAMCOMPUTEAGGREGATION_EARLIEST      ListStreamComputeAggregation = "earliest"
	LISTSTREAMCOMPUTEAGGREGATION_LATEST        ListStreamComputeAggregation = "latest"
	LISTSTREAMCOMPUTEAGGREGATION_MOST_FREQUENT ListStreamComputeAggregation = "most_frequent"
)

var allowedListStreamComputeAggregationEnumValues = []ListStreamComputeAggregation{
	LISTSTREAMCOMPUTEAGGREGATION_COUNT,
	LISTSTREAMCOMPUTEAGGREGATION_CARDINALITY,
	LISTSTREAMCOMPUTEAGGREGATION_MEDIAN,
	LISTSTREAMCOMPUTEAGGREGATION_PC75,
	LISTSTREAMCOMPUTEAGGREGATION_PC90,
	LISTSTREAMCOMPUTEAGGREGATION_PC95,
	LISTSTREAMCOMPUTEAGGREGATION_PC98,
	LISTSTREAMCOMPUTEAGGREGATION_PC99,
	LISTSTREAMCOMPUTEAGGREGATION_SUM,
	LISTSTREAMCOMPUTEAGGREGATION_MIN,
	LISTSTREAMCOMPUTEAGGREGATION_MAX,
	LISTSTREAMCOMPUTEAGGREGATION_AVG,
	LISTSTREAMCOMPUTEAGGREGATION_EARLIEST,
	LISTSTREAMCOMPUTEAGGREGATION_LATEST,
	LISTSTREAMCOMPUTEAGGREGATION_MOST_FREQUENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamComputeAggregation) GetAllowedValues() []ListStreamComputeAggregation {
	return allowedListStreamComputeAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamComputeAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamComputeAggregation(value)
	return nil
}

// NewListStreamComputeAggregationFromValue returns a pointer to a valid ListStreamComputeAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamComputeAggregationFromValue(v string) (*ListStreamComputeAggregation, error) {
	ev := ListStreamComputeAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamComputeAggregation: valid values are %v", v, allowedListStreamComputeAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamComputeAggregation) IsValid() bool {
	for _, existing := range allowedListStreamComputeAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamComputeAggregation value.
func (v ListStreamComputeAggregation) Ptr() *ListStreamComputeAggregation {
	return &v
}
