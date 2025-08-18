// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationAppendType Operation type.
type LogsArrayProcessorOperationAppendType string

// List of LogsArrayProcessorOperationAppendType.
const (
	LOGSARRAYPROCESSOROPERATIONAPPENDTYPE_APPEND LogsArrayProcessorOperationAppendType = "append"
)

var allowedLogsArrayProcessorOperationAppendTypeEnumValues = []LogsArrayProcessorOperationAppendType{
	LOGSARRAYPROCESSOROPERATIONAPPENDTYPE_APPEND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArrayProcessorOperationAppendType) GetAllowedValues() []LogsArrayProcessorOperationAppendType {
	return allowedLogsArrayProcessorOperationAppendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArrayProcessorOperationAppendType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArrayProcessorOperationAppendType(value)
	return nil
}

// NewLogsArrayProcessorOperationAppendTypeFromValue returns a pointer to a valid LogsArrayProcessorOperationAppendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArrayProcessorOperationAppendTypeFromValue(v string) (*LogsArrayProcessorOperationAppendType, error) {
	ev := LogsArrayProcessorOperationAppendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArrayProcessorOperationAppendType: valid values are %v", v, allowedLogsArrayProcessorOperationAppendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArrayProcessorOperationAppendType) IsValid() bool {
	for _, existing := range allowedLogsArrayProcessorOperationAppendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArrayProcessorOperationAppendType value.
func (v LogsArrayProcessorOperationAppendType) Ptr() *LogsArrayProcessorOperationAppendType {
	return &v
}
