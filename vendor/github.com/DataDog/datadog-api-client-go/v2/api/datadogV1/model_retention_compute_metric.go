// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionComputeMetric Metric for retention compute.
type RetentionComputeMetric string

// List of RetentionComputeMetric.
const (
	RETENTIONCOMPUTEMETRIC_RETENTION      RetentionComputeMetric = "__dd.retention"
	RETENTIONCOMPUTEMETRIC_RETENTION_RATE RetentionComputeMetric = "__dd.retention_rate"
)

var allowedRetentionComputeMetricEnumValues = []RetentionComputeMetric{
	RETENTIONCOMPUTEMETRIC_RETENTION,
	RETENTIONCOMPUTEMETRIC_RETENTION_RATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionComputeMetric) GetAllowedValues() []RetentionComputeMetric {
	return allowedRetentionComputeMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionComputeMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionComputeMetric(value)
	return nil
}

// NewRetentionComputeMetricFromValue returns a pointer to a valid RetentionComputeMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionComputeMetricFromValue(v string) (*RetentionComputeMetric, error) {
	ev := RetentionComputeMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionComputeMetric: valid values are %v", v, allowedRetentionComputeMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionComputeMetric) IsValid() bool {
	for _, existing := range allowedRetentionComputeMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionComputeMetric value.
func (v RetentionComputeMetric) Ptr() *RetentionComputeMetric {
	return &v
}
