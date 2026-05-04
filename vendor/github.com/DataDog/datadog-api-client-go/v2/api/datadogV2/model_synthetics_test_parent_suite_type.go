// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestParentSuiteType Type of the parent suite resource.
type SyntheticsTestParentSuiteType string

// List of SyntheticsTestParentSuiteType.
const (
	SYNTHETICSTESTPARENTSUITETYPE_PARENT_SUITE SyntheticsTestParentSuiteType = "parent_suite"
)

var allowedSyntheticsTestParentSuiteTypeEnumValues = []SyntheticsTestParentSuiteType{
	SYNTHETICSTESTPARENTSUITETYPE_PARENT_SUITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestParentSuiteType) GetAllowedValues() []SyntheticsTestParentSuiteType {
	return allowedSyntheticsTestParentSuiteTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestParentSuiteType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestParentSuiteType(value)
	return nil
}

// NewSyntheticsTestParentSuiteTypeFromValue returns a pointer to a valid SyntheticsTestParentSuiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestParentSuiteTypeFromValue(v string) (*SyntheticsTestParentSuiteType, error) {
	ev := SyntheticsTestParentSuiteType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestParentSuiteType: valid values are %v", v, allowedSyntheticsTestParentSuiteTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestParentSuiteType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestParentSuiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestParentSuiteType value.
func (v SyntheticsTestParentSuiteType) Ptr() *SyntheticsTestParentSuiteType {
	return &v
}
