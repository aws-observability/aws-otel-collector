// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagKeyType Type of the Cloud Cost Management tag key resource.
type CostTagKeyType string

// List of CostTagKeyType.
const (
	COSTTAGKEYTYPE_COST_TAG_KEY CostTagKeyType = "cost_tag_key"
)

var allowedCostTagKeyTypeEnumValues = []CostTagKeyType{
	COSTTAGKEYTYPE_COST_TAG_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagKeyType) GetAllowedValues() []CostTagKeyType {
	return allowedCostTagKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagKeyType(value)
	return nil
}

// NewCostTagKeyTypeFromValue returns a pointer to a valid CostTagKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagKeyTypeFromValue(v string) (*CostTagKeyType, error) {
	ev := CostTagKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagKeyType: valid values are %v", v, allowedCostTagKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagKeyType) IsValid() bool {
	for _, existing := range allowedCostTagKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagKeyType value.
func (v CostTagKeyType) Ptr() *CostTagKeyType {
	return &v
}
