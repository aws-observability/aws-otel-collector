// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionXPathOperator Assertion operator to apply.
type SyntheticsAssertionXPathOperator string

// List of SyntheticsAssertionXPathOperator.
const (
	SYNTHETICSASSERTIONXPATHOPERATOR_VALIDATES_X_PATH SyntheticsAssertionXPathOperator = "validatesXPath"
)

var allowedSyntheticsAssertionXPathOperatorEnumValues = []SyntheticsAssertionXPathOperator{
	SYNTHETICSASSERTIONXPATHOPERATOR_VALIDATES_X_PATH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionXPathOperator) GetAllowedValues() []SyntheticsAssertionXPathOperator {
	return allowedSyntheticsAssertionXPathOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionXPathOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionXPathOperator(value)
	return nil
}

// NewSyntheticsAssertionXPathOperatorFromValue returns a pointer to a valid SyntheticsAssertionXPathOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionXPathOperatorFromValue(v string) (*SyntheticsAssertionXPathOperator, error) {
	ev := SyntheticsAssertionXPathOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionXPathOperator: valid values are %v", v, allowedSyntheticsAssertionXPathOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionXPathOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionXPathOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionXPathOperator value.
func (v SyntheticsAssertionXPathOperator) Ptr() *SyntheticsAssertionXPathOperator {
	return &v
}
