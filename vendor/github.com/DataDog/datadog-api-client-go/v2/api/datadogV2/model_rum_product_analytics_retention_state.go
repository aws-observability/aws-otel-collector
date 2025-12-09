// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMProductAnalyticsRetentionState Controls the retention policy for Product Analytics data derived from RUM events.
type RUMProductAnalyticsRetentionState string

// List of RUMProductAnalyticsRetentionState.
const (
	RUMPRODUCTANALYTICSRETENTIONSTATE_MAX  RUMProductAnalyticsRetentionState = "MAX"
	RUMPRODUCTANALYTICSRETENTIONSTATE_NONE RUMProductAnalyticsRetentionState = "NONE"
)

var allowedRUMProductAnalyticsRetentionStateEnumValues = []RUMProductAnalyticsRetentionState{
	RUMPRODUCTANALYTICSRETENTIONSTATE_MAX,
	RUMPRODUCTANALYTICSRETENTIONSTATE_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMProductAnalyticsRetentionState) GetAllowedValues() []RUMProductAnalyticsRetentionState {
	return allowedRUMProductAnalyticsRetentionStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMProductAnalyticsRetentionState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMProductAnalyticsRetentionState(value)
	return nil
}

// NewRUMProductAnalyticsRetentionStateFromValue returns a pointer to a valid RUMProductAnalyticsRetentionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMProductAnalyticsRetentionStateFromValue(v string) (*RUMProductAnalyticsRetentionState, error) {
	ev := RUMProductAnalyticsRetentionState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMProductAnalyticsRetentionState: valid values are %v", v, allowedRUMProductAnalyticsRetentionStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMProductAnalyticsRetentionState) IsValid() bool {
	for _, existing := range allowedRUMProductAnalyticsRetentionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMProductAnalyticsRetentionState value.
func (v RUMProductAnalyticsRetentionState) Ptr() *RUMProductAnalyticsRetentionState {
	return &v
}
