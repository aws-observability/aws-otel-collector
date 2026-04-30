// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkAssertionMultiNetworkHopType Type of the multi-network hop assertion.
type SyntheticsNetworkAssertionMultiNetworkHopType string

// List of SyntheticsNetworkAssertionMultiNetworkHopType.
const (
	SYNTHETICSNETWORKASSERTIONMULTINETWORKHOPTYPE_MULTI_NETWORK_HOP SyntheticsNetworkAssertionMultiNetworkHopType = "multiNetworkHop"
)

var allowedSyntheticsNetworkAssertionMultiNetworkHopTypeEnumValues = []SyntheticsNetworkAssertionMultiNetworkHopType{
	SYNTHETICSNETWORKASSERTIONMULTINETWORKHOPTYPE_MULTI_NETWORK_HOP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsNetworkAssertionMultiNetworkHopType) GetAllowedValues() []SyntheticsNetworkAssertionMultiNetworkHopType {
	return allowedSyntheticsNetworkAssertionMultiNetworkHopTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsNetworkAssertionMultiNetworkHopType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsNetworkAssertionMultiNetworkHopType(value)
	return nil
}

// NewSyntheticsNetworkAssertionMultiNetworkHopTypeFromValue returns a pointer to a valid SyntheticsNetworkAssertionMultiNetworkHopType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsNetworkAssertionMultiNetworkHopTypeFromValue(v string) (*SyntheticsNetworkAssertionMultiNetworkHopType, error) {
	ev := SyntheticsNetworkAssertionMultiNetworkHopType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsNetworkAssertionMultiNetworkHopType: valid values are %v", v, allowedSyntheticsNetworkAssertionMultiNetworkHopTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsNetworkAssertionMultiNetworkHopType) IsValid() bool {
	for _, existing := range allowedSyntheticsNetworkAssertionMultiNetworkHopTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsNetworkAssertionMultiNetworkHopType value.
func (v SyntheticsNetworkAssertionMultiNetworkHopType) Ptr() *SyntheticsNetworkAssertionMultiNetworkHopType {
	return &v
}
