// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestRequestTCPMethod For TCP tests, the TCP traceroute strategy.
type SyntheticsNetworkTestRequestTCPMethod string

// List of SyntheticsNetworkTestRequestTCPMethod.
const (
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_PREFER_SACK SyntheticsNetworkTestRequestTCPMethod = "prefer_sack"
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_SYN         SyntheticsNetworkTestRequestTCPMethod = "syn"
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_SACK        SyntheticsNetworkTestRequestTCPMethod = "sack"
)

var allowedSyntheticsNetworkTestRequestTCPMethodEnumValues = []SyntheticsNetworkTestRequestTCPMethod{
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_PREFER_SACK,
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_SYN,
	SYNTHETICSNETWORKTESTREQUESTTCPMETHOD_SACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkTestRequestTCPMethod) GetAllowedValues() []SyntheticsNetworkTestRequestTCPMethod {
	return allowedSyntheticsNetworkTestRequestTCPMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkTestRequestTCPMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkTestRequestTCPMethod(value)
	return nil
}

// NewSyntheticsNetworkTestRequestTCPMethodFromValue returns a pointer to a valid SyntheticsNetworkTestRequestTCPMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkTestRequestTCPMethodFromValue(v string) (*SyntheticsNetworkTestRequestTCPMethod, error) {
	ev := SyntheticsNetworkTestRequestTCPMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkTestRequestTCPMethod: valid values are %v", v, allowedSyntheticsNetworkTestRequestTCPMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkTestRequestTCPMethod) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkTestRequestTCPMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkTestRequestTCPMethod value.
func (v SyntheticsNetworkTestRequestTCPMethod) Ptr() *SyntheticsNetworkTestRequestTCPMethod {
	return &v
}
