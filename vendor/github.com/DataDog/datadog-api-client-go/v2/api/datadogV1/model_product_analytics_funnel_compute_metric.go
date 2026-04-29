// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelComputeMetric Metric for user journey funnel compute. `__dd.conversion` and `__dd.conversion_rate` accept `count` (unique users/sessions) and `cardinality` (total users/sessions) as aggregations.
type ProductAnalyticsFunnelComputeMetric string

// List of ProductAnalyticsFunnelComputeMetric.
const (
	PRODUCTANALYTICSFUNNELCOMPUTEMETRIC_CONVERSION      ProductAnalyticsFunnelComputeMetric = "__dd.conversion"
	PRODUCTANALYTICSFUNNELCOMPUTEMETRIC_CONVERSION_RATE ProductAnalyticsFunnelComputeMetric = "__dd.conversion_rate"
)

var allowedProductAnalyticsFunnelComputeMetricEnumValues = []ProductAnalyticsFunnelComputeMetric{
	PRODUCTANALYTICSFUNNELCOMPUTEMETRIC_CONVERSION,
	PRODUCTANALYTICSFUNNELCOMPUTEMETRIC_CONVERSION_RATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelComputeMetric) GetAllowedValues() []ProductAnalyticsFunnelComputeMetric {
	return allowedProductAnalyticsFunnelComputeMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelComputeMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelComputeMetric(value)
	return nil
}

// NewProductAnalyticsFunnelComputeMetricFromValue returns a pointer to a valid ProductAnalyticsFunnelComputeMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelComputeMetricFromValue(v string) (*ProductAnalyticsFunnelComputeMetric, error) {
	ev := ProductAnalyticsFunnelComputeMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelComputeMetric: valid values are %v", v, allowedProductAnalyticsFunnelComputeMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelComputeMetric) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelComputeMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelComputeMetric value.
func (v ProductAnalyticsFunnelComputeMetric) Ptr() *ProductAnalyticsFunnelComputeMetric {
	return &v
}
