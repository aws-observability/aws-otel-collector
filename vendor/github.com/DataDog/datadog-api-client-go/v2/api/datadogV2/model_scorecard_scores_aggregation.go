// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardScoresAggregation Dimension to group scores by.
type ScorecardScoresAggregation string

// List of ScorecardScoresAggregation.
const (
	SCORECARDSCORESAGGREGATION_BY_ENTITY    ScorecardScoresAggregation = "by-entity"
	SCORECARDSCORESAGGREGATION_BY_RULE      ScorecardScoresAggregation = "by-rule"
	SCORECARDSCORESAGGREGATION_BY_SCORECARD ScorecardScoresAggregation = "by-scorecard"
	SCORECARDSCORESAGGREGATION_BY_TEAM      ScorecardScoresAggregation = "by-team"
	SCORECARDSCORESAGGREGATION_BY_KIND      ScorecardScoresAggregation = "by-kind"
)

var allowedScorecardScoresAggregationEnumValues = []ScorecardScoresAggregation{
	SCORECARDSCORESAGGREGATION_BY_ENTITY,
	SCORECARDSCORESAGGREGATION_BY_RULE,
	SCORECARDSCORESAGGREGATION_BY_SCORECARD,
	SCORECARDSCORESAGGREGATION_BY_TEAM,
	SCORECARDSCORESAGGREGATION_BY_KIND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScorecardScoresAggregation) GetAllowedValues() []ScorecardScoresAggregation {
	return allowedScorecardScoresAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScorecardScoresAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScorecardScoresAggregation(value)
	return nil
}

// NewScorecardScoresAggregationFromValue returns a pointer to a valid ScorecardScoresAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScorecardScoresAggregationFromValue(v string) (*ScorecardScoresAggregation, error) {
	ev := ScorecardScoresAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScorecardScoresAggregation: valid values are %v", v, allowedScorecardScoresAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScorecardScoresAggregation) IsValid() bool {
	for _, existing := range allowedScorecardScoresAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScorecardScoresAggregation value.
func (v ScorecardScoresAggregation) Ptr() *ScorecardScoresAggregation {
	return &v
}
