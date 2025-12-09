// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType Rulesets resource type.
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType string

// List of GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType.
const (
	GETMULTIPLERULESETSRESPONSEDATAATTRIBUTESRULESETSITEMSDATATYPE_RULESETS GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType = "rulesets"
)

var allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeEnumValues = []GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType{
	GETMULTIPLERULESETSRESPONSEDATAATTRIBUTESRULESETSITEMSDATATYPE_RULESETS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType) GetAllowedValues() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType {
	return allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType(value)
	return nil
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeFromValue returns a pointer to a valid GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeFromValue(v string) (*GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType, error) {
	ev := GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType: valid values are %v", v, allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType) IsValid() bool {
	for _, existing := range allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType value.
func (v GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType) Ptr() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsDataType {
	return &v
}
