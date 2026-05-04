// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmResourceStatName The APM resource statistic to query.
type ApmResourceStatName string

// List of ApmResourceStatName.
const (
	APMRESOURCESTATNAME_ERROR_RATE           ApmResourceStatName = "error_rate"
	APMRESOURCESTATNAME_ERRORS               ApmResourceStatName = "errors"
	APMRESOURCESTATNAME_HITS                 ApmResourceStatName = "hits"
	APMRESOURCESTATNAME_LATENCY_AVG          ApmResourceStatName = "latency_avg"
	APMRESOURCESTATNAME_LATENCY_MAX          ApmResourceStatName = "latency_max"
	APMRESOURCESTATNAME_LATENCY_P50          ApmResourceStatName = "latency_p50"
	APMRESOURCESTATNAME_LATENCY_P75          ApmResourceStatName = "latency_p75"
	APMRESOURCESTATNAME_LATENCY_P90          ApmResourceStatName = "latency_p90"
	APMRESOURCESTATNAME_LATENCY_P95          ApmResourceStatName = "latency_p95"
	APMRESOURCESTATNAME_LATENCY_P99          ApmResourceStatName = "latency_p99"
	APMRESOURCESTATNAME_LATENCY_DISTRIBUTION ApmResourceStatName = "latency_distribution"
	APMRESOURCESTATNAME_TOTAL_TIME           ApmResourceStatName = "total_time"
)

var allowedApmResourceStatNameEnumValues = []ApmResourceStatName{
	APMRESOURCESTATNAME_ERROR_RATE,
	APMRESOURCESTATNAME_ERRORS,
	APMRESOURCESTATNAME_HITS,
	APMRESOURCESTATNAME_LATENCY_AVG,
	APMRESOURCESTATNAME_LATENCY_MAX,
	APMRESOURCESTATNAME_LATENCY_P50,
	APMRESOURCESTATNAME_LATENCY_P75,
	APMRESOURCESTATNAME_LATENCY_P90,
	APMRESOURCESTATNAME_LATENCY_P95,
	APMRESOURCESTATNAME_LATENCY_P99,
	APMRESOURCESTATNAME_LATENCY_DISTRIBUTION,
	APMRESOURCESTATNAME_TOTAL_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmResourceStatName) GetAllowedValues() []ApmResourceStatName {
	return allowedApmResourceStatNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmResourceStatName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmResourceStatName(value)
	return nil
}

// NewApmResourceStatNameFromValue returns a pointer to a valid ApmResourceStatName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmResourceStatNameFromValue(v string) (*ApmResourceStatName, error) {
	ev := ApmResourceStatName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmResourceStatName: valid values are %v", v, allowedApmResourceStatNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmResourceStatName) IsValid() bool {
	for _, existing := range allowedApmResourceStatNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmResourceStatName value.
func (v ApmResourceStatName) Ptr() *ApmResourceStatName {
	return &v
}
