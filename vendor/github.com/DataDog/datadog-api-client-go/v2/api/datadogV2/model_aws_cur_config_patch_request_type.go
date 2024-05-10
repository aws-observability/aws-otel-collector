// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigPatchRequestType Type of AWS CUR config Patch Request.
type AwsCURConfigPatchRequestType string

// List of AwsCURConfigPatchRequestType.
const (
	AWSCURCONFIGPATCHREQUESTTYPE_AWS_CUR_CONFIG_PATCH_REQUEST AwsCURConfigPatchRequestType = "aws_cur_config_patch_request"
)

var allowedAwsCURConfigPatchRequestTypeEnumValues = []AwsCURConfigPatchRequestType{
	AWSCURCONFIGPATCHREQUESTTYPE_AWS_CUR_CONFIG_PATCH_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AwsCURConfigPatchRequestType) GetAllowedValues() []AwsCURConfigPatchRequestType {
	return allowedAwsCURConfigPatchRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AwsCURConfigPatchRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AwsCURConfigPatchRequestType(value)
	return nil
}

// NewAwsCURConfigPatchRequestTypeFromValue returns a pointer to a valid AwsCURConfigPatchRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAwsCURConfigPatchRequestTypeFromValue(v string) (*AwsCURConfigPatchRequestType, error) {
	ev := AwsCURConfigPatchRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsCURConfigPatchRequestType: valid values are %v", v, allowedAwsCURConfigPatchRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AwsCURConfigPatchRequestType) IsValid() bool {
	for _, existing := range allowedAwsCURConfigPatchRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsCURConfigPatchRequestType value.
func (v AwsCURConfigPatchRequestType) Ptr() *AwsCURConfigPatchRequestType {
	return &v
}
