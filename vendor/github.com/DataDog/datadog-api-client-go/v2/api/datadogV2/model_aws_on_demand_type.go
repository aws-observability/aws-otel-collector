// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsOnDemandType The type of the on demand task. The value should always be `aws_resource`.
type AwsOnDemandType string

// List of AwsOnDemandType.
const (
	AWSONDEMANDTYPE_AWS_RESOURCE AwsOnDemandType = "aws_resource"
)

var allowedAwsOnDemandTypeEnumValues = []AwsOnDemandType{
	AWSONDEMANDTYPE_AWS_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsOnDemandType) GetAllowedValues() []AwsOnDemandType {
	return allowedAwsOnDemandTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsOnDemandType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsOnDemandType(value)
	return nil
}

// NewAwsOnDemandTypeFromValue returns a pointer to a valid AwsOnDemandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsOnDemandTypeFromValue(v string) (*AwsOnDemandType, error) {
	ev := AwsOnDemandType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsOnDemandType: valid values are %v", v, allowedAwsOnDemandTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsOnDemandType) IsValid() bool {
	for _, existing := range allowedAwsOnDemandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsOnDemandType value.
func (v AwsOnDemandType) Ptr() *AwsOnDemandType {
	return &v
}
