// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesValidateQueryRequestDataType Validate query resource type.
type RulesValidateQueryRequestDataType string

// List of RulesValidateQueryRequestDataType.
const (
	RULESVALIDATEQUERYREQUESTDATATYPE_VALIDATE_QUERY RulesValidateQueryRequestDataType = "validate_query"
)

var allowedRulesValidateQueryRequestDataTypeEnumValues = []RulesValidateQueryRequestDataType{
	RULESVALIDATEQUERYREQUESTDATATYPE_VALIDATE_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RulesValidateQueryRequestDataType) GetAllowedValues() []RulesValidateQueryRequestDataType {
	return allowedRulesValidateQueryRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RulesValidateQueryRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RulesValidateQueryRequestDataType(value)
	return nil
}

// NewRulesValidateQueryRequestDataTypeFromValue returns a pointer to a valid RulesValidateQueryRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRulesValidateQueryRequestDataTypeFromValue(v string) (*RulesValidateQueryRequestDataType, error) {
	ev := RulesValidateQueryRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RulesValidateQueryRequestDataType: valid values are %v", v, allowedRulesValidateQueryRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RulesValidateQueryRequestDataType) IsValid() bool {
	for _, existing := range allowedRulesValidateQueryRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RulesValidateQueryRequestDataType value.
func (v RulesValidateQueryRequestDataType) Ptr() *RulesValidateQueryRequestDataType {
	return &v
}
