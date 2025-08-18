// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArrayProcessorOperationSelectType Operation type.
type LogsArrayProcessorOperationSelectType string

// List of LogsArrayProcessorOperationSelectType.
const (
	LOGSARRAYPROCESSOROPERATIONSELECTTYPE_SELECT LogsArrayProcessorOperationSelectType = "select"
)

var allowedLogsArrayProcessorOperationSelectTypeEnumValues = []LogsArrayProcessorOperationSelectType{
	LOGSARRAYPROCESSOROPERATIONSELECTTYPE_SELECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsArrayProcessorOperationSelectType) GetAllowedValues() []LogsArrayProcessorOperationSelectType {
	return allowedLogsArrayProcessorOperationSelectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsArrayProcessorOperationSelectType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsArrayProcessorOperationSelectType(value)
	return nil
}

// NewLogsArrayProcessorOperationSelectTypeFromValue returns a pointer to a valid LogsArrayProcessorOperationSelectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsArrayProcessorOperationSelectTypeFromValue(v string) (*LogsArrayProcessorOperationSelectType, error) {
	ev := LogsArrayProcessorOperationSelectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsArrayProcessorOperationSelectType: valid values are %v", v, allowedLogsArrayProcessorOperationSelectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsArrayProcessorOperationSelectType) IsValid() bool {
	for _, existing := range allowedLogsArrayProcessorOperationSelectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsArrayProcessorOperationSelectType value.
func (v LogsArrayProcessorOperationSelectType) Ptr() *LogsArrayProcessorOperationSelectType {
	return &v
}
