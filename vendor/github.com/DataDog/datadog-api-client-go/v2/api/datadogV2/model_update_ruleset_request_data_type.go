// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataType Update ruleset resource type.
type UpdateRulesetRequestDataType string

// List of UpdateRulesetRequestDataType.
const (
	UPDATERULESETREQUESTDATATYPE_UPDATE_RULESET UpdateRulesetRequestDataType = "update_ruleset"
)

var allowedUpdateRulesetRequestDataTypeEnumValues = []UpdateRulesetRequestDataType{
	UPDATERULESETREQUESTDATATYPE_UPDATE_RULESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateRulesetRequestDataType) GetAllowedValues() []UpdateRulesetRequestDataType {
	return allowedUpdateRulesetRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateRulesetRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateRulesetRequestDataType(value)
	return nil
}

// NewUpdateRulesetRequestDataTypeFromValue returns a pointer to a valid UpdateRulesetRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateRulesetRequestDataTypeFromValue(v string) (*UpdateRulesetRequestDataType, error) {
	ev := UpdateRulesetRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateRulesetRequestDataType: valid values are %v", v, allowedUpdateRulesetRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateRulesetRequestDataType) IsValid() bool {
	for _, existing := range allowedUpdateRulesetRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateRulesetRequestDataType value.
func (v UpdateRulesetRequestDataType) Ptr() *UpdateRulesetRequestDataType {
	return &v
}
