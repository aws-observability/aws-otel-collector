// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActiveBillingDimensionsType Type of active billing dimensions data.
type ActiveBillingDimensionsType string

// List of ActiveBillingDimensionsType.
const (
	ACTIVEBILLINGDIMENSIONSTYPE_BILLING_DIMENSIONS ActiveBillingDimensionsType = "billing_dimensions"
)

var allowedActiveBillingDimensionsTypeEnumValues = []ActiveBillingDimensionsType{
	ACTIVEBILLINGDIMENSIONSTYPE_BILLING_DIMENSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ActiveBillingDimensionsType) GetAllowedValues() []ActiveBillingDimensionsType {
	return allowedActiveBillingDimensionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ActiveBillingDimensionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ActiveBillingDimensionsType(value)
	return nil
}

// NewActiveBillingDimensionsTypeFromValue returns a pointer to a valid ActiveBillingDimensionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewActiveBillingDimensionsTypeFromValue(v string) (*ActiveBillingDimensionsType, error) {
	ev := ActiveBillingDimensionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ActiveBillingDimensionsType: valid values are %v", v, allowedActiveBillingDimensionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ActiveBillingDimensionsType) IsValid() bool {
	for _, existing := range allowedActiveBillingDimensionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActiveBillingDimensionsType value.
func (v ActiveBillingDimensionsType) Ptr() *ActiveBillingDimensionsType {
	return &v
}
