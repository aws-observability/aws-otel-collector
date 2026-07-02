// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitMode The rate limit mode. `custom` enforces a fixed session limit, while
// `adaptive` dynamically adjusts retention.
type RumRateLimitMode string

// List of RumRateLimitMode.
const (
	RUMRATELIMITMODE_CUSTOM   RumRateLimitMode = "custom"
	RUMRATELIMITMODE_ADAPTIVE RumRateLimitMode = "adaptive"
)

var allowedRumRateLimitModeEnumValues = []RumRateLimitMode{
	RUMRATELIMITMODE_CUSTOM,
	RUMRATELIMITMODE_ADAPTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRateLimitMode) GetAllowedValues() []RumRateLimitMode {
	return allowedRumRateLimitModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRateLimitMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRateLimitMode(value)
	return nil
}

// NewRumRateLimitModeFromValue returns a pointer to a valid RumRateLimitMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRateLimitModeFromValue(v string) (*RumRateLimitMode, error) {
	ev := RumRateLimitMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRateLimitMode: valid values are %v", v, allowedRumRateLimitModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRateLimitMode) IsValid() bool {
	for _, existing := range allowedRumRateLimitModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRateLimitMode value.
func (v RumRateLimitMode) Ptr() *RumRateLimitMode {
	return &v
}
