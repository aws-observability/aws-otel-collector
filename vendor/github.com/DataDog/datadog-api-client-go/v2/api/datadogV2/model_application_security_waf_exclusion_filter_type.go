// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafExclusionFilterType Type of the resource. The value should always be `exclusion_filter`.
type ApplicationSecurityWafExclusionFilterType string

// List of ApplicationSecurityWafExclusionFilterType.
const (
	APPLICATIONSECURITYWAFEXCLUSIONFILTERTYPE_EXCLUSION_FILTER ApplicationSecurityWafExclusionFilterType = "exclusion_filter"
)

var allowedApplicationSecurityWafExclusionFilterTypeEnumValues = []ApplicationSecurityWafExclusionFilterType{
	APPLICATIONSECURITYWAFEXCLUSIONFILTERTYPE_EXCLUSION_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityWafExclusionFilterType) GetAllowedValues() []ApplicationSecurityWafExclusionFilterType {
	return allowedApplicationSecurityWafExclusionFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityWafExclusionFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityWafExclusionFilterType(value)
	return nil
}

// NewApplicationSecurityWafExclusionFilterTypeFromValue returns a pointer to a valid ApplicationSecurityWafExclusionFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityWafExclusionFilterTypeFromValue(v string) (*ApplicationSecurityWafExclusionFilterType, error) {
	ev := ApplicationSecurityWafExclusionFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityWafExclusionFilterType: valid values are %v", v, allowedApplicationSecurityWafExclusionFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityWafExclusionFilterType) IsValid() bool {
	for _, existing := range allowedApplicationSecurityWafExclusionFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityWafExclusionFilterType value.
func (v ApplicationSecurityWafExclusionFilterType) Ptr() *ApplicationSecurityWafExclusionFilterType {
	return &v
}
