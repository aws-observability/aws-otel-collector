// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyPolicyType The type of the policy. Only `org_config` is supported, indicating a policy backed by an organization configuration setting.
type OrgGroupPolicyPolicyType string

// List of OrgGroupPolicyPolicyType.
const (
	ORGGROUPPOLICYPOLICYTYPE_ORG_CONFIG OrgGroupPolicyPolicyType = "org_config"
)

var allowedOrgGroupPolicyPolicyTypeEnumValues = []OrgGroupPolicyPolicyType{
	ORGGROUPPOLICYPOLICYTYPE_ORG_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyPolicyType) GetAllowedValues() []OrgGroupPolicyPolicyType {
	return allowedOrgGroupPolicyPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyPolicyType(value)
	return nil
}

// NewOrgGroupPolicyPolicyTypeFromValue returns a pointer to a valid OrgGroupPolicyPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyPolicyTypeFromValue(v string) (*OrgGroupPolicyPolicyType, error) {
	ev := OrgGroupPolicyPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyPolicyType: valid values are %v", v, allowedOrgGroupPolicyPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyPolicyType) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyPolicyType value.
func (v OrgGroupPolicyPolicyType) Ptr() *OrgGroupPolicyPolicyType {
	return &v
}
