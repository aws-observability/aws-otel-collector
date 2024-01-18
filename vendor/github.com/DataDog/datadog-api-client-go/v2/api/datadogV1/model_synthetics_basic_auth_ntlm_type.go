// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthNTLMType The type of authentication to use when performing the test.
type SyntheticsBasicAuthNTLMType string

// List of SyntheticsBasicAuthNTLMType.
const (
	SYNTHETICSBASICAUTHNTLMTYPE_NTLM SyntheticsBasicAuthNTLMType = "ntlm"
)

var allowedSyntheticsBasicAuthNTLMTypeEnumValues = []SyntheticsBasicAuthNTLMType{
	SYNTHETICSBASICAUTHNTLMTYPE_NTLM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthNTLMType) GetAllowedValues() []SyntheticsBasicAuthNTLMType {
	return allowedSyntheticsBasicAuthNTLMTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthNTLMType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthNTLMType(value)
	return nil
}

// NewSyntheticsBasicAuthNTLMTypeFromValue returns a pointer to a valid SyntheticsBasicAuthNTLMType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthNTLMTypeFromValue(v string) (*SyntheticsBasicAuthNTLMType, error) {
	ev := SyntheticsBasicAuthNTLMType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthNTLMType: valid values are %v", v, allowedSyntheticsBasicAuthNTLMTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthNTLMType) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthNTLMTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthNTLMType value.
func (v SyntheticsBasicAuthNTLMType) Ptr() *SyntheticsBasicAuthNTLMType {
	return &v
}
