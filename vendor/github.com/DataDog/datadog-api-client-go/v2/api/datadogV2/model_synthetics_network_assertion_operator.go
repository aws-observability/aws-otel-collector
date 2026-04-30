// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionOperator Assertion operator to apply.
type SyntheticsNetworkAssertionOperator string

// List of SyntheticsNetworkAssertionOperator.
const (
	SYNTHETICSNETWORKASSERTIONOPERATOR_IS                 SyntheticsNetworkAssertionOperator = "is"
	SYNTHETICSNETWORKASSERTIONOPERATOR_IS_NOT             SyntheticsNetworkAssertionOperator = "isNot"
	SYNTHETICSNETWORKASSERTIONOPERATOR_LESS_THAN          SyntheticsNetworkAssertionOperator = "lessThan"
	SYNTHETICSNETWORKASSERTIONOPERATOR_LESS_THAN_OR_EQUAL SyntheticsNetworkAssertionOperator = "lessThanOrEqual"
	SYNTHETICSNETWORKASSERTIONOPERATOR_MORE_THAN          SyntheticsNetworkAssertionOperator = "moreThan"
	SYNTHETICSNETWORKASSERTIONOPERATOR_MORE_THAN_OR_EQUAL SyntheticsNetworkAssertionOperator = "moreThanOrEqual"
)

var allowedSyntheticsNetworkAssertionOperatorEnumValues = []SyntheticsNetworkAssertionOperator{
	SYNTHETICSNETWORKASSERTIONOPERATOR_IS,
	SYNTHETICSNETWORKASSERTIONOPERATOR_IS_NOT,
	SYNTHETICSNETWORKASSERTIONOPERATOR_LESS_THAN,
	SYNTHETICSNETWORKASSERTIONOPERATOR_LESS_THAN_OR_EQUAL,
	SYNTHETICSNETWORKASSERTIONOPERATOR_MORE_THAN,
	SYNTHETICSNETWORKASSERTIONOPERATOR_MORE_THAN_OR_EQUAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkAssertionOperator) GetAllowedValues() []SyntheticsNetworkAssertionOperator {
	return allowedSyntheticsNetworkAssertionOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkAssertionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkAssertionOperator(value)
	return nil
}

// NewSyntheticsNetworkAssertionOperatorFromValue returns a pointer to a valid SyntheticsNetworkAssertionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkAssertionOperatorFromValue(v string) (*SyntheticsNetworkAssertionOperator, error) {
	ev := SyntheticsNetworkAssertionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkAssertionOperator: valid values are %v", v, allowedSyntheticsNetworkAssertionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkAssertionOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkAssertionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkAssertionOperator value.
func (v SyntheticsNetworkAssertionOperator) Ptr() *SyntheticsNetworkAssertionOperator {
	return &v
}
