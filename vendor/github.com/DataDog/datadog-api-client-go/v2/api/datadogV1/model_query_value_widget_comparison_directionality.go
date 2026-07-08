// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueWidgetComparisonDirectionality Color-coding direction: `increase_better` (green on rise), `decrease_better` (green on drop), or `neutral` (no color).
type QueryValueWidgetComparisonDirectionality string

// List of QueryValueWidgetComparisonDirectionality.
const (
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_INCREASE_BETTER QueryValueWidgetComparisonDirectionality = "increase_better"
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_DECREASE_BETTER QueryValueWidgetComparisonDirectionality = "decrease_better"
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_NEUTRAL         QueryValueWidgetComparisonDirectionality = "neutral"
)

var allowedQueryValueWidgetComparisonDirectionalityEnumValues = []QueryValueWidgetComparisonDirectionality{
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_INCREASE_BETTER,
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_DECREASE_BETTER,
	QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_NEUTRAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *QueryValueWidgetComparisonDirectionality) GetAllowedValues() []QueryValueWidgetComparisonDirectionality {
	return allowedQueryValueWidgetComparisonDirectionalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *QueryValueWidgetComparisonDirectionality) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = QueryValueWidgetComparisonDirectionality(value)
	return nil
}

// NewQueryValueWidgetComparisonDirectionalityFromValue returns a pointer to a valid QueryValueWidgetComparisonDirectionality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewQueryValueWidgetComparisonDirectionalityFromValue(v string) (*QueryValueWidgetComparisonDirectionality, error) {
	ev := QueryValueWidgetComparisonDirectionality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for QueryValueWidgetComparisonDirectionality: valid values are %v", v, allowedQueryValueWidgetComparisonDirectionalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v QueryValueWidgetComparisonDirectionality) IsValid() bool {
	for _, existing := range allowedQueryValueWidgetComparisonDirectionalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryValueWidgetComparisonDirectionality value.
func (v QueryValueWidgetComparisonDirectionality) Ptr() *QueryValueWidgetComparisonDirectionality {
	return &v
}
