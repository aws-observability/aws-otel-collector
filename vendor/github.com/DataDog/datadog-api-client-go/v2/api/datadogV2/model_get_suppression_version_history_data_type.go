// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetSuppressionVersionHistoryDataType Type of data.
type GetSuppressionVersionHistoryDataType string

// List of GetSuppressionVersionHistoryDataType.
const (
	GETSUPPRESSIONVERSIONHISTORYDATATYPE_SUPPRESSIONVERSIONHISTORY GetSuppressionVersionHistoryDataType = "suppression_version_history"
)

var allowedGetSuppressionVersionHistoryDataTypeEnumValues = []GetSuppressionVersionHistoryDataType{
	GETSUPPRESSIONVERSIONHISTORYDATATYPE_SUPPRESSIONVERSIONHISTORY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetSuppressionVersionHistoryDataType) GetAllowedValues() []GetSuppressionVersionHistoryDataType {
	return allowedGetSuppressionVersionHistoryDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetSuppressionVersionHistoryDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetSuppressionVersionHistoryDataType(value)
	return nil
}

// NewGetSuppressionVersionHistoryDataTypeFromValue returns a pointer to a valid GetSuppressionVersionHistoryDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetSuppressionVersionHistoryDataTypeFromValue(v string) (*GetSuppressionVersionHistoryDataType, error) {
	ev := GetSuppressionVersionHistoryDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetSuppressionVersionHistoryDataType: valid values are %v", v, allowedGetSuppressionVersionHistoryDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetSuppressionVersionHistoryDataType) IsValid() bool {
	for _, existing := range allowedGetSuppressionVersionHistoryDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetSuppressionVersionHistoryDataType value.
func (v GetSuppressionVersionHistoryDataType) Ptr() *GetSuppressionVersionHistoryDataType {
	return &v
}
