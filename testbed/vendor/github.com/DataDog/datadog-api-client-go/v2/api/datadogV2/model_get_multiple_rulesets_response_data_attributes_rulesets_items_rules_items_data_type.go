// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType Rules resource type.
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType string

// List of GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType.
const (
	GETMULTIPLERULESETSRESPONSEDATAATTRIBUTESRULESETSITEMSRULESITEMSDATATYPE_RULES GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType = "rules"
)

var allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeEnumValues = []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType{
	GETMULTIPLERULESETSRESPONSEDATAATTRIBUTESRULESETSITEMSRULESITEMSDATATYPE_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType) GetAllowedValues() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType {
	return allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType(value)
	return nil
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeFromValue returns a pointer to a valid GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeFromValue(v string) (*GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType, error) {
	ev := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType: valid values are %v", v, allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType) IsValid() bool {
	for _, existing := range allowedGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType value.
func (v GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType) Ptr() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsDataType {
	return &v
}
