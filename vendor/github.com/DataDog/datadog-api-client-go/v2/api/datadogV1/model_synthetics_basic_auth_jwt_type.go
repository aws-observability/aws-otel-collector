// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthJWTType The type of authentication to use when performing the test.
type SyntheticsBasicAuthJWTType string

// List of SyntheticsBasicAuthJWTType.
const (
	SYNTHETICSBASICAUTHJWTTYPE_JWT SyntheticsBasicAuthJWTType = "jwt"
)

var allowedSyntheticsBasicAuthJWTTypeEnumValues = []SyntheticsBasicAuthJWTType{
	SYNTHETICSBASICAUTHJWTTYPE_JWT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthJWTType) GetAllowedValues() []SyntheticsBasicAuthJWTType {
	return allowedSyntheticsBasicAuthJWTTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthJWTType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthJWTType(value)
	return nil
}

// NewSyntheticsBasicAuthJWTTypeFromValue returns a pointer to a valid SyntheticsBasicAuthJWTType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthJWTTypeFromValue(v string) (*SyntheticsBasicAuthJWTType, error) {
	ev := SyntheticsBasicAuthJWTType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthJWTType: valid values are %v", v, allowedSyntheticsBasicAuthJWTTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthJWTType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthJWTTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthJWTType value.
func (v SyntheticsBasicAuthJWTType) Ptr() *SyntheticsBasicAuthJWTType {
	return &v
}
