// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsAggregateSortType The type of sorting algorithm
type LogsAggregateSortType string

// List of LogsAggregateSortType.
const (
	LOGSAGGREGATESORTTYPE_ALPHABETICAL LogsAggregateSortType = "alphabetical"
	LOGSAGGREGATESORTTYPE_MEASURE      LogsAggregateSortType = "measure"
)

var allowedLogsAggregateSortTypeEnumValues = []LogsAggregateSortType{
	LOGSAGGREGATESORTTYPE_ALPHABETICAL,
	LOGSAGGREGATESORTTYPE_MEASURE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsAggregateSortType) GetAllowedValues() []LogsAggregateSortType {
	return allowedLogsAggregateSortTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsAggregateSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsAggregateSortType(value)
	return nil
}

// NewLogsAggregateSortTypeFromValue returns a pointer to a valid LogsAggregateSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsAggregateSortTypeFromValue(v string) (*LogsAggregateSortType, error) {
	ev := LogsAggregateSortType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsAggregateSortType: valid values are %v", v, allowedLogsAggregateSortTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsAggregateSortType) IsValid() bool {
	for _, existing := range allowedLogsAggregateSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsAggregateSortType value.
func (v LogsAggregateSortType) Ptr() *LogsAggregateSortType {
	return &v
}
