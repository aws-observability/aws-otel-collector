// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSort Time-ascending `asc` or time-descending `desc` results.
type LogsSort string

// List of LogsSort.
const (
	LOGSSORT_TIME_ASCENDING  LogsSort = "asc"
	LOGSSORT_TIME_DESCENDING LogsSort = "desc"
)

var allowedLogsSortEnumValues = []LogsSort{
	LOGSSORT_TIME_ASCENDING,
	LOGSSORT_TIME_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSort) GetAllowedValues() []LogsSort {
	return allowedLogsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSort(value)
	return nil
}

// NewLogsSortFromValue returns a pointer to a valid LogsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSortFromValue(v string) (*LogsSort, error) {
	ev := LogsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSort: valid values are %v", v, allowedLogsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSort) IsValid() bool {
	for _, existing := range allowedLogsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSort value.
func (v LogsSort) Ptr() *LogsSort {
	return &v
}
