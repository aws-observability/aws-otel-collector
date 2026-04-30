// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAggregationType Controls how costs are aggregated when using `start_date`. The `cumulative` option returns month-to-date running totals.
type CostAggregationType string

// List of CostAggregationType.
const (
	COSTAGGREGATIONTYPE_CUMULATIVE CostAggregationType = "cumulative"
)

var allowedCostAggregationTypeEnumValues = []CostAggregationType{
	COSTAGGREGATIONTYPE_CUMULATIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostAggregationType) GetAllowedValues() []CostAggregationType {
	return allowedCostAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostAggregationType(value)
	return nil
}

// NewCostAggregationTypeFromValue returns a pointer to a valid CostAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostAggregationTypeFromValue(v string) (*CostAggregationType, error) {
	ev := CostAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostAggregationType: valid values are %v", v, allowedCostAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostAggregationType) IsValid() bool {
	for _, existing := range allowedCostAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostAggregationType value.
func (v CostAggregationType) Ptr() *CostAggregationType {
	return &v
}
