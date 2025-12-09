// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateRulesetRequestDataType Create ruleset resource type.
type CreateRulesetRequestDataType string

// List of CreateRulesetRequestDataType.
const (
	CREATERULESETREQUESTDATATYPE_CREATE_RULESET CreateRulesetRequestDataType = "create_ruleset"
)

var allowedCreateRulesetRequestDataTypeEnumValues = []CreateRulesetRequestDataType{
	CREATERULESETREQUESTDATATYPE_CREATE_RULESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateRulesetRequestDataType) GetAllowedValues() []CreateRulesetRequestDataType {
	return allowedCreateRulesetRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateRulesetRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateRulesetRequestDataType(value)
	return nil
}

// NewCreateRulesetRequestDataTypeFromValue returns a pointer to a valid CreateRulesetRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateRulesetRequestDataTypeFromValue(v string) (*CreateRulesetRequestDataType, error) {
	ev := CreateRulesetRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateRulesetRequestDataType: valid values are %v", v, allowedCreateRulesetRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateRulesetRequestDataType) IsValid() bool {
	for _, existing := range allowedCreateRulesetRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateRulesetRequestDataType value.
func (v CreateRulesetRequestDataType) Ptr() *CreateRulesetRequestDataType {
	return &v
}
