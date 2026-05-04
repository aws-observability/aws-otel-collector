// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestSubType Subtype of the Synthetic Network Path test: `tcp`, `udp`, or `icmp`.
type SyntheticsNetworkTestSubType string

// List of SyntheticsNetworkTestSubType.
const (
	SYNTHETICSNETWORKTESTSUBTYPE_TCP  SyntheticsNetworkTestSubType = "tcp"
	SYNTHETICSNETWORKTESTSUBTYPE_UDP  SyntheticsNetworkTestSubType = "udp"
	SYNTHETICSNETWORKTESTSUBTYPE_ICMP SyntheticsNetworkTestSubType = "icmp"
)

var allowedSyntheticsNetworkTestSubTypeEnumValues = []SyntheticsNetworkTestSubType{
	SYNTHETICSNETWORKTESTSUBTYPE_TCP,
	SYNTHETICSNETWORKTESTSUBTYPE_UDP,
	SYNTHETICSNETWORKTESTSUBTYPE_ICMP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkTestSubType) GetAllowedValues() []SyntheticsNetworkTestSubType {
	return allowedSyntheticsNetworkTestSubTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkTestSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkTestSubType(value)
	return nil
}

// NewSyntheticsNetworkTestSubTypeFromValue returns a pointer to a valid SyntheticsNetworkTestSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkTestSubTypeFromValue(v string) (*SyntheticsNetworkTestSubType, error) {
	ev := SyntheticsNetworkTestSubType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkTestSubType: valid values are %v", v, allowedSyntheticsNetworkTestSubTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkTestSubType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkTestSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkTestSubType value.
func (v SyntheticsNetworkTestSubType) Ptr() *SyntheticsNetworkTestSubType {
	return &v
}
