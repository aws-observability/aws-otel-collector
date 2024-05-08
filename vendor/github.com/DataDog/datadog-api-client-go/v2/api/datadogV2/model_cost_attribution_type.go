// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAttributionType Type of cost attribution data.
type CostAttributionType string

// List of CostAttributionType.
const (
	COSTATTRIBUTIONTYPE_COST_BY_TAG CostAttributionType = "cost_by_tag"
)

var allowedCostAttributionTypeEnumValues = []CostAttributionType{
	COSTATTRIBUTIONTYPE_COST_BY_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostAttributionType) GetAllowedValues() []CostAttributionType {
	return allowedCostAttributionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostAttributionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostAttributionType(value)
	return nil
}

// NewCostAttributionTypeFromValue returns a pointer to a valid CostAttributionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostAttributionTypeFromValue(v string) (*CostAttributionType, error) {
	ev := CostAttributionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostAttributionType: valid values are %v", v, allowedCostAttributionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostAttributionType) IsValid() bool {
	for _, existing := range allowedCostAttributionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostAttributionType value.
func (v CostAttributionType) Ptr() *CostAttributionType {
	return &v
}
