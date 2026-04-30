// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestDetailsType Type of the Synthetic test.
type SyntheticsTestDetailsType string

// List of SyntheticsTestDetailsType.
const (
	SYNTHETICSTESTDETAILSTYPE_API     SyntheticsTestDetailsType = "api"
	SYNTHETICSTESTDETAILSTYPE_BROWSER SyntheticsTestDetailsType = "browser"
	SYNTHETICSTESTDETAILSTYPE_MOBILE  SyntheticsTestDetailsType = "mobile"
	SYNTHETICSTESTDETAILSTYPE_NETWORK SyntheticsTestDetailsType = "network"
)

var allowedSyntheticsTestDetailsTypeEnumValues = []SyntheticsTestDetailsType{
	SYNTHETICSTESTDETAILSTYPE_API,
	SYNTHETICSTESTDETAILSTYPE_BROWSER,
	SYNTHETICSTESTDETAILSTYPE_MOBILE,
	SYNTHETICSTESTDETAILSTYPE_NETWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestDetailsType) GetAllowedValues() []SyntheticsTestDetailsType {
	return allowedSyntheticsTestDetailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestDetailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestDetailsType(value)
	return nil
}

// NewSyntheticsTestDetailsTypeFromValue returns a pointer to a valid SyntheticsTestDetailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestDetailsTypeFromValue(v string) (*SyntheticsTestDetailsType, error) {
	ev := SyntheticsTestDetailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestDetailsType: valid values are %v", v, allowedSyntheticsTestDetailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestDetailsType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestDetailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestDetailsType value.
func (v SyntheticsTestDetailsType) Ptr() *SyntheticsTestDetailsType {
	return &v
}
