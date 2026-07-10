// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFilterVersionType The type of the resource. The value should always be `security_filters_configuration`.
type SecurityFilterVersionType string

// List of SecurityFilterVersionType.
const (
	SECURITYFILTERVERSIONTYPE_SECURITY_FILTERS_CONFIGURATION SecurityFilterVersionType = "security_filters_configuration"
)

var allowedSecurityFilterVersionTypeEnumValues = []SecurityFilterVersionType{
	SECURITYFILTERVERSIONTYPE_SECURITY_FILTERS_CONFIGURATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFilterVersionType) GetAllowedValues() []SecurityFilterVersionType {
	return allowedSecurityFilterVersionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFilterVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFilterVersionType(value)
	return nil
}

// NewSecurityFilterVersionTypeFromValue returns a pointer to a valid SecurityFilterVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFilterVersionTypeFromValue(v string) (*SecurityFilterVersionType, error) {
	ev := SecurityFilterVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFilterVersionType: valid values are %v", v, allowedSecurityFilterVersionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFilterVersionType) IsValid() bool {
	for _, existing := range allowedSecurityFilterVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFilterVersionType value.
func (v SecurityFilterVersionType) Ptr() *SecurityFilterVersionType {
	return &v
}
