// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsRestrictionQueriesType Restriction query resource type.
type LogsRestrictionQueriesType string

// List of LogsRestrictionQueriesType.
const (
	LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES LogsRestrictionQueriesType = "logs_restriction_queries"
)

var allowedLogsRestrictionQueriesTypeEnumValues = []LogsRestrictionQueriesType{
	LOGSRESTRICTIONQUERIESTYPE_LOGS_RESTRICTION_QUERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsRestrictionQueriesType) GetAllowedValues() []LogsRestrictionQueriesType {
	return allowedLogsRestrictionQueriesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsRestrictionQueriesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsRestrictionQueriesType(value)
	return nil
}

// NewLogsRestrictionQueriesTypeFromValue returns a pointer to a valid LogsRestrictionQueriesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsRestrictionQueriesTypeFromValue(v string) (*LogsRestrictionQueriesType, error) {
	ev := LogsRestrictionQueriesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsRestrictionQueriesType: valid values are %v", v, allowedLogsRestrictionQueriesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsRestrictionQueriesType) IsValid() bool {
	for _, existing := range allowedLogsRestrictionQueriesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsRestrictionQueriesType value.
func (v LogsRestrictionQueriesType) Ptr() *LogsRestrictionQueriesType {
	return &v
}
