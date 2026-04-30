// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RolloutStrategy The progression strategy used by a progressive rollout.
type RolloutStrategy string

// List of RolloutStrategy.
const (
	ROLLOUTSTRATEGY_UNIFORM_INTERVALS RolloutStrategy = "UNIFORM_INTERVALS"
	ROLLOUTSTRATEGY_NO_ROLLOUT        RolloutStrategy = "NO_ROLLOUT"
)

var allowedRolloutStrategyEnumValues = []RolloutStrategy{
	ROLLOUTSTRATEGY_UNIFORM_INTERVALS,
	ROLLOUTSTRATEGY_NO_ROLLOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RolloutStrategy) GetAllowedValues() []RolloutStrategy {
	return allowedRolloutStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RolloutStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RolloutStrategy(value)
	return nil
}

// NewRolloutStrategyFromValue returns a pointer to a valid RolloutStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRolloutStrategyFromValue(v string) (*RolloutStrategy, error) {
	ev := RolloutStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RolloutStrategy: valid values are %v", v, allowedRolloutStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RolloutStrategy) IsValid() bool {
	for _, existing := range allowedRolloutStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RolloutStrategy value.
func (v RolloutStrategy) Ptr() *RolloutStrategy {
	return &v
}
