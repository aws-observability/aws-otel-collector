// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReorderRulesetResourceDataType Ruleset resource type.
type ReorderRulesetResourceDataType string

// List of ReorderRulesetResourceDataType.
const (
	REORDERRULESETRESOURCEDATATYPE_RULESET ReorderRulesetResourceDataType = "ruleset"
)

var allowedReorderRulesetResourceDataTypeEnumValues = []ReorderRulesetResourceDataType{
	REORDERRULESETRESOURCEDATATYPE_RULESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReorderRulesetResourceDataType) GetAllowedValues() []ReorderRulesetResourceDataType {
	return allowedReorderRulesetResourceDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReorderRulesetResourceDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReorderRulesetResourceDataType(value)
	return nil
}

// NewReorderRulesetResourceDataTypeFromValue returns a pointer to a valid ReorderRulesetResourceDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReorderRulesetResourceDataTypeFromValue(v string) (*ReorderRulesetResourceDataType, error) {
	ev := ReorderRulesetResourceDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReorderRulesetResourceDataType: valid values are %v", v, allowedReorderRulesetResourceDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReorderRulesetResourceDataType) IsValid() bool {
	for _, existing := range allowedReorderRulesetResourceDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReorderRulesetResourceDataType value.
func (v ReorderRulesetResourceDataType) Ptr() *ReorderRulesetResourceDataType {
	return &v
}
