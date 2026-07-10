// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationDataType Recommendation resource type.
type CostRecommendationDataType string

// List of CostRecommendationDataType.
const (
	COSTRECOMMENDATIONDATATYPE_RECOMMENDATION CostRecommendationDataType = "recommendation"
)

var allowedCostRecommendationDataTypeEnumValues = []CostRecommendationDataType{
	COSTRECOMMENDATIONDATATYPE_RECOMMENDATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostRecommendationDataType) GetAllowedValues() []CostRecommendationDataType {
	return allowedCostRecommendationDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostRecommendationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostRecommendationDataType(value)
	return nil
}

// NewCostRecommendationDataTypeFromValue returns a pointer to a valid CostRecommendationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostRecommendationDataTypeFromValue(v string) (*CostRecommendationDataType, error) {
	ev := CostRecommendationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostRecommendationDataType: valid values are %v", v, allowedCostRecommendationDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostRecommendationDataType) IsValid() bool {
	for _, existing := range allowedCostRecommendationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostRecommendationDataType value.
func (v CostRecommendationDataType) Ptr() *CostRecommendationDataType {
	return &v
}
