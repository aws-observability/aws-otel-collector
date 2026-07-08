// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedWaterfallPerformanceCriteriaMetric Performance metric used to filter view instances by threshold.
type AggregatedWaterfallPerformanceCriteriaMetric string

// List of AggregatedWaterfallPerformanceCriteriaMetric.
const (
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_LOADING_TIME              AggregatedWaterfallPerformanceCriteriaMetric = "loading_time"
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_LARGEST_CONTENTFUL_PAINT  AggregatedWaterfallPerformanceCriteriaMetric = "largest_contentful_paint"
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_FIRST_CONTENTFUL_PAINT    AggregatedWaterfallPerformanceCriteriaMetric = "first_contentful_paint"
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_INTERACTION_TO_NEXT_PAINT AggregatedWaterfallPerformanceCriteriaMetric = "interaction_to_next_paint"
)

var allowedAggregatedWaterfallPerformanceCriteriaMetricEnumValues = []AggregatedWaterfallPerformanceCriteriaMetric{
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_LOADING_TIME,
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_LARGEST_CONTENTFUL_PAINT,
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_FIRST_CONTENTFUL_PAINT,
	AGGREGATEDWATERFALLPERFORMANCECRITERIAMETRIC_INTERACTION_TO_NEXT_PAINT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregatedWaterfallPerformanceCriteriaMetric) GetAllowedValues() []AggregatedWaterfallPerformanceCriteriaMetric {
	return allowedAggregatedWaterfallPerformanceCriteriaMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregatedWaterfallPerformanceCriteriaMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregatedWaterfallPerformanceCriteriaMetric(value)
	return nil
}

// NewAggregatedWaterfallPerformanceCriteriaMetricFromValue returns a pointer to a valid AggregatedWaterfallPerformanceCriteriaMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregatedWaterfallPerformanceCriteriaMetricFromValue(v string) (*AggregatedWaterfallPerformanceCriteriaMetric, error) {
	ev := AggregatedWaterfallPerformanceCriteriaMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregatedWaterfallPerformanceCriteriaMetric: valid values are %v", v, allowedAggregatedWaterfallPerformanceCriteriaMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregatedWaterfallPerformanceCriteriaMetric) IsValid() bool {
	for _, existing := range allowedAggregatedWaterfallPerformanceCriteriaMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregatedWaterfallPerformanceCriteriaMetric value.
func (v AggregatedWaterfallPerformanceCriteriaMetric) Ptr() *AggregatedWaterfallPerformanceCriteriaMetric {
	return &v
}
