// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDecoderProcessorType Type of logs decoder processor.
type LogsDecoderProcessorType string

// List of LogsDecoderProcessorType.
const (
	LOGSDECODERPROCESSORTYPE_DECODER_PROCESSOR LogsDecoderProcessorType = "decoder-processor"
)

var allowedLogsDecoderProcessorTypeEnumValues = []LogsDecoderProcessorType{
	LOGSDECODERPROCESSORTYPE_DECODER_PROCESSOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsDecoderProcessorType) GetAllowedValues() []LogsDecoderProcessorType {
	return allowedLogsDecoderProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsDecoderProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsDecoderProcessorType(value)
	return nil
}

// NewLogsDecoderProcessorTypeFromValue returns a pointer to a valid LogsDecoderProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsDecoderProcessorTypeFromValue(v string) (*LogsDecoderProcessorType, error) {
	ev := LogsDecoderProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsDecoderProcessorType: valid values are %v", v, allowedLogsDecoderProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsDecoderProcessorType) IsValid() bool {
	for _, existing := range allowedLogsDecoderProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsDecoderProcessorType value.
func (v LogsDecoderProcessorType) Ptr() *LogsDecoderProcessorType {
	return &v
}
