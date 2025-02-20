// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricComputeAggregationType The type of aggregation to use.
type RumMetricComputeAggregationType string

// List of RumMetricComputeAggregationType.
const (
	RUMMETRICCOMPUTEAGGREGATIONTYPE_COUNT        RumMetricComputeAggregationType = "count"
	RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION RumMetricComputeAggregationType = "distribution"
)

var allowedRumMetricComputeAggregationTypeEnumValues = []RumMetricComputeAggregationType{
	RUMMETRICCOMPUTEAGGREGATIONTYPE_COUNT,
	RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumMetricComputeAggregationType) GetAllowedValues() []RumMetricComputeAggregationType {
	return allowedRumMetricComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumMetricComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumMetricComputeAggregationType(value)
	return nil
}

// NewRumMetricComputeAggregationTypeFromValue returns a pointer to a valid RumMetricComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumMetricComputeAggregationTypeFromValue(v string) (*RumMetricComputeAggregationType, error) {
	ev := RumMetricComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumMetricComputeAggregationType: valid values are %v", v, allowedRumMetricComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumMetricComputeAggregationType) IsValid() bool {
	for _, existing := range allowedRumMetricComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumMetricComputeAggregationType value.
func (v RumMetricComputeAggregationType) Ptr() *RumMetricComputeAggregationType {
	return &v
}
