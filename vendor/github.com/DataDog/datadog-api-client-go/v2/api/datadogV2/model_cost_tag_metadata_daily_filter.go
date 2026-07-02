// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagMetadataDailyFilter Granularity for tag metadata results. `true` returns one row per day, `false` (or omitted) returns the monthly roll-up.
type CostTagMetadataDailyFilter string

// List of CostTagMetadataDailyFilter.
const (
	COSTTAGMETADATADAILYFILTER_TRUE  CostTagMetadataDailyFilter = "true"
	COSTTAGMETADATADAILYFILTER_FALSE CostTagMetadataDailyFilter = "false"
)

var allowedCostTagMetadataDailyFilterEnumValues = []CostTagMetadataDailyFilter{
	COSTTAGMETADATADAILYFILTER_TRUE,
	COSTTAGMETADATADAILYFILTER_FALSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagMetadataDailyFilter) GetAllowedValues() []CostTagMetadataDailyFilter {
	return allowedCostTagMetadataDailyFilterEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagMetadataDailyFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagMetadataDailyFilter(value)
	return nil
}

// NewCostTagMetadataDailyFilterFromValue returns a pointer to a valid CostTagMetadataDailyFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagMetadataDailyFilterFromValue(v string) (*CostTagMetadataDailyFilter, error) {
	ev := CostTagMetadataDailyFilter(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagMetadataDailyFilter: valid values are %v", v, allowedCostTagMetadataDailyFilterEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagMetadataDailyFilter) IsValid() bool {
	for _, existing := range allowedCostTagMetadataDailyFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagMetadataDailyFilter value.
func (v CostTagMetadataDailyFilter) Ptr() *CostTagMetadataDailyFilter {
	return &v
}
