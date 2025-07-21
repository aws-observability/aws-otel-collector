// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAPITestStepSubtype The subtype of the Synthetic multi-step API test step.
type SyntheticsAPITestStepSubtype string

// List of SyntheticsAPITestStepSubtype.
const (
	SYNTHETICSAPITESTSTEPSUBTYPE_HTTP      SyntheticsAPITestStepSubtype = "http"
	SYNTHETICSAPITESTSTEPSUBTYPE_GRPC      SyntheticsAPITestStepSubtype = "grpc"
	SYNTHETICSAPITESTSTEPSUBTYPE_SSL       SyntheticsAPITestStepSubtype = "ssl"
	SYNTHETICSAPITESTSTEPSUBTYPE_DNS       SyntheticsAPITestStepSubtype = "dns"
	SYNTHETICSAPITESTSTEPSUBTYPE_TCP       SyntheticsAPITestStepSubtype = "tcp"
	SYNTHETICSAPITESTSTEPSUBTYPE_UDP       SyntheticsAPITestStepSubtype = "udp"
	SYNTHETICSAPITESTSTEPSUBTYPE_ICMP      SyntheticsAPITestStepSubtype = "icmp"
	SYNTHETICSAPITESTSTEPSUBTYPE_WEBSOCKET SyntheticsAPITestStepSubtype = "websocket"
)

var allowedSyntheticsAPITestStepSubtypeEnumValues = []SyntheticsAPITestStepSubtype{
	SYNTHETICSAPITESTSTEPSUBTYPE_HTTP,
	SYNTHETICSAPITESTSTEPSUBTYPE_GRPC,
	SYNTHETICSAPITESTSTEPSUBTYPE_SSL,
	SYNTHETICSAPITESTSTEPSUBTYPE_DNS,
	SYNTHETICSAPITESTSTEPSUBTYPE_TCP,
	SYNTHETICSAPITESTSTEPSUBTYPE_UDP,
	SYNTHETICSAPITESTSTEPSUBTYPE_ICMP,
	SYNTHETICSAPITESTSTEPSUBTYPE_WEBSOCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsAPITestStepSubtype) GetAllowedValues() []SyntheticsAPITestStepSubtype {
	return allowedSyntheticsAPITestStepSubtypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsAPITestStepSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsAPITestStepSubtype(value)
	return nil
}

// NewSyntheticsAPITestStepSubtypeFromValue returns a pointer to a valid SyntheticsAPITestStepSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsAPITestStepSubtypeFromValue(v string) (*SyntheticsAPITestStepSubtype, error) {
	ev := SyntheticsAPITestStepSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsAPITestStepSubtype: valid values are %v", v, allowedSyntheticsAPITestStepSubtypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsAPITestStepSubtype) IsValid() bool {
	for _, existing := range allowedSyntheticsAPITestStepSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAPITestStepSubtype value.
func (v SyntheticsAPITestStepSubtype) Ptr() *SyntheticsAPITestStepSubtype {
	return &v
}
