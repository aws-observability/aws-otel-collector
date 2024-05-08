// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthDigestType The type of basic authentication to use when performing the test.
type SyntheticsBasicAuthDigestType string

// List of SyntheticsBasicAuthDigestType.
const (
	SYNTHETICSBASICAUTHDIGESTTYPE_DIGEST SyntheticsBasicAuthDigestType = "digest"
)

var allowedSyntheticsBasicAuthDigestTypeEnumValues = []SyntheticsBasicAuthDigestType{
	SYNTHETICSBASICAUTHDIGESTTYPE_DIGEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthDigestType) GetAllowedValues() []SyntheticsBasicAuthDigestType {
	return allowedSyntheticsBasicAuthDigestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthDigestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthDigestType(value)
	return nil
}

// NewSyntheticsBasicAuthDigestTypeFromValue returns a pointer to a valid SyntheticsBasicAuthDigestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthDigestTypeFromValue(v string) (*SyntheticsBasicAuthDigestType, error) {
	ev := SyntheticsBasicAuthDigestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthDigestType: valid values are %v", v, allowedSyntheticsBasicAuthDigestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthDigestType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthDigestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthDigestType value.
func (v SyntheticsBasicAuthDigestType) Ptr() *SyntheticsBasicAuthDigestType {
	return &v
}
