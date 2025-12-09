// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorType Type of logs array processor.
type LogsArrayProcessorType string

// List of LogsArrayProcessorType.
const (
	LOGSARRAYPROCESSORTYPE_ARRAY_PROCESSOR LogsArrayProcessorType = "array-processor"
)

var allowedLogsArrayProcessorTypeEnumValues = []LogsArrayProcessorType{
	LOGSARRAYPROCESSORTYPE_ARRAY_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArrayProcessorType) GetAllowedValues() []LogsArrayProcessorType {
	return allowedLogsArrayProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArrayProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArrayProcessorType(value)
	return nil
}

// NewLogsArrayProcessorTypeFromValue returns a pointer to a valid LogsArrayProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArrayProcessorTypeFromValue(v string) (*LogsArrayProcessorType, error) {
	ev := LogsArrayProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArrayProcessorType: valid values are %v", v, allowedLogsArrayProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArrayProcessorType) IsValid() bool {
	for _, existing := range allowedLogsArrayProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArrayProcessorType value.
func (v LogsArrayProcessorType) Ptr() *LogsArrayProcessorType {
	return &v
}
