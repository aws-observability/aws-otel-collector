// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeySourceType Type of the Cloud Cost Management tag source resource.
type CostTagKeySourceType string

// List of CostTagKeySourceType.
const (
	COSTTAGKEYSOURCETYPE_COST_TAG_KEY_SOURCE CostTagKeySourceType = "cost_tag_key_source"
)

var allowedCostTagKeySourceTypeEnumValues = []CostTagKeySourceType{
	COSTTAGKEYSOURCETYPE_COST_TAG_KEY_SOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagKeySourceType) GetAllowedValues() []CostTagKeySourceType {
	return allowedCostTagKeySourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagKeySourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagKeySourceType(value)
	return nil
}

// NewCostTagKeySourceTypeFromValue returns a pointer to a valid CostTagKeySourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagKeySourceTypeFromValue(v string) (*CostTagKeySourceType, error) {
	ev := CostTagKeySourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagKeySourceType: valid values are %v", v, allowedCostTagKeySourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagKeySourceType) IsValid() bool {
	for _, existing := range allowedCostTagKeySourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagKeySourceType value.
func (v CostTagKeySourceType) Ptr() *CostTagKeySourceType {
	return &v
}
