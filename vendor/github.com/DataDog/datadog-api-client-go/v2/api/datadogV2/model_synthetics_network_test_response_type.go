// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestResponseType Type of response, `network_test`.
type SyntheticsNetworkTestResponseType string

// List of SyntheticsNetworkTestResponseType.
const (
	SYNTHETICSNETWORKTESTRESPONSETYPE_NETWORK_TEST SyntheticsNetworkTestResponseType = "network_test"
)

var allowedSyntheticsNetworkTestResponseTypeEnumValues = []SyntheticsNetworkTestResponseType{
	SYNTHETICSNETWORKTESTRESPONSETYPE_NETWORK_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkTestResponseType) GetAllowedValues() []SyntheticsNetworkTestResponseType {
	return allowedSyntheticsNetworkTestResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkTestResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkTestResponseType(value)
	return nil
}

// NewSyntheticsNetworkTestResponseTypeFromValue returns a pointer to a valid SyntheticsNetworkTestResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkTestResponseTypeFromValue(v string) (*SyntheticsNetworkTestResponseType, error) {
	ev := SyntheticsNetworkTestResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkTestResponseType: valid values are %v", v, allowedSyntheticsNetworkTestResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkTestResponseType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkTestResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkTestResponseType value.
func (v SyntheticsNetworkTestResponseType) Ptr() *SyntheticsNetworkTestResponseType {
	return &v
}
