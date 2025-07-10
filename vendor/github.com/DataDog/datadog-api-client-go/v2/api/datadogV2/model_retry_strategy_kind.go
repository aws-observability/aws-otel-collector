// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetryStrategyKind The definition of `RetryStrategyKind` object.
type RetryStrategyKind string

// List of RetryStrategyKind.
const (
	RETRYSTRATEGYKIND_RETRY_STRATEGY_LINEAR RetryStrategyKind = "RETRY_STRATEGY_LINEAR"
)

var allowedRetryStrategyKindEnumValues = []RetryStrategyKind{
	RETRYSTRATEGYKIND_RETRY_STRATEGY_LINEAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetryStrategyKind) GetAllowedValues() []RetryStrategyKind {
	return allowedRetryStrategyKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetryStrategyKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetryStrategyKind(value)
	return nil
}

// NewRetryStrategyKindFromValue returns a pointer to a valid RetryStrategyKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetryStrategyKindFromValue(v string) (*RetryStrategyKind, error) {
	ev := RetryStrategyKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetryStrategyKind: valid values are %v", v, allowedRetryStrategyKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetryStrategyKind) IsValid() bool {
	for _, existing := range allowedRetryStrategyKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetryStrategyKind value.
func (v RetryStrategyKind) Ptr() *RetryStrategyKind {
	return &v
}
