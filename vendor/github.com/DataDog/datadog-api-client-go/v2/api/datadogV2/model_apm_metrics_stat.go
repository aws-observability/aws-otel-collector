// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricsStat The APM metric statistic to query.
type ApmMetricsStat string

// List of ApmMetricsStat.
const (
	APMMETRICSSTAT_ERROR_RATE           ApmMetricsStat = "error_rate"
	APMMETRICSSTAT_ERRORS               ApmMetricsStat = "errors"
	APMMETRICSSTAT_ERRORS_PER_SECOND    ApmMetricsStat = "errors_per_second"
	APMMETRICSSTAT_HITS                 ApmMetricsStat = "hits"
	APMMETRICSSTAT_HITS_PER_SECOND      ApmMetricsStat = "hits_per_second"
	APMMETRICSSTAT_APDEX                ApmMetricsStat = "apdex"
	APMMETRICSSTAT_LATENCY_AVG          ApmMetricsStat = "latency_avg"
	APMMETRICSSTAT_LATENCY_MAX          ApmMetricsStat = "latency_max"
	APMMETRICSSTAT_LATENCY_P50          ApmMetricsStat = "latency_p50"
	APMMETRICSSTAT_LATENCY_P75          ApmMetricsStat = "latency_p75"
	APMMETRICSSTAT_LATENCY_P90          ApmMetricsStat = "latency_p90"
	APMMETRICSSTAT_LATENCY_P95          ApmMetricsStat = "latency_p95"
	APMMETRICSSTAT_LATENCY_P99          ApmMetricsStat = "latency_p99"
	APMMETRICSSTAT_LATENCY_P999         ApmMetricsStat = "latency_p999"
	APMMETRICSSTAT_LATENCY_DISTRIBUTION ApmMetricsStat = "latency_distribution"
	APMMETRICSSTAT_TOTAL_TIME           ApmMetricsStat = "total_time"
)

var allowedApmMetricsStatEnumValues = []ApmMetricsStat{
	APMMETRICSSTAT_ERROR_RATE,
	APMMETRICSSTAT_ERRORS,
	APMMETRICSSTAT_ERRORS_PER_SECOND,
	APMMETRICSSTAT_HITS,
	APMMETRICSSTAT_HITS_PER_SECOND,
	APMMETRICSSTAT_APDEX,
	APMMETRICSSTAT_LATENCY_AVG,
	APMMETRICSSTAT_LATENCY_MAX,
	APMMETRICSSTAT_LATENCY_P50,
	APMMETRICSSTAT_LATENCY_P75,
	APMMETRICSSTAT_LATENCY_P90,
	APMMETRICSSTAT_LATENCY_P95,
	APMMETRICSSTAT_LATENCY_P99,
	APMMETRICSSTAT_LATENCY_P999,
	APMMETRICSSTAT_LATENCY_DISTRIBUTION,
	APMMETRICSSTAT_TOTAL_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricsStat) GetAllowedValues() []ApmMetricsStat {
	return allowedApmMetricsStatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricsStat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricsStat(value)
	return nil
}

// NewApmMetricsStatFromValue returns a pointer to a valid ApmMetricsStat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricsStatFromValue(v string) (*ApmMetricsStat, error) {
	ev := ApmMetricsStat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricsStat: valid values are %v", v, allowedApmMetricsStatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricsStat) IsValid() bool {
	for _, existing := range allowedApmMetricsStatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricsStat value.
func (v ApmMetricsStat) Ptr() *ApmMetricsStat {
	return &v
}
