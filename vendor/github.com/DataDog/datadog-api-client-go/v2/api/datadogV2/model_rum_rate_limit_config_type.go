// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitConfigType The type of the resource, always `rum_rate_limit_config`.
type RumRateLimitConfigType string

// List of RumRateLimitConfigType.
const (
	RUMRATELIMITCONFIGTYPE_RUM_RATE_LIMIT_CONFIG RumRateLimitConfigType = "rum_rate_limit_config"
)

var allowedRumRateLimitConfigTypeEnumValues = []RumRateLimitConfigType{
	RUMRATELIMITCONFIGTYPE_RUM_RATE_LIMIT_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRateLimitConfigType) GetAllowedValues() []RumRateLimitConfigType {
	return allowedRumRateLimitConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRateLimitConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRateLimitConfigType(value)
	return nil
}

// NewRumRateLimitConfigTypeFromValue returns a pointer to a valid RumRateLimitConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRateLimitConfigTypeFromValue(v string) (*RumRateLimitConfigType, error) {
	ev := RumRateLimitConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRateLimitConfigType: valid values are %v", v, allowedRumRateLimitConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRateLimitConfigType) IsValid() bool {
	for _, existing := range allowedRumRateLimitConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRateLimitConfigType value.
func (v RumRateLimitConfigType) Ptr() *RumRateLimitConfigType {
	return &v
}
