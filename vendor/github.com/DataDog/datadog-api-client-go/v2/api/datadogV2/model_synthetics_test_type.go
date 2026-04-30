// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestType Type of the Synthetic test that produced this result.
type SyntheticsTestType string

// List of SyntheticsTestType.
const (
	SYNTHETICSTESTTYPE_API     SyntheticsTestType = "api"
	SYNTHETICSTESTTYPE_BROWSER SyntheticsTestType = "browser"
	SYNTHETICSTESTTYPE_MOBILE  SyntheticsTestType = "mobile"
	SYNTHETICSTESTTYPE_NETWORK SyntheticsTestType = "network"
)

var allowedSyntheticsTestTypeEnumValues = []SyntheticsTestType{
	SYNTHETICSTESTTYPE_API,
	SYNTHETICSTESTTYPE_BROWSER,
	SYNTHETICSTESTTYPE_MOBILE,
	SYNTHETICSTESTTYPE_NETWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestType) GetAllowedValues() []SyntheticsTestType {
	return allowedSyntheticsTestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestType(value)
	return nil
}

// NewSyntheticsTestTypeFromValue returns a pointer to a valid SyntheticsTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestTypeFromValue(v string) (*SyntheticsTestType, error) {
	ev := SyntheticsTestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestType: valid values are %v", v, allowedSyntheticsTestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestType value.
func (v SyntheticsTestType) Ptr() *SyntheticsTestType {
	return &v
}
