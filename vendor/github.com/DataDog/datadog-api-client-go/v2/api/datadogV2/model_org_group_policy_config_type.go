// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyConfigType Org group policy configs resource type.
type OrgGroupPolicyConfigType string

// List of OrgGroupPolicyConfigType.
const (
	ORGGROUPPOLICYCONFIGTYPE_ORG_GROUP_POLICY_CONFIGS OrgGroupPolicyConfigType = "org_group_policy_configs"
)

var allowedOrgGroupPolicyConfigTypeEnumValues = []OrgGroupPolicyConfigType{
	ORGGROUPPOLICYCONFIGTYPE_ORG_GROUP_POLICY_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyConfigType) GetAllowedValues() []OrgGroupPolicyConfigType {
	return allowedOrgGroupPolicyConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyConfigType(value)
	return nil
}

// NewOrgGroupPolicyConfigTypeFromValue returns a pointer to a valid OrgGroupPolicyConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyConfigTypeFromValue(v string) (*OrgGroupPolicyConfigType, error) {
	ev := OrgGroupPolicyConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyConfigType: valid values are %v", v, allowedOrgGroupPolicyConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyConfigType) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyConfigType value.
func (v OrgGroupPolicyConfigType) Ptr() *OrgGroupPolicyConfigType {
	return &v
}
