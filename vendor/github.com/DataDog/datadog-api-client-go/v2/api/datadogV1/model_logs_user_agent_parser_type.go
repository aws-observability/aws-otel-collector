// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsUserAgentParserType Type of logs User-Agent parser.
type LogsUserAgentParserType string

// List of LogsUserAgentParserType.
const (
	LOGSUSERAGENTPARSERTYPE_USER_AGENT_PARSER LogsUserAgentParserType = "user-agent-parser"
)

var allowedLogsUserAgentParserTypeEnumValues = []LogsUserAgentParserType{
	LOGSUSERAGENTPARSERTYPE_USER_AGENT_PARSER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsUserAgentParserType) GetAllowedValues() []LogsUserAgentParserType {
	return allowedLogsUserAgentParserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsUserAgentParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsUserAgentParserType(value)
	return nil
}

// NewLogsUserAgentParserTypeFromValue returns a pointer to a valid LogsUserAgentParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsUserAgentParserTypeFromValue(v string) (*LogsUserAgentParserType, error) {
	ev := LogsUserAgentParserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsUserAgentParserType: valid values are %v", v, allowedLogsUserAgentParserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsUserAgentParserType) IsValid() bool {
	for _, existing := range allowedLogsUserAgentParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsUserAgentParserType value.
func (v LogsUserAgentParserType) Ptr() *LogsUserAgentParserType {
	return &v
}
