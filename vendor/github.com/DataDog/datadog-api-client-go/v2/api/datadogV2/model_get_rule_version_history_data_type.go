// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetRuleVersionHistoryDataType Type of data.
type GetRuleVersionHistoryDataType string

// List of GetRuleVersionHistoryDataType.
const (
	GETRULEVERSIONHISTORYDATATYPE_GETRULEVERSIONHISTORYRESPONSE GetRuleVersionHistoryDataType = "GetRuleVersionHistoryResponse"
)

var allowedGetRuleVersionHistoryDataTypeEnumValues = []GetRuleVersionHistoryDataType{
	GETRULEVERSIONHISTORYDATATYPE_GETRULEVERSIONHISTORYRESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetRuleVersionHistoryDataType) GetAllowedValues() []GetRuleVersionHistoryDataType {
	return allowedGetRuleVersionHistoryDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetRuleVersionHistoryDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetRuleVersionHistoryDataType(value)
	return nil
}

// NewGetRuleVersionHistoryDataTypeFromValue returns a pointer to a valid GetRuleVersionHistoryDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetRuleVersionHistoryDataTypeFromValue(v string) (*GetRuleVersionHistoryDataType, error) {
	ev := GetRuleVersionHistoryDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetRuleVersionHistoryDataType: valid values are %v", v, allowedGetRuleVersionHistoryDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetRuleVersionHistoryDataType) IsValid() bool {
	for _, existing := range allowedGetRuleVersionHistoryDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetRuleVersionHistoryDataType value.
func (v GetRuleVersionHistoryDataType) Ptr() *GetRuleVersionHistoryDataType {
	return &v
}
