// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafExclusionFilterOnMatch The action taken when the exclusion filter matches. When set to `monitor`, security traces are emitted but the requests are not blocked. By default, security traces are not emitted and the requests are not blocked.
type ApplicationSecurityWafExclusionFilterOnMatch string

// List of ApplicationSecurityWafExclusionFilterOnMatch.
const (
	APPLICATIONSECURITYWAFEXCLUSIONFILTERONMATCH_MONITOR ApplicationSecurityWafExclusionFilterOnMatch = "monitor"
)

var allowedApplicationSecurityWafExclusionFilterOnMatchEnumValues = []ApplicationSecurityWafExclusionFilterOnMatch{
	APPLICATIONSECURITYWAFEXCLUSIONFILTERONMATCH_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafExclusionFilterOnMatch) GetAllowedValues() []ApplicationSecurityWafExclusionFilterOnMatch {
	return allowedApplicationSecurityWafExclusionFilterOnMatchEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafExclusionFilterOnMatch) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafExclusionFilterOnMatch(value)
	return nil
}

// NewApplicationSecurityWafExclusionFilterOnMatchFromValue returns a pointer to a valid ApplicationSecurityWafExclusionFilterOnMatch
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafExclusionFilterOnMatchFromValue(v string) (*ApplicationSecurityWafExclusionFilterOnMatch, error) {
	ev := ApplicationSecurityWafExclusionFilterOnMatch(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafExclusionFilterOnMatch: valid values are %v", v, allowedApplicationSecurityWafExclusionFilterOnMatchEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafExclusionFilterOnMatch) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafExclusionFilterOnMatchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafExclusionFilterOnMatch value.
func (v ApplicationSecurityWafExclusionFilterOnMatch) Ptr() *ApplicationSecurityWafExclusionFilterOnMatch {
	return &v
}
