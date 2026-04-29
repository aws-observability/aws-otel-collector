// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionPacketLossPercentageType Type of the packet loss percentage assertion.
type SyntheticsNetworkAssertionPacketLossPercentageType string

// List of SyntheticsNetworkAssertionPacketLossPercentageType.
const (
	SYNTHETICSNETWORKASSERTIONPACKETLOSSPERCENTAGETYPE_PACKET_LOSS_PERCENTAGE SyntheticsNetworkAssertionPacketLossPercentageType = "packetLossPercentage"
)

var allowedSyntheticsNetworkAssertionPacketLossPercentageTypeEnumValues = []SyntheticsNetworkAssertionPacketLossPercentageType{
	SYNTHETICSNETWORKASSERTIONPACKETLOSSPERCENTAGETYPE_PACKET_LOSS_PERCENTAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkAssertionPacketLossPercentageType) GetAllowedValues() []SyntheticsNetworkAssertionPacketLossPercentageType {
	return allowedSyntheticsNetworkAssertionPacketLossPercentageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkAssertionPacketLossPercentageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkAssertionPacketLossPercentageType(value)
	return nil
}

// NewSyntheticsNetworkAssertionPacketLossPercentageTypeFromValue returns a pointer to a valid SyntheticsNetworkAssertionPacketLossPercentageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkAssertionPacketLossPercentageTypeFromValue(v string) (*SyntheticsNetworkAssertionPacketLossPercentageType, error) {
	ev := SyntheticsNetworkAssertionPacketLossPercentageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkAssertionPacketLossPercentageType: valid values are %v", v, allowedSyntheticsNetworkAssertionPacketLossPercentageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkAssertionPacketLossPercentageType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkAssertionPacketLossPercentageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkAssertionPacketLossPercentageType value.
func (v SyntheticsNetworkAssertionPacketLossPercentageType) Ptr() *SyntheticsNetworkAssertionPacketLossPercentageType {
	return &v
}
