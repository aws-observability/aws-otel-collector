// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataType Get multiple rulesets response resource type.
type GetMultipleRulesetsResponseDataType string

// List of GetMultipleRulesetsResponseDataType.
const (
	GETMULTIPLERULESETSRESPONSEDATATYPE_GET_MULTIPLE_RULESETS_RESPONSE GetMultipleRulesetsResponseDataType = "get_multiple_rulesets_response"
)

var allowedGetMultipleRulesetsResponseDataTypeEnumValues = []GetMultipleRulesetsResponseDataType{
	GETMULTIPLERULESETSRESPONSEDATATYPE_GET_MULTIPLE_RULESETS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetMultipleRulesetsResponseDataType) GetAllowedValues() []GetMultipleRulesetsResponseDataType {
	return allowedGetMultipleRulesetsResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetMultipleRulesetsResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetMultipleRulesetsResponseDataType(value)
	return nil
}

// NewGetMultipleRulesetsResponseDataTypeFromValue returns a pointer to a valid GetMultipleRulesetsResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetMultipleRulesetsResponseDataTypeFromValue(v string) (*GetMultipleRulesetsResponseDataType, error) {
	ev := GetMultipleRulesetsResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetMultipleRulesetsResponseDataType: valid values are %v", v, allowedGetMultipleRulesetsResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetMultipleRulesetsResponseDataType) IsValid() bool {
	for _, existing := range allowedGetMultipleRulesetsResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetMultipleRulesetsResponseDataType value.
func (v GetMultipleRulesetsResponseDataType) Ptr() *GetMultipleRulesetsResponseDataType {
	return &v
}
