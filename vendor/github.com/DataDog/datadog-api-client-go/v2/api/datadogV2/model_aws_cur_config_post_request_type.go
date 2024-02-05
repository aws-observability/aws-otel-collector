// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigPostRequestType Type of AWS CUR config Post Request.
type AwsCURConfigPostRequestType string

// List of AwsCURConfigPostRequestType.
const (
	AWSCURCONFIGPOSTREQUESTTYPE_AWS_CUR_CONFIG_POST_REQUEST AwsCURConfigPostRequestType = "aws_cur_config_post_request"
)

var allowedAwsCURConfigPostRequestTypeEnumValues = []AwsCURConfigPostRequestType{
	AWSCURCONFIGPOSTREQUESTTYPE_AWS_CUR_CONFIG_POST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsCURConfigPostRequestType) GetAllowedValues() []AwsCURConfigPostRequestType {
	return allowedAwsCURConfigPostRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsCURConfigPostRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsCURConfigPostRequestType(value)
	return nil
}

// NewAwsCURConfigPostRequestTypeFromValue returns a pointer to a valid AwsCURConfigPostRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsCURConfigPostRequestTypeFromValue(v string) (*AwsCURConfigPostRequestType, error) {
	ev := AwsCURConfigPostRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsCURConfigPostRequestType: valid values are %v", v, allowedAwsCURConfigPostRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsCURConfigPostRequestType) IsValid() bool {
	for _, existing := range allowedAwsCURConfigPostRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsCURConfigPostRequestType value.
func (v AwsCURConfigPostRequestType) Ptr() *AwsCURConfigPostRequestType {
	return &v
}
