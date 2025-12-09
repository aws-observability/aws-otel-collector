// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReorderRuleResourceDataType Arbitrary rule resource type.
type ReorderRuleResourceDataType string

// List of ReorderRuleResourceDataType.
const (
	REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE ReorderRuleResourceDataType = "arbitrary_rule"
)

var allowedReorderRuleResourceDataTypeEnumValues = []ReorderRuleResourceDataType{
	REORDERRULERESOURCEDATATYPE_ARBITRARY_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReorderRuleResourceDataType) GetAllowedValues() []ReorderRuleResourceDataType {
	return allowedReorderRuleResourceDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReorderRuleResourceDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReorderRuleResourceDataType(value)
	return nil
}

// NewReorderRuleResourceDataTypeFromValue returns a pointer to a valid ReorderRuleResourceDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReorderRuleResourceDataTypeFromValue(v string) (*ReorderRuleResourceDataType, error) {
	ev := ReorderRuleResourceDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReorderRuleResourceDataType: valid values are %v", v, allowedReorderRuleResourceDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReorderRuleResourceDataType) IsValid() bool {
	for _, existing := range allowedReorderRuleResourceDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReorderRuleResourceDataType value.
func (v ReorderRuleResourceDataType) Ptr() *ReorderRuleResourceDataType {
	return &v
}
