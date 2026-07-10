// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsExcludeAttributeProcessorType Type of logs exclude attribute processor.
type LogsExcludeAttributeProcessorType string

// List of LogsExcludeAttributeProcessorType.
const (
	LOGSEXCLUDEATTRIBUTEPROCESSORTYPE_EXCLUDE_ATTRIBUTE LogsExcludeAttributeProcessorType = "exclude-attribute"
)

var allowedLogsExcludeAttributeProcessorTypeEnumValues = []LogsExcludeAttributeProcessorType{
	LOGSEXCLUDEATTRIBUTEPROCESSORTYPE_EXCLUDE_ATTRIBUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsExcludeAttributeProcessorType) GetAllowedValues() []LogsExcludeAttributeProcessorType {
	return allowedLogsExcludeAttributeProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsExcludeAttributeProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsExcludeAttributeProcessorType(value)
	return nil
}

// NewLogsExcludeAttributeProcessorTypeFromValue returns a pointer to a valid LogsExcludeAttributeProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsExcludeAttributeProcessorTypeFromValue(v string) (*LogsExcludeAttributeProcessorType, error) {
	ev := LogsExcludeAttributeProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsExcludeAttributeProcessorType: valid values are %v", v, allowedLogsExcludeAttributeProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsExcludeAttributeProcessorType) IsValid() bool {
	for _, existing := range allowedLogsExcludeAttributeProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsExcludeAttributeProcessorType value.
func (v LogsExcludeAttributeProcessorType) Ptr() *LogsExcludeAttributeProcessorType {
	return &v
}
