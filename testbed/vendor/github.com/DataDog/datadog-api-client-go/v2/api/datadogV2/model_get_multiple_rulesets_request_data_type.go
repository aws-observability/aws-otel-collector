// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsRequestDataType Get multiple rulesets request resource type.
type GetMultipleRulesetsRequestDataType string

// List of GetMultipleRulesetsRequestDataType.
const (
	GETMULTIPLERULESETSREQUESTDATATYPE_GET_MULTIPLE_RULESETS_REQUEST GetMultipleRulesetsRequestDataType = "get_multiple_rulesets_request"
)

var allowedGetMultipleRulesetsRequestDataTypeEnumValues = []GetMultipleRulesetsRequestDataType{
	GETMULTIPLERULESETSREQUESTDATATYPE_GET_MULTIPLE_RULESETS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetMultipleRulesetsRequestDataType) GetAllowedValues() []GetMultipleRulesetsRequestDataType {
	return allowedGetMultipleRulesetsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetMultipleRulesetsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetMultipleRulesetsRequestDataType(value)
	return nil
}

// NewGetMultipleRulesetsRequestDataTypeFromValue returns a pointer to a valid GetMultipleRulesetsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetMultipleRulesetsRequestDataTypeFromValue(v string) (*GetMultipleRulesetsRequestDataType, error) {
	ev := GetMultipleRulesetsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetMultipleRulesetsRequestDataType: valid values are %v", v, allowedGetMultipleRulesetsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetMultipleRulesetsRequestDataType) IsValid() bool {
	for _, existing := range allowedGetMultipleRulesetsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetMultipleRulesetsRequestDataType value.
func (v GetMultipleRulesetsRequestDataType) Ptr() *GetMultipleRulesetsRequestDataType {
	return &v
}
