// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationLengthType Operation type.
type LogsArrayProcessorOperationLengthType string

// List of LogsArrayProcessorOperationLengthType.
const (
	LOGSARRAYPROCESSOROPERATIONLENGTHTYPE_LENGTH LogsArrayProcessorOperationLengthType = "length"
)

var allowedLogsArrayProcessorOperationLengthTypeEnumValues = []LogsArrayProcessorOperationLengthType{
	LOGSARRAYPROCESSOROPERATIONLENGTHTYPE_LENGTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArrayProcessorOperationLengthType) GetAllowedValues() []LogsArrayProcessorOperationLengthType {
	return allowedLogsArrayProcessorOperationLengthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArrayProcessorOperationLengthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArrayProcessorOperationLengthType(value)
	return nil
}

// NewLogsArrayProcessorOperationLengthTypeFromValue returns a pointer to a valid LogsArrayProcessorOperationLengthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArrayProcessorOperationLengthTypeFromValue(v string) (*LogsArrayProcessorOperationLengthType, error) {
	ev := LogsArrayProcessorOperationLengthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArrayProcessorOperationLengthType: valid values are %v", v, allowedLogsArrayProcessorOperationLengthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArrayProcessorOperationLengthType) IsValid() bool {
	for _, existing := range allowedLogsArrayProcessorOperationLengthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArrayProcessorOperationLengthType value.
func (v LogsArrayProcessorOperationLengthType) Ptr() *LogsArrayProcessorOperationLengthType {
	return &v
}
