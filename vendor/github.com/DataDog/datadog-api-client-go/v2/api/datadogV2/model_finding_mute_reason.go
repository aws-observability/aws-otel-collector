// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingMuteReason The reason why this finding is muted or unmuted.
type FindingMuteReason string

// List of FindingMuteReason.
const (
	FINDINGMUTEREASON_PENDING_FIX             FindingMuteReason = "PENDING_FIX"
	FINDINGMUTEREASON_FALSE_POSITIVE          FindingMuteReason = "FALSE_POSITIVE"
	FINDINGMUTEREASON_ACCEPTED_RISK           FindingMuteReason = "ACCEPTED_RISK"
	FINDINGMUTEREASON_NO_PENDING_FIX          FindingMuteReason = "NO_PENDING_FIX"
	FINDINGMUTEREASON_HUMAN_ERROR             FindingMuteReason = "HUMAN_ERROR"
	FINDINGMUTEREASON_NO_LONGER_ACCEPTED_RISK FindingMuteReason = "NO_LONGER_ACCEPTED_RISK"
	FINDINGMUTEREASON_OTHER                   FindingMuteReason = "OTHER"
)

var allowedFindingMuteReasonEnumValues = []FindingMuteReason{
	FINDINGMUTEREASON_PENDING_FIX,
	FINDINGMUTEREASON_FALSE_POSITIVE,
	FINDINGMUTEREASON_ACCEPTED_RISK,
	FINDINGMUTEREASON_NO_PENDING_FIX,
	FINDINGMUTEREASON_HUMAN_ERROR,
	FINDINGMUTEREASON_NO_LONGER_ACCEPTED_RISK,
	FINDINGMUTEREASON_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingMuteReason) GetAllowedValues() []FindingMuteReason {
	return allowedFindingMuteReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingMuteReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingMuteReason(value)
	return nil
}

// NewFindingMuteReasonFromValue returns a pointer to a valid FindingMuteReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingMuteReasonFromValue(v string) (*FindingMuteReason, error) {
	ev := FindingMuteReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingMuteReason: valid values are %v", v, allowedFindingMuteReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingMuteReason) IsValid() bool {
	for _, existing := range allowedFindingMuteReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingMuteReason value.
func (v FindingMuteReason) Ptr() *FindingMuteReason {
	return &v
}
