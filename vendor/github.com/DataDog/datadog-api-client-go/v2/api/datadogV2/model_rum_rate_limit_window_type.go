// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitWindowType The window type over which the session limit is enforced.
type RumRateLimitWindowType string

// List of RumRateLimitWindowType.
const (
	RUMRATELIMITWINDOWTYPE_DAILY RumRateLimitWindowType = "daily"
)

var allowedRumRateLimitWindowTypeEnumValues = []RumRateLimitWindowType{
	RUMRATELIMITWINDOWTYPE_DAILY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRateLimitWindowType) GetAllowedValues() []RumRateLimitWindowType {
	return allowedRumRateLimitWindowTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRateLimitWindowType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRateLimitWindowType(value)
	return nil
}

// NewRumRateLimitWindowTypeFromValue returns a pointer to a valid RumRateLimitWindowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRateLimitWindowTypeFromValue(v string) (*RumRateLimitWindowType, error) {
	ev := RumRateLimitWindowType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRateLimitWindowType: valid values are %v", v, allowedRumRateLimitWindowTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRateLimitWindowType) IsValid() bool {
	for _, existing := range allowedRumRateLimitWindowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRateLimitWindowType value.
func (v RumRateLimitWindowType) Ptr() *RumRateLimitWindowType {
	return &v
}
