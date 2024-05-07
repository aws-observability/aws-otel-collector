// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthSigv4Type The type of authentication to use when performing the test.
type SyntheticsBasicAuthSigv4Type string

// List of SyntheticsBasicAuthSigv4Type.
const (
	SYNTHETICSBASICAUTHSIGV4TYPE_SIGV4 SyntheticsBasicAuthSigv4Type = "sigv4"
)

var allowedSyntheticsBasicAuthSigv4TypeEnumValues = []SyntheticsBasicAuthSigv4Type{
	SYNTHETICSBASICAUTHSIGV4TYPE_SIGV4,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthSigv4Type) GetAllowedValues() []SyntheticsBasicAuthSigv4Type {
	return allowedSyntheticsBasicAuthSigv4TypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthSigv4Type) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthSigv4Type(value)
	return nil
}

// NewSyntheticsBasicAuthSigv4TypeFromValue returns a pointer to a valid SyntheticsBasicAuthSigv4Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthSigv4TypeFromValue(v string) (*SyntheticsBasicAuthSigv4Type, error) {
	ev := SyntheticsBasicAuthSigv4Type(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthSigv4Type: valid values are %v", v, allowedSyntheticsBasicAuthSigv4TypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthSigv4Type) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthSigv4TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthSigv4Type value.
func (v SyntheticsBasicAuthSigv4Type) Ptr() *SyntheticsBasicAuthSigv4Type {
	return &v
}
