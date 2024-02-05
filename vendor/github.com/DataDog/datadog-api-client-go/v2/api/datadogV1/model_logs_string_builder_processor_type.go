// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsStringBuilderProcessorType Type of logs string builder processor.
type LogsStringBuilderProcessorType string

// List of LogsStringBuilderProcessorType.
const (
	LOGSSTRINGBUILDERPROCESSORTYPE_STRING_BUILDER_PROCESSOR LogsStringBuilderProcessorType = "string-builder-processor"
)

var allowedLogsStringBuilderProcessorTypeEnumValues = []LogsStringBuilderProcessorType{
	LOGSSTRINGBUILDERPROCESSORTYPE_STRING_BUILDER_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsStringBuilderProcessorType) GetAllowedValues() []LogsStringBuilderProcessorType {
	return allowedLogsStringBuilderProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsStringBuilderProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsStringBuilderProcessorType(value)
	return nil
}

// NewLogsStringBuilderProcessorTypeFromValue returns a pointer to a valid LogsStringBuilderProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsStringBuilderProcessorTypeFromValue(v string) (*LogsStringBuilderProcessorType, error) {
	ev := LogsStringBuilderProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsStringBuilderProcessorType: valid values are %v", v, allowedLogsStringBuilderProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsStringBuilderProcessorType) IsValid() bool {
	for _, existing := range allowedLogsStringBuilderProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsStringBuilderProcessorType value.
func (v LogsStringBuilderProcessorType) Ptr() *LogsStringBuilderProcessorType {
	return &v
}
