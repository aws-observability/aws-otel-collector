// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyType The type of the resource, must always be `policy`
type CloudWorkloadSecurityAgentPolicyType string

// List of CloudWorkloadSecurityAgentPolicyType.
const (
	CLOUDWORKLOADSECURITYAGENTPOLICYTYPE_POLICY CloudWorkloadSecurityAgentPolicyType = "policy"
)

var allowedCloudWorkloadSecurityAgentPolicyTypeEnumValues = []CloudWorkloadSecurityAgentPolicyType{
	CLOUDWORKLOADSECURITYAGENTPOLICYTYPE_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudWorkloadSecurityAgentPolicyType) GetAllowedValues() []CloudWorkloadSecurityAgentPolicyType {
	return allowedCloudWorkloadSecurityAgentPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudWorkloadSecurityAgentPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudWorkloadSecurityAgentPolicyType(value)
	return nil
}

// NewCloudWorkloadSecurityAgentPolicyTypeFromValue returns a pointer to a valid CloudWorkloadSecurityAgentPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudWorkloadSecurityAgentPolicyTypeFromValue(v string) (*CloudWorkloadSecurityAgentPolicyType, error) {
	ev := CloudWorkloadSecurityAgentPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudWorkloadSecurityAgentPolicyType: valid values are %v", v, allowedCloudWorkloadSecurityAgentPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudWorkloadSecurityAgentPolicyType) IsValid() bool {
	for _, existing := range allowedCloudWorkloadSecurityAgentPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudWorkloadSecurityAgentPolicyType value.
func (v CloudWorkloadSecurityAgentPolicyType) Ptr() *CloudWorkloadSecurityAgentPolicyType {
	return &v
}
