// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCurConfigResponseDataType AWS CUR config resource type.
type AwsCurConfigResponseDataType string

// List of AwsCurConfigResponseDataType.
const (
	AWSCURCONFIGRESPONSEDATATYPE_AWS_CUR_CONFIG AwsCurConfigResponseDataType = "aws_cur_config"
)

var allowedAwsCurConfigResponseDataTypeEnumValues = []AwsCurConfigResponseDataType{
	AWSCURCONFIGRESPONSEDATATYPE_AWS_CUR_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsCurConfigResponseDataType) GetAllowedValues() []AwsCurConfigResponseDataType {
	return allowedAwsCurConfigResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsCurConfigResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsCurConfigResponseDataType(value)
	return nil
}

// NewAwsCurConfigResponseDataTypeFromValue returns a pointer to a valid AwsCurConfigResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsCurConfigResponseDataTypeFromValue(v string) (*AwsCurConfigResponseDataType, error) {
	ev := AwsCurConfigResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsCurConfigResponseDataType: valid values are %v", v, allowedAwsCurConfigResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsCurConfigResponseDataType) IsValid() bool {
	for _, existing := range allowedAwsCurConfigResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsCurConfigResponseDataType value.
func (v AwsCurConfigResponseDataType) Ptr() *AwsCurConfigResponseDataType {
	return &v
}
