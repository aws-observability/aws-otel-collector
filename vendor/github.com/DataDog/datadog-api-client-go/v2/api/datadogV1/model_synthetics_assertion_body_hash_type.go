// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionBodyHashType Type of the assertion.
type SyntheticsAssertionBodyHashType string

// List of SyntheticsAssertionBodyHashType.
const (
	SYNTHETICSASSERTIONBODYHASHTYPE_BODY_HASH SyntheticsAssertionBodyHashType = "bodyHash"
)

var allowedSyntheticsAssertionBodyHashTypeEnumValues = []SyntheticsAssertionBodyHashType{
	SYNTHETICSASSERTIONBODYHASHTYPE_BODY_HASH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionBodyHashType) GetAllowedValues() []SyntheticsAssertionBodyHashType {
	return allowedSyntheticsAssertionBodyHashTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionBodyHashType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionBodyHashType(value)
	return nil
}

// NewSyntheticsAssertionBodyHashTypeFromValue returns a pointer to a valid SyntheticsAssertionBodyHashType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionBodyHashTypeFromValue(v string) (*SyntheticsAssertionBodyHashType, error) {
	ev := SyntheticsAssertionBodyHashType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionBodyHashType: valid values are %v", v, allowedSyntheticsAssertionBodyHashTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionBodyHashType) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionBodyHashTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionBodyHashType value.
func (v SyntheticsAssertionBodyHashType) Ptr() *SyntheticsAssertionBodyHashType {
	return &v
}
