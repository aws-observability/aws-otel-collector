// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardScoreDataType The JSON:API resource type.
type ScorecardScoreDataType string

// List of ScorecardScoreDataType.
const (
	SCORECARDSCOREDATATYPE_SCORE ScorecardScoreDataType = "score"
)

var allowedScorecardScoreDataTypeEnumValues = []ScorecardScoreDataType{
	SCORECARDSCOREDATATYPE_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScorecardScoreDataType) GetAllowedValues() []ScorecardScoreDataType {
	return allowedScorecardScoreDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScorecardScoreDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScorecardScoreDataType(value)
	return nil
}

// NewScorecardScoreDataTypeFromValue returns a pointer to a valid ScorecardScoreDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScorecardScoreDataTypeFromValue(v string) (*ScorecardScoreDataType, error) {
	ev := ScorecardScoreDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScorecardScoreDataType: valid values are %v", v, allowedScorecardScoreDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScorecardScoreDataType) IsValid() bool {
	for _, existing := range allowedScorecardScoreDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScorecardScoreDataType value.
func (v ScorecardScoreDataType) Ptr() *ScorecardScoreDataType {
	return &v
}
