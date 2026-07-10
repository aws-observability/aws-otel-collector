// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRateLimitQuotaReachedAction The action to take when the session quota is reached.
type RumRateLimitQuotaReachedAction string

// List of RumRateLimitQuotaReachedAction.
const (
	RUMRATELIMITQUOTAREACHEDACTION_STOP     RumRateLimitQuotaReachedAction = "stop"
	RUMRATELIMITQUOTAREACHEDACTION_SLOWDOWN RumRateLimitQuotaReachedAction = "slowdown"
)

var allowedRumRateLimitQuotaReachedActionEnumValues = []RumRateLimitQuotaReachedAction{
	RUMRATELIMITQUOTAREACHEDACTION_STOP,
	RUMRATELIMITQUOTAREACHEDACTION_SLOWDOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRateLimitQuotaReachedAction) GetAllowedValues() []RumRateLimitQuotaReachedAction {
	return allowedRumRateLimitQuotaReachedActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRateLimitQuotaReachedAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRateLimitQuotaReachedAction(value)
	return nil
}

// NewRumRateLimitQuotaReachedActionFromValue returns a pointer to a valid RumRateLimitQuotaReachedAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRateLimitQuotaReachedActionFromValue(v string) (*RumRateLimitQuotaReachedAction, error) {
	ev := RumRateLimitQuotaReachedAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRateLimitQuotaReachedAction: valid values are %v", v, allowedRumRateLimitQuotaReachedActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRateLimitQuotaReachedAction) IsValid() bool {
	for _, existing := range allowedRumRateLimitQuotaReachedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRateLimitQuotaReachedAction value.
func (v RumRateLimitQuotaReachedAction) Ptr() *RumRateLimitQuotaReachedAction {
	return &v
}
