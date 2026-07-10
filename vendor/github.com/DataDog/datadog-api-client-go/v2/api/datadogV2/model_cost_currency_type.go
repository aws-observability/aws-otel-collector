// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostCurrencyType Type of the Cloud Cost Management billing currency resource.
type CostCurrencyType string

// List of CostCurrencyType.
const (
	COSTCURRENCYTYPE_COST_CURRENCY CostCurrencyType = "cost_currency"
)

var allowedCostCurrencyTypeEnumValues = []CostCurrencyType{
	COSTCURRENCYTYPE_COST_CURRENCY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostCurrencyType) GetAllowedValues() []CostCurrencyType {
	return allowedCostCurrencyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostCurrencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostCurrencyType(value)
	return nil
}

// NewCostCurrencyTypeFromValue returns a pointer to a valid CostCurrencyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostCurrencyTypeFromValue(v string) (*CostCurrencyType, error) {
	ev := CostCurrencyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostCurrencyType: valid values are %v", v, allowedCostCurrencyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostCurrencyType) IsValid() bool {
	for _, existing := range allowedCostCurrencyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostCurrencyType value.
func (v CostCurrencyType) Ptr() *CostCurrencyType {
	return &v
}
