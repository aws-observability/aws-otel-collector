// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RevertCustomRuleRevisionDataType Request type
type RevertCustomRuleRevisionDataType string

// List of RevertCustomRuleRevisionDataType.
const (
	REVERTCUSTOMRULEREVISIONDATATYPE_REVERT_CUSTOM_RULE_REVISION_REQUEST RevertCustomRuleRevisionDataType = "revert_custom_rule_revision_request"
)

var allowedRevertCustomRuleRevisionDataTypeEnumValues = []RevertCustomRuleRevisionDataType{
	REVERTCUSTOMRULEREVISIONDATATYPE_REVERT_CUSTOM_RULE_REVISION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RevertCustomRuleRevisionDataType) GetAllowedValues() []RevertCustomRuleRevisionDataType {
	return allowedRevertCustomRuleRevisionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RevertCustomRuleRevisionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RevertCustomRuleRevisionDataType(value)
	return nil
}

// NewRevertCustomRuleRevisionDataTypeFromValue returns a pointer to a valid RevertCustomRuleRevisionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRevertCustomRuleRevisionDataTypeFromValue(v string) (*RevertCustomRuleRevisionDataType, error) {
	ev := RevertCustomRuleRevisionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RevertCustomRuleRevisionDataType: valid values are %v", v, allowedRevertCustomRuleRevisionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RevertCustomRuleRevisionDataType) IsValid() bool {
	for _, existing := range allowedRevertCustomRuleRevisionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RevertCustomRuleRevisionDataType value.
func (v RevertCustomRuleRevisionDataType) Ptr() *RevertCustomRuleRevisionDataType {
	return &v
}
