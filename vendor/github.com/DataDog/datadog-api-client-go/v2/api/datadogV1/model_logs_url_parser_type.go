// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsURLParserType Type of logs URL parser.
type LogsURLParserType string

// List of LogsURLParserType.
const (
	LOGSURLPARSERTYPE_URL_PARSER LogsURLParserType = "url-parser"
)

var allowedLogsURLParserTypeEnumValues = []LogsURLParserType{
	LOGSURLPARSERTYPE_URL_PARSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsURLParserType) GetAllowedValues() []LogsURLParserType {
	return allowedLogsURLParserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsURLParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsURLParserType(value)
	return nil
}

// NewLogsURLParserTypeFromValue returns a pointer to a valid LogsURLParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsURLParserTypeFromValue(v string) (*LogsURLParserType, error) {
	ev := LogsURLParserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsURLParserType: valid values are %v", v, allowedLogsURLParserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsURLParserType) IsValid() bool {
	for _, existing := range allowedLogsURLParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsURLParserType value.
func (v LogsURLParserType) Ptr() *LogsURLParserType {
	return &v
}
