// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesValidateQueryResponseDataType Validate response resource type.
type RulesValidateQueryResponseDataType string

// List of RulesValidateQueryResponseDataType.
const (
	RULESVALIDATEQUERYRESPONSEDATATYPE_VALIDATE_RESPONSE RulesValidateQueryResponseDataType = "validate_response"
)

var allowedRulesValidateQueryResponseDataTypeEnumValues = []RulesValidateQueryResponseDataType{
	RULESVALIDATEQUERYRESPONSEDATATYPE_VALIDATE_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RulesValidateQueryResponseDataType) GetAllowedValues() []RulesValidateQueryResponseDataType {
	return allowedRulesValidateQueryResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RulesValidateQueryResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RulesValidateQueryResponseDataType(value)
	return nil
}

// NewRulesValidateQueryResponseDataTypeFromValue returns a pointer to a valid RulesValidateQueryResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRulesValidateQueryResponseDataTypeFromValue(v string) (*RulesValidateQueryResponseDataType, error) {
	ev := RulesValidateQueryResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RulesValidateQueryResponseDataType: valid values are %v", v, allowedRulesValidateQueryResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RulesValidateQueryResponseDataType) IsValid() bool {
	for _, existing := range allowedRulesValidateQueryResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RulesValidateQueryResponseDataType value.
func (v RulesValidateQueryResponseDataType) Ptr() *RulesValidateQueryResponseDataType {
	return &v
}
