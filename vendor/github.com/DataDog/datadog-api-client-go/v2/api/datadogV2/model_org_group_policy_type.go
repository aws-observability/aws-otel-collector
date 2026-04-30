// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyType Org group policies resource type.
type OrgGroupPolicyType string

// List of OrgGroupPolicyType.
const (
	ORGGROUPPOLICYTYPE_ORG_GROUP_POLICIES OrgGroupPolicyType = "org_group_policies"
)

var allowedOrgGroupPolicyTypeEnumValues = []OrgGroupPolicyType{
	ORGGROUPPOLICYTYPE_ORG_GROUP_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyType) GetAllowedValues() []OrgGroupPolicyType {
	return allowedOrgGroupPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyType(value)
	return nil
}

// NewOrgGroupPolicyTypeFromValue returns a pointer to a valid OrgGroupPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyTypeFromValue(v string) (*OrgGroupPolicyType, error) {
	ev := OrgGroupPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyType: valid values are %v", v, allowedOrgGroupPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyType) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyType value.
func (v OrgGroupPolicyType) Ptr() *OrgGroupPolicyType {
	return &v
}
