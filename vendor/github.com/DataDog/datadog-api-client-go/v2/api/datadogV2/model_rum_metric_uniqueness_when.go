// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricUniquenessWhen When to count updatable events. `match` when the event is first seen, or `end` when the event is complete.
type RumMetricUniquenessWhen string

// List of RumMetricUniquenessWhen.
const (
	RUMMETRICUNIQUENESSWHEN_WHEN_MATCH RumMetricUniquenessWhen = "match"
	RUMMETRICUNIQUENESSWHEN_WHEN_END   RumMetricUniquenessWhen = "end"
)

var allowedRumMetricUniquenessWhenEnumValues = []RumMetricUniquenessWhen{
	RUMMETRICUNIQUENESSWHEN_WHEN_MATCH,
	RUMMETRICUNIQUENESSWHEN_WHEN_END,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumMetricUniquenessWhen) GetAllowedValues() []RumMetricUniquenessWhen {
	return allowedRumMetricUniquenessWhenEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumMetricUniquenessWhen) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumMetricUniquenessWhen(value)
	return nil
}

// NewRumMetricUniquenessWhenFromValue returns a pointer to a valid RumMetricUniquenessWhen
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumMetricUniquenessWhenFromValue(v string) (*RumMetricUniquenessWhen, error) {
	ev := RumMetricUniquenessWhen(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumMetricUniquenessWhen: valid values are %v", v, allowedRumMetricUniquenessWhenEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumMetricUniquenessWhen) IsValid() bool {
	for _, existing := range allowedRumMetricUniquenessWhenEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumMetricUniquenessWhen value.
func (v RumMetricUniquenessWhen) Ptr() *RumMetricUniquenessWhen {
	return &v
}
