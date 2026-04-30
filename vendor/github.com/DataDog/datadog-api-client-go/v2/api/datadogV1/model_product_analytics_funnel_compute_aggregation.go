// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelComputeAggregation Aggregation type for user journey funnel compute.
type ProductAnalyticsFunnelComputeAggregation string

// List of ProductAnalyticsFunnelComputeAggregation.
const (
	PRODUCTANALYTICSFUNNELCOMPUTEAGGREGATION_CARDINALITY ProductAnalyticsFunnelComputeAggregation = "cardinality"
	PRODUCTANALYTICSFUNNELCOMPUTEAGGREGATION_COUNT       ProductAnalyticsFunnelComputeAggregation = "count"
)

var allowedProductAnalyticsFunnelComputeAggregationEnumValues = []ProductAnalyticsFunnelComputeAggregation{
	PRODUCTANALYTICSFUNNELCOMPUTEAGGREGATION_CARDINALITY,
	PRODUCTANALYTICSFUNNELCOMPUTEAGGREGATION_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelComputeAggregation) GetAllowedValues() []ProductAnalyticsFunnelComputeAggregation {
	return allowedProductAnalyticsFunnelComputeAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelComputeAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelComputeAggregation(value)
	return nil
}

// NewProductAnalyticsFunnelComputeAggregationFromValue returns a pointer to a valid ProductAnalyticsFunnelComputeAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelComputeAggregationFromValue(v string) (*ProductAnalyticsFunnelComputeAggregation, error) {
	ev := ProductAnalyticsFunnelComputeAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelComputeAggregation: valid values are %v", v, allowedProductAnalyticsFunnelComputeAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelComputeAggregation) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelComputeAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelComputeAggregation value.
func (v ProductAnalyticsFunnelComputeAggregation) Ptr() *ProductAnalyticsFunnelComputeAggregation {
	return &v
}
