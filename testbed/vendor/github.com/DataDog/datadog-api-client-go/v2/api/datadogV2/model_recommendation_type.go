// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationType JSON:API resource type for Spark Pod Autosizing recommendations. Identifies the Recommendation resource returned by SPA.
type RecommendationType string

// List of RecommendationType.
const (
	RECOMMENDATIONTYPE_RECOMMENDATION RecommendationType = "recommendation"
)

var allowedRecommendationTypeEnumValues = []RecommendationType{
	RECOMMENDATIONTYPE_RECOMMENDATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RecommendationType) GetAllowedValues() []RecommendationType {
	return allowedRecommendationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RecommendationType(value)
	return nil
}

// NewRecommendationTypeFromValue returns a pointer to a valid RecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRecommendationTypeFromValue(v string) (*RecommendationType, error) {
	ev := RecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RecommendationType: valid values are %v", v, allowedRecommendationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RecommendationType) IsValid() bool {
	for _, existing := range allowedRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationType value.
func (v RecommendationType) Ptr() *RecommendationType {
	return &v
}
