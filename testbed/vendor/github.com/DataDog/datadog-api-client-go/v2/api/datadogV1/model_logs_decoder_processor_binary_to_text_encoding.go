// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsDecoderProcessorBinaryToTextEncoding The encoding used to represent the binary data.
type LogsDecoderProcessorBinaryToTextEncoding string

// List of LogsDecoderProcessorBinaryToTextEncoding.
const (
	LOGSDECODERPROCESSORBINARYTOTEXTENCODING_BASE64 LogsDecoderProcessorBinaryToTextEncoding = "base64"
	LOGSDECODERPROCESSORBINARYTOTEXTENCODING_BASE16 LogsDecoderProcessorBinaryToTextEncoding = "base16"
)

var allowedLogsDecoderProcessorBinaryToTextEncodingEnumValues = []LogsDecoderProcessorBinaryToTextEncoding{
	LOGSDECODERPROCESSORBINARYTOTEXTENCODING_BASE64,
	LOGSDECODERPROCESSORBINARYTOTEXTENCODING_BASE16,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsDecoderProcessorBinaryToTextEncoding) GetAllowedValues() []LogsDecoderProcessorBinaryToTextEncoding {
	return allowedLogsDecoderProcessorBinaryToTextEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsDecoderProcessorBinaryToTextEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsDecoderProcessorBinaryToTextEncoding(value)
	return nil
}

// NewLogsDecoderProcessorBinaryToTextEncodingFromValue returns a pointer to a valid LogsDecoderProcessorBinaryToTextEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsDecoderProcessorBinaryToTextEncodingFromValue(v string) (*LogsDecoderProcessorBinaryToTextEncoding, error) {
	ev := LogsDecoderProcessorBinaryToTextEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsDecoderProcessorBinaryToTextEncoding: valid values are %v", v, allowedLogsDecoderProcessorBinaryToTextEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsDecoderProcessorBinaryToTextEncoding) IsValid() bool {
	for _, existing := range allowedLogsDecoderProcessorBinaryToTextEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsDecoderProcessorBinaryToTextEncoding value.
func (v LogsDecoderProcessorBinaryToTextEncoding) Ptr() *LogsDecoderProcessorBinaryToTextEncoding {
	return &v
}
