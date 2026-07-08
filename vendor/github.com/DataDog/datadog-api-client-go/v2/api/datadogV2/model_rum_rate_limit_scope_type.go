// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitScopeType The type of scope the rate limit configuration applies to.
type RumRateLimitScopeType string

// List of RumRateLimitScopeType.
const (
	RUMRATELIMITSCOPETYPE_APPLICATION RumRateLimitScopeType = "application"
)

var allowedRumRateLimitScopeTypeEnumValues = []RumRateLimitScopeType{
	RUMRATELIMITSCOPETYPE_APPLICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRateLimitScopeType) GetAllowedValues() []RumRateLimitScopeType {
	return allowedRumRateLimitScopeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRateLimitScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRateLimitScopeType(value)
	return nil
}

// NewRumRateLimitScopeTypeFromValue returns a pointer to a valid RumRateLimitScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRateLimitScopeTypeFromValue(v string) (*RumRateLimitScopeType, error) {
	ev := RumRateLimitScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRateLimitScopeType: valid values are %v", v, allowedRumRateLimitScopeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRateLimitScopeType) IsValid() bool {
	for _, existing := range allowedRumRateLimitScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRateLimitScopeType value.
func (v RumRateLimitScopeType) Ptr() *RumRateLimitScopeType {
	return &v
}
