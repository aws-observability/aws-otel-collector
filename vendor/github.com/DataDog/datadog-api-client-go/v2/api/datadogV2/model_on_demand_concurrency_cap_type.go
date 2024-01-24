// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnDemandConcurrencyCapType On-demand concurrency cap type.
type OnDemandConcurrencyCapType string

// List of OnDemandConcurrencyCapType.
const (
	ONDEMANDCONCURRENCYCAPTYPE_ON_DEMAND_CONCURRENCY_CAP OnDemandConcurrencyCapType = "on_demand_concurrency_cap"
)

var allowedOnDemandConcurrencyCapTypeEnumValues = []OnDemandConcurrencyCapType{
	ONDEMANDCONCURRENCYCAPTYPE_ON_DEMAND_CONCURRENCY_CAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnDemandConcurrencyCapType) GetAllowedValues() []OnDemandConcurrencyCapType {
	return allowedOnDemandConcurrencyCapTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnDemandConcurrencyCapType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnDemandConcurrencyCapType(value)
	return nil
}

// NewOnDemandConcurrencyCapTypeFromValue returns a pointer to a valid OnDemandConcurrencyCapType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnDemandConcurrencyCapTypeFromValue(v string) (*OnDemandConcurrencyCapType, error) {
	ev := OnDemandConcurrencyCapType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnDemandConcurrencyCapType: valid values are %v", v, allowedOnDemandConcurrencyCapTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnDemandConcurrencyCapType) IsValid() bool {
	for _, existing := range allowedOnDemandConcurrencyCapTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnDemandConcurrencyCapType value.
func (v OnDemandConcurrencyCapType) Ptr() *OnDemandConcurrencyCapType {
	return &v
}
