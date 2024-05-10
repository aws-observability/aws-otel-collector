// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsMetricType The type of the resource. The value should always be logs_metrics.
type LogsMetricType string

// List of LogsMetricType.
const (
	LOGSMETRICTYPE_LOGS_METRICS LogsMetricType = "logs_metrics"
)

var allowedLogsMetricTypeEnumValues = []LogsMetricType{
	LOGSMETRICTYPE_LOGS_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsMetricType) GetAllowedValues() []LogsMetricType {
	return allowedLogsMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsMetricType(value)
	return nil
}

// NewLogsMetricTypeFromValue returns a pointer to a valid LogsMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsMetricTypeFromValue(v string) (*LogsMetricType, error) {
	ev := LogsMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsMetricType: valid values are %v", v, allowedLogsMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsMetricType) IsValid() bool {
	for _, existing := range allowedLogsMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsMetricType value.
func (v LogsMetricType) Ptr() *LogsMetricType {
	return &v
}
