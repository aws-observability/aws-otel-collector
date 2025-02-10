// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountType AWS Account resource type.
type AWSAccountType string

// List of AWSAccountType.
const (
	AWSACCOUNTTYPE_ACCOUNT AWSAccountType = "account"
)

var allowedAWSAccountTypeEnumValues = []AWSAccountType{
	AWSACCOUNTTYPE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSAccountType) GetAllowedValues() []AWSAccountType {
	return allowedAWSAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSAccountType(value)
	return nil
}

// NewAWSAccountTypeFromValue returns a pointer to a valid AWSAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSAccountTypeFromValue(v string) (*AWSAccountType, error) {
	ev := AWSAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSAccountType: valid values are %v", v, allowedAWSAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSAccountType) IsValid() bool {
	for _, existing := range allowedAWSAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSAccountType value.
func (v AWSAccountType) Ptr() *AWSAccountType {
	return &v
}
