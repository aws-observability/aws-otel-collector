// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogStreamWidgetDefinitionType Type of the log stream widget.
type LogStreamWidgetDefinitionType string

// List of LogStreamWidgetDefinitionType.
const (
	LOGSTREAMWIDGETDEFINITIONTYPE_LOG_STREAM LogStreamWidgetDefinitionType = "log_stream"
)

var allowedLogStreamWidgetDefinitionTypeEnumValues = []LogStreamWidgetDefinitionType{
	LOGSTREAMWIDGETDEFINITIONTYPE_LOG_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogStreamWidgetDefinitionType) GetAllowedValues() []LogStreamWidgetDefinitionType {
	return allowedLogStreamWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogStreamWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogStreamWidgetDefinitionType(value)
	return nil
}

// NewLogStreamWidgetDefinitionTypeFromValue returns a pointer to a valid LogStreamWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogStreamWidgetDefinitionTypeFromValue(v string) (*LogStreamWidgetDefinitionType, error) {
	ev := LogStreamWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogStreamWidgetDefinitionType: valid values are %v", v, allowedLogStreamWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogStreamWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedLogStreamWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogStreamWidgetDefinitionType value.
func (v LogStreamWidgetDefinitionType) Ptr() *LogStreamWidgetDefinitionType {
	return &v
}
