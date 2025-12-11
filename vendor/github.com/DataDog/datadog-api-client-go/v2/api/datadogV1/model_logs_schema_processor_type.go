// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaProcessorType Type of logs schema processor.
type LogsSchemaProcessorType string

// List of LogsSchemaProcessorType.
const (
	LOGSSCHEMAPROCESSORTYPE_SCHEMA_PROCESSOR LogsSchemaProcessorType = "schema-processor"
)

var allowedLogsSchemaProcessorTypeEnumValues = []LogsSchemaProcessorType{
	LOGSSCHEMAPROCESSORTYPE_SCHEMA_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSchemaProcessorType) GetAllowedValues() []LogsSchemaProcessorType {
	return allowedLogsSchemaProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSchemaProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSchemaProcessorType(value)
	return nil
}

// NewLogsSchemaProcessorTypeFromValue returns a pointer to a valid LogsSchemaProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSchemaProcessorTypeFromValue(v string) (*LogsSchemaProcessorType, error) {
	ev := LogsSchemaProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSchemaProcessorType: valid values are %v", v, allowedLogsSchemaProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSchemaProcessorType) IsValid() bool {
	for _, existing := range allowedLogsSchemaProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSchemaProcessorType value.
func (v LogsSchemaProcessorType) Ptr() *LogsSchemaProcessorType {
	return &v
}
