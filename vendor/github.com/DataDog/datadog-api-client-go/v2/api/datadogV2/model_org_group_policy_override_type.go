// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideType Org group policy overrides resource type.
type OrgGroupPolicyOverrideType string

// List of OrgGroupPolicyOverrideType.
const (
	ORGGROUPPOLICYOVERRIDETYPE_ORG_GROUP_POLICY_OVERRIDES OrgGroupPolicyOverrideType = "org_group_policy_overrides"
)

var allowedOrgGroupPolicyOverrideTypeEnumValues = []OrgGroupPolicyOverrideType{
	ORGGROUPPOLICYOVERRIDETYPE_ORG_GROUP_POLICY_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyOverrideType) GetAllowedValues() []OrgGroupPolicyOverrideType {
	return allowedOrgGroupPolicyOverrideTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyOverrideType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyOverrideType(value)
	return nil
}

// NewOrgGroupPolicyOverrideTypeFromValue returns a pointer to a valid OrgGroupPolicyOverrideType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyOverrideTypeFromValue(v string) (*OrgGroupPolicyOverrideType, error) {
	ev := OrgGroupPolicyOverrideType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyOverrideType: valid values are %v", v, allowedOrgGroupPolicyOverrideTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyOverrideType) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyOverrideTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyOverrideType value.
func (v OrgGroupPolicyOverrideType) Ptr() *OrgGroupPolicyOverrideType {
	return &v
}
