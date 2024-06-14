// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionJSONSchemaOperator Assertion operator to apply.
type SyntheticsAssertionJSONSchemaOperator string

// List of SyntheticsAssertionJSONSchemaOperator.
const (
	SYNTHETICSASSERTIONJSONSCHEMAOPERATOR_VALIDATES_JSON_SCHEMA SyntheticsAssertionJSONSchemaOperator = "validatesJSONSchema"
)

var allowedSyntheticsAssertionJSONSchemaOperatorEnumValues = []SyntheticsAssertionJSONSchemaOperator{
	SYNTHETICSASSERTIONJSONSCHEMAOPERATOR_VALIDATES_JSON_SCHEMA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionJSONSchemaOperator) GetAllowedValues() []SyntheticsAssertionJSONSchemaOperator {
	return allowedSyntheticsAssertionJSONSchemaOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionJSONSchemaOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionJSONSchemaOperator(value)
	return nil
}

// NewSyntheticsAssertionJSONSchemaOperatorFromValue returns a pointer to a valid SyntheticsAssertionJSONSchemaOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionJSONSchemaOperatorFromValue(v string) (*SyntheticsAssertionJSONSchemaOperator, error) {
	ev := SyntheticsAssertionJSONSchemaOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionJSONSchemaOperator: valid values are %v", v, allowedSyntheticsAssertionJSONSchemaOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionJSONSchemaOperator) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionJSONSchemaOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionJSONSchemaOperator value.
func (v SyntheticsAssertionJSONSchemaOperator) Ptr() *SyntheticsAssertionJSONSchemaOperator {
	return &v
}
