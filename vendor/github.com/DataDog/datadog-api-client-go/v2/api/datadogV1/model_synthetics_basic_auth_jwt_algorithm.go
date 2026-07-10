// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsBasicAuthJWTAlgorithm Algorithm to use for the JWT authentication.
type SyntheticsBasicAuthJWTAlgorithm string

// List of SyntheticsBasicAuthJWTAlgorithm.
const (
	SYNTHETICSBASICAUTHJWTALGORITHM_HS256 SyntheticsBasicAuthJWTAlgorithm = "HS256"
	SYNTHETICSBASICAUTHJWTALGORITHM_RS256 SyntheticsBasicAuthJWTAlgorithm = "RS256"
	SYNTHETICSBASICAUTHJWTALGORITHM_ES256 SyntheticsBasicAuthJWTAlgorithm = "ES256"
)

var allowedSyntheticsBasicAuthJWTAlgorithmEnumValues = []SyntheticsBasicAuthJWTAlgorithm{
	SYNTHETICSBASICAUTHJWTALGORITHM_HS256,
	SYNTHETICSBASICAUTHJWTALGORITHM_RS256,
	SYNTHETICSBASICAUTHJWTALGORITHM_ES256,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsBasicAuthJWTAlgorithm) GetAllowedValues() []SyntheticsBasicAuthJWTAlgorithm {
	return allowedSyntheticsBasicAuthJWTAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsBasicAuthJWTAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsBasicAuthJWTAlgorithm(value)
	return nil
}

// NewSyntheticsBasicAuthJWTAlgorithmFromValue returns a pointer to a valid SyntheticsBasicAuthJWTAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsBasicAuthJWTAlgorithmFromValue(v string) (*SyntheticsBasicAuthJWTAlgorithm, error) {
	ev := SyntheticsBasicAuthJWTAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsBasicAuthJWTAlgorithm: valid values are %v", v, allowedSyntheticsBasicAuthJWTAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsBasicAuthJWTAlgorithm) IsValid() bool {
	for _, existing := range allowedSyntheticsBasicAuthJWTAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsBasicAuthJWTAlgorithm value.
func (v SyntheticsBasicAuthJWTAlgorithm) Ptr() *SyntheticsBasicAuthJWTAlgorithm {
	return &v
}
