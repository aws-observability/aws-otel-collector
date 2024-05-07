// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSRelatedAccountType Type of AWS related account.
type AWSRelatedAccountType string

// List of AWSRelatedAccountType.
const (
	AWSRELATEDACCOUNTTYPE_AWS_ACCOUNT AWSRelatedAccountType = "aws_account"
)

var allowedAWSRelatedAccountTypeEnumValues = []AWSRelatedAccountType{
	AWSRELATEDACCOUNTTYPE_AWS_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSRelatedAccountType) GetAllowedValues() []AWSRelatedAccountType {
	return allowedAWSRelatedAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSRelatedAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSRelatedAccountType(value)
	return nil
}

// NewAWSRelatedAccountTypeFromValue returns a pointer to a valid AWSRelatedAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSRelatedAccountTypeFromValue(v string) (*AWSRelatedAccountType, error) {
	ev := AWSRelatedAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSRelatedAccountType: valid values are %v", v, allowedAWSRelatedAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSRelatedAccountType) IsValid() bool {
	for _, existing := range allowedAWSRelatedAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSRelatedAccountType value.
func (v AWSRelatedAccountType) Ptr() *AWSRelatedAccountType {
	return &v
}
