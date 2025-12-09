// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaRemapperType Type of logs schema remapper.
type LogsSchemaRemapperType string

// List of LogsSchemaRemapperType.
const (
	LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER LogsSchemaRemapperType = "schema-remapper"
)

var allowedLogsSchemaRemapperTypeEnumValues = []LogsSchemaRemapperType{
	LOGSSCHEMAREMAPPERTYPE_SCHEMA_REMAPPER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSchemaRemapperType) GetAllowedValues() []LogsSchemaRemapperType {
	return allowedLogsSchemaRemapperTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSchemaRemapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSchemaRemapperType(value)
	return nil
}

// NewLogsSchemaRemapperTypeFromValue returns a pointer to a valid LogsSchemaRemapperType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSchemaRemapperTypeFromValue(v string) (*LogsSchemaRemapperType, error) {
	ev := LogsSchemaRemapperType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSchemaRemapperType: valid values are %v", v, allowedLogsSchemaRemapperTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSchemaRemapperType) IsValid() bool {
	for _, existing := range allowedLogsSchemaRemapperTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSchemaRemapperType value.
func (v LogsSchemaRemapperType) Ptr() *LogsSchemaRemapperType {
	return &v
}
