// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDecoderProcessorInputRepresentation The original representation of input string.
type LogsDecoderProcessorInputRepresentation string

// List of LogsDecoderProcessorInputRepresentation.
const (
	LOGSDECODERPROCESSORINPUTREPRESENTATION_UTF_8   LogsDecoderProcessorInputRepresentation = "utf_8"
	LOGSDECODERPROCESSORINPUTREPRESENTATION_INTEGER LogsDecoderProcessorInputRepresentation = "integer"
)

var allowedLogsDecoderProcessorInputRepresentationEnumValues = []LogsDecoderProcessorInputRepresentation{
	LOGSDECODERPROCESSORINPUTREPRESENTATION_UTF_8,
	LOGSDECODERPROCESSORINPUTREPRESENTATION_INTEGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsDecoderProcessorInputRepresentation) GetAllowedValues() []LogsDecoderProcessorInputRepresentation {
	return allowedLogsDecoderProcessorInputRepresentationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsDecoderProcessorInputRepresentation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsDecoderProcessorInputRepresentation(value)
	return nil
}

// NewLogsDecoderProcessorInputRepresentationFromValue returns a pointer to a valid LogsDecoderProcessorInputRepresentation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsDecoderProcessorInputRepresentationFromValue(v string) (*LogsDecoderProcessorInputRepresentation, error) {
	ev := LogsDecoderProcessorInputRepresentation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsDecoderProcessorInputRepresentation: valid values are %v", v, allowedLogsDecoderProcessorInputRepresentationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsDecoderProcessorInputRepresentation) IsValid() bool {
	for _, existing := range allowedLogsDecoderProcessorInputRepresentationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsDecoderProcessorInputRepresentation value.
func (v LogsDecoderProcessorInputRepresentation) Ptr() *LogsDecoderProcessorInputRepresentation {
	return &v
}
