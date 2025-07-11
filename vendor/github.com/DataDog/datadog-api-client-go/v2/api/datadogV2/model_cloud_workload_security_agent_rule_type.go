// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleType The type of the resource, must always be `agent_rule`
type CloudWorkloadSecurityAgentRuleType string

// List of CloudWorkloadSecurityAgentRuleType.
const (
	CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE CloudWorkloadSecurityAgentRuleType = "agent_rule"
)

var allowedCloudWorkloadSecurityAgentRuleTypeEnumValues = []CloudWorkloadSecurityAgentRuleType{
	CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudWorkloadSecurityAgentRuleType) GetAllowedValues() []CloudWorkloadSecurityAgentRuleType {
	return allowedCloudWorkloadSecurityAgentRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudWorkloadSecurityAgentRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudWorkloadSecurityAgentRuleType(value)
	return nil
}

// NewCloudWorkloadSecurityAgentRuleTypeFromValue returns a pointer to a valid CloudWorkloadSecurityAgentRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudWorkloadSecurityAgentRuleTypeFromValue(v string) (*CloudWorkloadSecurityAgentRuleType, error) {
	ev := CloudWorkloadSecurityAgentRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudWorkloadSecurityAgentRuleType: valid values are %v", v, allowedCloudWorkloadSecurityAgentRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudWorkloadSecurityAgentRuleType) IsValid() bool {
	for _, existing := range allowedCloudWorkloadSecurityAgentRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudWorkloadSecurityAgentRuleType value.
func (v CloudWorkloadSecurityAgentRuleType) Ptr() *CloudWorkloadSecurityAgentRuleType {
	return &v
}
