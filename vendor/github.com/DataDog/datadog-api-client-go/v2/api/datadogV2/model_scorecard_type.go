// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScorecardType The JSON:API type for scorecard.
type ScorecardType string

// List of ScorecardType.
const (
	SCORECARDTYPE_SCORECARD ScorecardType = "scorecard"
)

var allowedScorecardTypeEnumValues = []ScorecardType{
	SCORECARDTYPE_SCORECARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScorecardType) GetAllowedValues() []ScorecardType {
	return allowedScorecardTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScorecardType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScorecardType(value)
	return nil
}

// NewScorecardTypeFromValue returns a pointer to a valid ScorecardType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScorecardTypeFromValue(v string) (*ScorecardType, error) {
	ev := ScorecardType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScorecardType: valid values are %v", v, allowedScorecardTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScorecardType) IsValid() bool {
	for _, existing := range allowedScorecardTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScorecardType value.
func (v ScorecardType) Ptr() *ScorecardType {
	return &v
}
