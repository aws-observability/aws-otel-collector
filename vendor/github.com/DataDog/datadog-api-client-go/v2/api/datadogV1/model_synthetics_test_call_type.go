// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestCallType The type of gRPC call to perform.
type SyntheticsTestCallType string

// List of SyntheticsTestCallType.
const (
	SYNTHETICSTESTCALLTYPE_HEALTHCHECK SyntheticsTestCallType = "healthcheck"
	SYNTHETICSTESTCALLTYPE_UNARY       SyntheticsTestCallType = "unary"
)

var allowedSyntheticsTestCallTypeEnumValues = []SyntheticsTestCallType{
	SYNTHETICSTESTCALLTYPE_HEALTHCHECK,
	SYNTHETICSTESTCALLTYPE_UNARY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsTestCallType) GetAllowedValues() []SyntheticsTestCallType {
	return allowedSyntheticsTestCallTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestCallType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestCallType(value)
	return nil
}

// NewSyntheticsTestCallTypeFromValue returns a pointer to a valid SyntheticsTestCallType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestCallTypeFromValue(v string) (*SyntheticsTestCallType, error) {
	ev := SyntheticsTestCallType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestCallType: valid values are %v", v, allowedSyntheticsTestCallTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestCallType) IsValid() bool {
	for _, existing := range allowedSyntheticsTestCallTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestCallType value.
func (v SyntheticsTestCallType) Ptr() *SyntheticsTestCallType {
	return &v
}
