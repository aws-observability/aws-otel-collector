// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionType Type of the version resource.
type SyntheticsTestVersionType string

// List of SyntheticsTestVersionType.
const (
	SYNTHETICSTESTVERSIONTYPE_VERSION SyntheticsTestVersionType = "version"
)

var allowedSyntheticsTestVersionTypeEnumValues = []SyntheticsTestVersionType{
	SYNTHETICSTESTVERSIONTYPE_VERSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestVersionType) GetAllowedValues() []SyntheticsTestVersionType {
	return allowedSyntheticsTestVersionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestVersionType(value)
	return nil
}

// NewSyntheticsTestVersionTypeFromValue returns a pointer to a valid SyntheticsTestVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestVersionTypeFromValue(v string) (*SyntheticsTestVersionType, error) {
	ev := SyntheticsTestVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestVersionType: valid values are %v", v, allowedSyntheticsTestVersionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestVersionType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestVersionType value.
func (v SyntheticsTestVersionType) Ptr() *SyntheticsTestVersionType {
	return &v
}
