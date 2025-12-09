// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsMetricComputeAggregationType The type of aggregation to use.
type LogsMetricComputeAggregationType string

// List of LogsMetricComputeAggregationType.
const (
	LOGSMETRICCOMPUTEAGGREGATIONTYPE_COUNT        LogsMetricComputeAggregationType = "count"
	LOGSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION LogsMetricComputeAggregationType = "distribution"
)

var allowedLogsMetricComputeAggregationTypeEnumValues = []LogsMetricComputeAggregationType{
	LOGSMETRICCOMPUTEAGGREGATIONTYPE_COUNT,
	LOGSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsMetricComputeAggregationType) GetAllowedValues() []LogsMetricComputeAggregationType {
	return allowedLogsMetricComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsMetricComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsMetricComputeAggregationType(value)
	return nil
}

// NewLogsMetricComputeAggregationTypeFromValue returns a pointer to a valid LogsMetricComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsMetricComputeAggregationTypeFromValue(v string) (*LogsMetricComputeAggregationType, error) {
	ev := LogsMetricComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsMetricComputeAggregationType: valid values are %v", v, allowedLogsMetricComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsMetricComputeAggregationType) IsValid() bool {
	for _, existing := range allowedLogsMetricComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsMetricComputeAggregationType value.
func (v LogsMetricComputeAggregationType) Ptr() *LogsMetricComputeAggregationType {
	return &v
}
