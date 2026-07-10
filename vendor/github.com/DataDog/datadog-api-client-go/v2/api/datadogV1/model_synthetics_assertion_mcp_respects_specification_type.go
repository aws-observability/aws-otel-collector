// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionMCPRespectsSpecificationType Type of the assertion.
type SyntheticsAssertionMCPRespectsSpecificationType string

// List of SyntheticsAssertionMCPRespectsSpecificationType.
const (
	SYNTHETICSASSERTIONMCPRESPECTSSPECIFICATIONTYPE_MCP_RESPECTS_SPECIFICATION SyntheticsAssertionMCPRespectsSpecificationType = "mcpRespectsSpecification"
)

var allowedSyntheticsAssertionMCPRespectsSpecificationTypeEnumValues = []SyntheticsAssertionMCPRespectsSpecificationType{
	SYNTHETICSASSERTIONMCPRESPECTSSPECIFICATIONTYPE_MCP_RESPECTS_SPECIFICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAssertionMCPRespectsSpecificationType) GetAllowedValues() []SyntheticsAssertionMCPRespectsSpecificationType {
	return allowedSyntheticsAssertionMCPRespectsSpecificationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAssertionMCPRespectsSpecificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAssertionMCPRespectsSpecificationType(value)
	return nil
}

// NewSyntheticsAssertionMCPRespectsSpecificationTypeFromValue returns a pointer to a valid SyntheticsAssertionMCPRespectsSpecificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAssertionMCPRespectsSpecificationTypeFromValue(v string) (*SyntheticsAssertionMCPRespectsSpecificationType, error) {
	ev := SyntheticsAssertionMCPRespectsSpecificationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAssertionMCPRespectsSpecificationType: valid values are %v", v, allowedSyntheticsAssertionMCPRespectsSpecificationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAssertionMCPRespectsSpecificationType) IsValid() bool {
	for _, existing := range allowedSyntheticsAssertionMCPRespectsSpecificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAssertionMCPRespectsSpecificationType value.
func (v SyntheticsAssertionMCPRespectsSpecificationType) Ptr() *SyntheticsAssertionMCPRespectsSpecificationType {
	return &v
}
