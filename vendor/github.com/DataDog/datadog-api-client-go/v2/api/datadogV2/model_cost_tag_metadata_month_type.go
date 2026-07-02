// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagMetadataMonthType Type of the Cloud Cost Management tag metadata month resource.
type CostTagMetadataMonthType string

// List of CostTagMetadataMonthType.
const (
	COSTTAGMETADATAMONTHTYPE_COST_TAG_METADATA_MONTH CostTagMetadataMonthType = "cost_tag_metadata_month"
)

var allowedCostTagMetadataMonthTypeEnumValues = []CostTagMetadataMonthType{
	COSTTAGMETADATAMONTHTYPE_COST_TAG_METADATA_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagMetadataMonthType) GetAllowedValues() []CostTagMetadataMonthType {
	return allowedCostTagMetadataMonthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagMetadataMonthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagMetadataMonthType(value)
	return nil
}

// NewCostTagMetadataMonthTypeFromValue returns a pointer to a valid CostTagMetadataMonthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagMetadataMonthTypeFromValue(v string) (*CostTagMetadataMonthType, error) {
	ev := CostTagMetadataMonthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagMetadataMonthType: valid values are %v", v, allowedCostTagMetadataMonthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagMetadataMonthType) IsValid() bool {
	for _, existing := range allowedCostTagMetadataMonthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagMetadataMonthType value.
func (v CostTagMetadataMonthType) Ptr() *CostTagMetadataMonthType {
	return &v
}
