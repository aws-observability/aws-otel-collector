// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionTimingsScope Timings scope for response time assertions.
type SyntheticsAssertionTimingsScope string

// List of SyntheticsAssertionTimingsScope.
const (
	SYNTHETICSASSERTIONTIMINGSSCOPE_ALL         SyntheticsAssertionTimingsScope = "all"
	SYNTHETICSASSERTIONTIMINGSSCOPE_WITHOUT_DNS SyntheticsAssertionTimingsScope = "withoutDNS"
)

var allowedSyntheticsAssertionTimingsScopeEnumValues = []SyntheticsAssertionTimingsScope{
	SYNTHETICSASSERTIONTIMINGSSCOPE_ALL,
	SYNTHETICSASSERTIONTIMINGSSCOPE_WITHOUT_DNS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionTimingsScope) GetAllowedValues() []SyntheticsAssertionTimingsScope {
	return allowedSyntheticsAssertionTimingsScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionTimingsScope) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionTimingsScope(value)
	return nil
}

// NewSyntheticsAssertionTimingsScopeFromValue returns a pointer to a valid SyntheticsAssertionTimingsScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionTimingsScopeFromValue(v string) (*SyntheticsAssertionTimingsScope, error) {
	ev := SyntheticsAssertionTimingsScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionTimingsScope: valid values are %v", v, allowedSyntheticsAssertionTimingsScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionTimingsScope) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionTimingsScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionTimingsScope value.
func (v SyntheticsAssertionTimingsScope) Ptr() *SyntheticsAssertionTimingsScope {
	return &v
}
