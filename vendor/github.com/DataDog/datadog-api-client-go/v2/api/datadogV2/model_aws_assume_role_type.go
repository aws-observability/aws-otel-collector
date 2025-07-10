// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAssumeRoleType The definition of `AWSAssumeRoleType` object.
type AWSAssumeRoleType string

// List of AWSAssumeRoleType.
const (
	AWSASSUMEROLETYPE_AWSASSUMEROLE AWSAssumeRoleType = "AWSAssumeRole"
)

var allowedAWSAssumeRoleTypeEnumValues = []AWSAssumeRoleType{
	AWSASSUMEROLETYPE_AWSASSUMEROLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AWSAssumeRoleType) GetAllowedValues() []AWSAssumeRoleType {
	return allowedAWSAssumeRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AWSAssumeRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AWSAssumeRoleType(value)
	return nil
}

// NewAWSAssumeRoleTypeFromValue returns a pointer to a valid AWSAssumeRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAWSAssumeRoleTypeFromValue(v string) (*AWSAssumeRoleType, error) {
	ev := AWSAssumeRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AWSAssumeRoleType: valid values are %v", v, allowedAWSAssumeRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AWSAssumeRoleType) IsValid() bool {
	for _, existing := range allowedAWSAssumeRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSAssumeRoleType value.
func (v AWSAssumeRoleType) Ptr() *AWSAssumeRoleType {
	return &v
}
