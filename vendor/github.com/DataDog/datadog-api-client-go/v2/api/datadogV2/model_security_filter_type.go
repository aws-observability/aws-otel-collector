// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFilterType The type of the resource. The value should always be `security_filters`.
type SecurityFilterType string

// List of SecurityFilterType.
const (
	SECURITYFILTERTYPE_SECURITY_FILTERS SecurityFilterType = "security_filters"
)

var allowedSecurityFilterTypeEnumValues = []SecurityFilterType{
	SECURITYFILTERTYPE_SECURITY_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFilterType) GetAllowedValues() []SecurityFilterType {
	return allowedSecurityFilterTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFilterType(value)
	return nil
}

// NewSecurityFilterTypeFromValue returns a pointer to a valid SecurityFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFilterTypeFromValue(v string) (*SecurityFilterType, error) {
	ev := SecurityFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFilterType: valid values are %v", v, allowedSecurityFilterTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFilterType) IsValid() bool {
	for _, existing := range allowedSecurityFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFilterType value.
func (v SecurityFilterType) Ptr() *SecurityFilterType {
	return &v
}
