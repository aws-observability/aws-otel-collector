// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigType Type of AWS CUR config.
type AwsCURConfigType string

// List of AwsCURConfigType.
const (
	AWSCURCONFIGTYPE_AWS_CUR_CONFIG AwsCURConfigType = "aws_cur_config"
)

var allowedAwsCURConfigTypeEnumValues = []AwsCURConfigType{
	AWSCURCONFIGTYPE_AWS_CUR_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsCURConfigType) GetAllowedValues() []AwsCURConfigType {
	return allowedAwsCURConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsCURConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsCURConfigType(value)
	return nil
}

// NewAwsCURConfigTypeFromValue returns a pointer to a valid AwsCURConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsCURConfigTypeFromValue(v string) (*AwsCURConfigType, error) {
	ev := AwsCURConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsCURConfigType: valid values are %v", v, allowedAwsCURConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsCURConfigType) IsValid() bool {
	for _, existing := range allowedAwsCURConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsCURConfigType value.
func (v AwsCURConfigType) Ptr() *AwsCURConfigType {
	return &v
}
