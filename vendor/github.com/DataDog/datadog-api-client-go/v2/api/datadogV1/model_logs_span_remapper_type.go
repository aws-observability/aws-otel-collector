// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSpanRemapperType Type of logs span remapper.
type LogsSpanRemapperType string

// List of LogsSpanRemapperType.
const (
	LOGSSPANREMAPPERTYPE_SPAN_ID_REMAPPER LogsSpanRemapperType = "span-id-remapper"
)

var allowedLogsSpanRemapperTypeEnumValues = []LogsSpanRemapperType{
	LOGSSPANREMAPPERTYPE_SPAN_ID_REMAPPER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsSpanRemapperType) GetAllowedValues() []LogsSpanRemapperType {
	return allowedLogsSpanRemapperTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsSpanRemapperType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsSpanRemapperType(value)
	return nil
}

// NewLogsSpanRemapperTypeFromValue returns a pointer to a valid LogsSpanRemapperType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsSpanRemapperTypeFromValue(v string) (*LogsSpanRemapperType, error) {
	ev := LogsSpanRemapperType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsSpanRemapperType: valid values are %v", v, allowedLogsSpanRemapperTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsSpanRemapperType) IsValid() bool {
	for _, existing := range allowedLogsSpanRemapperTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsSpanRemapperType value.
func (v LogsSpanRemapperType) Ptr() *LogsSpanRemapperType {
	return &v
}
