// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostMetricType Type of the Cloud Cost Management available metric resource.
type CostMetricType string

// List of CostMetricType.
const (
	COSTMETRICTYPE_COST_METRIC CostMetricType = "cost_metric"
)

var allowedCostMetricTypeEnumValues = []CostMetricType{
	COSTMETRICTYPE_COST_METRIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostMetricType) GetAllowedValues() []CostMetricType {
	return allowedCostMetricTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostMetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostMetricType(value)
	return nil
}

// NewCostMetricTypeFromValue returns a pointer to a valid CostMetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostMetricTypeFromValue(v string) (*CostMetricType, error) {
	ev := CostMetricType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostMetricType: valid values are %v", v, allowedCostMetricTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostMetricType) IsValid() bool {
	for _, existing := range allowedCostMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostMetricType value.
func (v CostMetricType) Ptr() *CostMetricType {
	return &v
}
