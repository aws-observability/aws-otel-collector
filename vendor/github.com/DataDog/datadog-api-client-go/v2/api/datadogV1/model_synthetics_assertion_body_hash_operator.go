// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionBodyHashOperator Assertion operator to apply.
type SyntheticsAssertionBodyHashOperator string

// List of SyntheticsAssertionBodyHashOperator.
const (
	SYNTHETICSASSERTIONBODYHASHOPERATOR_MD5    SyntheticsAssertionBodyHashOperator = "md5"
	SYNTHETICSASSERTIONBODYHASHOPERATOR_SHA1   SyntheticsAssertionBodyHashOperator = "sha1"
	SYNTHETICSASSERTIONBODYHASHOPERATOR_SHA256 SyntheticsAssertionBodyHashOperator = "sha256"
)

var allowedSyntheticsAssertionBodyHashOperatorEnumValues = []SyntheticsAssertionBodyHashOperator{
	SYNTHETICSASSERTIONBODYHASHOPERATOR_MD5,
	SYNTHETICSASSERTIONBODYHASHOPERATOR_SHA1,
	SYNTHETICSASSERTIONBODYHASHOPERATOR_SHA256,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionBodyHashOperator) GetAllowedValues() []SyntheticsAssertionBodyHashOperator {
	return allowedSyntheticsAssertionBodyHashOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionBodyHashOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionBodyHashOperator(value)
	return nil
}

// NewSyntheticsAssertionBodyHashOperatorFromValue returns a pointer to a valid SyntheticsAssertionBodyHashOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionBodyHashOperatorFromValue(v string) (*SyntheticsAssertionBodyHashOperator, error) {
	ev := SyntheticsAssertionBodyHashOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionBodyHashOperator: valid values are %v", v, allowedSyntheticsAssertionBodyHashOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionBodyHashOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionBodyHashOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionBodyHashOperator value.
func (v SyntheticsAssertionBodyHashOperator) Ptr() *SyntheticsAssertionBodyHashOperator {
	return &v
}
