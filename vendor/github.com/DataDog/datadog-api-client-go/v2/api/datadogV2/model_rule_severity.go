// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleSeverity Severity of a security rule.
type RuleSeverity string

// List of RuleSeverity.
const (
	RULESEVERITY_CRITICAL RuleSeverity = "critical"
	RULESEVERITY_HIGH     RuleSeverity = "high"
	RULESEVERITY_MEDIUM   RuleSeverity = "medium"
	RULESEVERITY_LOW      RuleSeverity = "low"
	RULESEVERITY_UNKNOWN  RuleSeverity = "unknown"
	RULESEVERITY_INFO     RuleSeverity = "info"
)

var allowedRuleSeverityEnumValues = []RuleSeverity{
	RULESEVERITY_CRITICAL,
	RULESEVERITY_HIGH,
	RULESEVERITY_MEDIUM,
	RULESEVERITY_LOW,
	RULESEVERITY_UNKNOWN,
	RULESEVERITY_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RuleSeverity) GetAllowedValues() []RuleSeverity {
	return allowedRuleSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RuleSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RuleSeverity(value)
	return nil
}

// NewRuleSeverityFromValue returns a pointer to a valid RuleSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRuleSeverityFromValue(v string) (*RuleSeverity, error) {
	ev := RuleSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RuleSeverity: valid values are %v", v, allowedRuleSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RuleSeverity) IsValid() bool {
	for _, existing := range allowedRuleSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleSeverity value.
func (v RuleSeverity) Ptr() *RuleSeverity {
	return &v
}
