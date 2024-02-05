// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsMetricResponseComputeAggregationType The type of aggregation to use.
type LogsMetricResponseComputeAggregationType string

// List of LogsMetricResponseComputeAggregationType.
const (
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT        LogsMetricResponseComputeAggregationType = "count"
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION LogsMetricResponseComputeAggregationType = "distribution"
)

var allowedLogsMetricResponseComputeAggregationTypeEnumValues = []LogsMetricResponseComputeAggregationType{
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT,
	LOGSMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsMetricResponseComputeAggregationType) GetAllowedValues() []LogsMetricResponseComputeAggregationType {
	return allowedLogsMetricResponseComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsMetricResponseComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsMetricResponseComputeAggregationType(value)
	return nil
}

// NewLogsMetricResponseComputeAggregationTypeFromValue returns a pointer to a valid LogsMetricResponseComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsMetricResponseComputeAggregationTypeFromValue(v string) (*LogsMetricResponseComputeAggregationType, error) {
	ev := LogsMetricResponseComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsMetricResponseComputeAggregationType: valid values are %v", v, allowedLogsMetricResponseComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsMetricResponseComputeAggregationType) IsValid() bool {
	for _, existing := range allowedLogsMetricResponseComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsMetricResponseComputeAggregationType value.
func (v LogsMetricResponseComputeAggregationType) Ptr() *LogsMetricResponseComputeAggregationType {
	return &v
}
