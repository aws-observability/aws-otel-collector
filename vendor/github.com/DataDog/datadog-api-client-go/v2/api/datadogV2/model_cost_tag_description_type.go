// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagDescriptionType Type of the Cloud Cost Management tag description resource.
type CostTagDescriptionType string

// List of CostTagDescriptionType.
const (
	COSTTAGDESCRIPTIONTYPE_COST_TAG_DESCRIPTION CostTagDescriptionType = "cost_tag_description"
)

var allowedCostTagDescriptionTypeEnumValues = []CostTagDescriptionType{
	COSTTAGDESCRIPTIONTYPE_COST_TAG_DESCRIPTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagDescriptionType) GetAllowedValues() []CostTagDescriptionType {
	return allowedCostTagDescriptionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagDescriptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagDescriptionType(value)
	return nil
}

// NewCostTagDescriptionTypeFromValue returns a pointer to a valid CostTagDescriptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagDescriptionTypeFromValue(v string) (*CostTagDescriptionType, error) {
	ev := CostTagDescriptionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagDescriptionType: valid values are %v", v, allowedCostTagDescriptionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagDescriptionType) IsValid() bool {
	for _, existing := range allowedCostTagDescriptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagDescriptionType value.
func (v CostTagDescriptionType) Ptr() *CostTagDescriptionType {
	return &v
}
