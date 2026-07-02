// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagType Type of the Cloud Cost Management tag resource.
type CostTagType string

// List of CostTagType.
const (
	COSTTAGTYPE_COST_TAG CostTagType = "cost_tag"
)

var allowedCostTagTypeEnumValues = []CostTagType{
	COSTTAGTYPE_COST_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagType) GetAllowedValues() []CostTagType {
	return allowedCostTagTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagType(value)
	return nil
}

// NewCostTagTypeFromValue returns a pointer to a valid CostTagType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagTypeFromValue(v string) (*CostTagType, error) {
	ev := CostTagType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagType: valid values are %v", v, allowedCostTagTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagType) IsValid() bool {
	for _, existing := range allowedCostTagTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagType value.
func (v CostTagType) Ptr() *CostTagType {
	return &v
}
