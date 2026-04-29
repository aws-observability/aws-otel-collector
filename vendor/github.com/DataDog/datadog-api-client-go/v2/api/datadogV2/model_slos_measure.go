// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlosMeasure The SLO measurement to retrieve.
type SlosMeasure string

// List of SlosMeasure.
const (
	SLOSMEASURE_GOOD_EVENTS                    SlosMeasure = "good_events"
	SLOSMEASURE_BAD_EVENTS                     SlosMeasure = "bad_events"
	SLOSMEASURE_SLO_STATUS                     SlosMeasure = "slo_status"
	SLOSMEASURE_ERROR_BUDGET_REMAINING         SlosMeasure = "error_budget_remaining"
	SLOSMEASURE_ERROR_BUDGET_REMAINING_HISTORY SlosMeasure = "error_budget_remaining_history"
	SLOSMEASURE_ERROR_BUDGET_BURNDOWN          SlosMeasure = "error_budget_burndown"
	SLOSMEASURE_BURN_RATE                      SlosMeasure = "burn_rate"
	SLOSMEASURE_SLO_STATUS_HISTORY             SlosMeasure = "slo_status_history"
	SLOSMEASURE_GOOD_MINUTES                   SlosMeasure = "good_minutes"
	SLOSMEASURE_BAD_MINUTES                    SlosMeasure = "bad_minutes"
)

var allowedSlosMeasureEnumValues = []SlosMeasure{
	SLOSMEASURE_GOOD_EVENTS,
	SLOSMEASURE_BAD_EVENTS,
	SLOSMEASURE_SLO_STATUS,
	SLOSMEASURE_ERROR_BUDGET_REMAINING,
	SLOSMEASURE_ERROR_BUDGET_REMAINING_HISTORY,
	SLOSMEASURE_ERROR_BUDGET_BURNDOWN,
	SLOSMEASURE_BURN_RATE,
	SLOSMEASURE_SLO_STATUS_HISTORY,
	SLOSMEASURE_GOOD_MINUTES,
	SLOSMEASURE_BAD_MINUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SlosMeasure) GetAllowedValues() []SlosMeasure {
	return allowedSlosMeasureEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SlosMeasure) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SlosMeasure(value)
	return nil
}

// NewSlosMeasureFromValue returns a pointer to a valid SlosMeasure
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSlosMeasureFromValue(v string) (*SlosMeasure, error) {
	ev := SlosMeasure(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SlosMeasure: valid values are %v", v, allowedSlosMeasureEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SlosMeasure) IsValid() bool {
	for _, existing := range allowedSlosMeasureEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SlosMeasure value.
func (v SlosMeasure) Ptr() *SlosMeasure {
	return &v
}
