// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsComputeType The type of compute
type LogsComputeType string

// List of LogsComputeType.
const (
	LOGSCOMPUTETYPE_TIMESERIES LogsComputeType = "timeseries"
	LOGSCOMPUTETYPE_TOTAL      LogsComputeType = "total"
)

var allowedLogsComputeTypeEnumValues = []LogsComputeType{
	LOGSCOMPUTETYPE_TIMESERIES,
	LOGSCOMPUTETYPE_TOTAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsComputeType) GetAllowedValues() []LogsComputeType {
	return allowedLogsComputeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsComputeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsComputeType(value)
	return nil
}

// NewLogsComputeTypeFromValue returns a pointer to a valid LogsComputeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsComputeTypeFromValue(v string) (*LogsComputeType, error) {
	ev := LogsComputeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsComputeType: valid values are %v", v, allowedLogsComputeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsComputeType) IsValid() bool {
	for _, existing := range allowedLogsComputeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsComputeType value.
func (v LogsComputeType) Ptr() *LogsComputeType {
	return &v
}
