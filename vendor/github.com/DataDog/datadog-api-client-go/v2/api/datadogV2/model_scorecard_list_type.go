// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardListType The JSON:API type for scorecard list.
type ScorecardListType string

// List of ScorecardListType.
const (
	SCORECARDLISTTYPE_SCORECARD ScorecardListType = "scorecard"
)

var allowedScorecardListTypeEnumValues = []ScorecardListType{
	SCORECARDLISTTYPE_SCORECARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScorecardListType) GetAllowedValues() []ScorecardListType {
	return allowedScorecardListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScorecardListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScorecardListType(value)
	return nil
}

// NewScorecardListTypeFromValue returns a pointer to a valid ScorecardListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScorecardListTypeFromValue(v string) (*ScorecardListType, error) {
	ev := ScorecardListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScorecardListType: valid values are %v", v, allowedScorecardListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScorecardListType) IsValid() bool {
	for _, existing := range allowedScorecardListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScorecardListType value.
func (v ScorecardListType) Ptr() *ScorecardListType {
	return &v
}
