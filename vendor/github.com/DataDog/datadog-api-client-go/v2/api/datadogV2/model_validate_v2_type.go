// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValidateV2Type Resource type for the API key validation response.
type ValidateV2Type string

// List of ValidateV2Type.
const (
	VALIDATEV2TYPE_ValidateV2 ValidateV2Type = "validate_v2"
)

var allowedValidateV2TypeEnumValues = []ValidateV2Type{
	VALIDATEV2TYPE_ValidateV2,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ValidateV2Type) GetAllowedValues() []ValidateV2Type {
	return allowedValidateV2TypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ValidateV2Type) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ValidateV2Type(value)
	return nil
}

// NewValidateV2TypeFromValue returns a pointer to a valid ValidateV2Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewValidateV2TypeFromValue(v string) (*ValidateV2Type, error) {
	ev := ValidateV2Type(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ValidateV2Type: valid values are %v", v, allowedValidateV2TypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ValidateV2Type) IsValid() bool {
	for _, existing := range allowedValidateV2TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidateV2Type value.
func (v ValidateV2Type) Ptr() *ValidateV2Type {
	return &v
}
