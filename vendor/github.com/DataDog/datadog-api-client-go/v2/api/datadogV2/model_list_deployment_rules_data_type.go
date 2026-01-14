// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListDeploymentRulesDataType List deployment rule resource type.
type ListDeploymentRulesDataType string

// List of ListDeploymentRulesDataType.
const (
	LISTDEPLOYMENTRULESDATATYPE_LIST_DEPLOYMENT_RULES ListDeploymentRulesDataType = "list_deployment_rules"
)

var allowedListDeploymentRulesDataTypeEnumValues = []ListDeploymentRulesDataType{
	LISTDEPLOYMENTRULESDATATYPE_LIST_DEPLOYMENT_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListDeploymentRulesDataType) GetAllowedValues() []ListDeploymentRulesDataType {
	return allowedListDeploymentRulesDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListDeploymentRulesDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListDeploymentRulesDataType(value)
	return nil
}

// NewListDeploymentRulesDataTypeFromValue returns a pointer to a valid ListDeploymentRulesDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListDeploymentRulesDataTypeFromValue(v string) (*ListDeploymentRulesDataType, error) {
	ev := ListDeploymentRulesDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListDeploymentRulesDataType: valid values are %v", v, allowedListDeploymentRulesDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListDeploymentRulesDataType) IsValid() bool {
	for _, existing := range allowedListDeploymentRulesDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListDeploymentRulesDataType value.
func (v ListDeploymentRulesDataType) Ptr() *ListDeploymentRulesDataType {
	return &v
}
