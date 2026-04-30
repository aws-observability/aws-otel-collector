// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyOverrideSortOption Field to sort overrides by.
type OrgGroupPolicyOverrideSortOption string

// List of OrgGroupPolicyOverrideSortOption.
const (
	ORGGROUPPOLICYOVERRIDESORTOPTION_ID             OrgGroupPolicyOverrideSortOption = "id"
	ORGGROUPPOLICYOVERRIDESORTOPTION_MINUS_ID       OrgGroupPolicyOverrideSortOption = "-id"
	ORGGROUPPOLICYOVERRIDESORTOPTION_ORG_UUID       OrgGroupPolicyOverrideSortOption = "org_uuid"
	ORGGROUPPOLICYOVERRIDESORTOPTION_MINUS_ORG_UUID OrgGroupPolicyOverrideSortOption = "-org_uuid"
)

var allowedOrgGroupPolicyOverrideSortOptionEnumValues = []OrgGroupPolicyOverrideSortOption{
	ORGGROUPPOLICYOVERRIDESORTOPTION_ID,
	ORGGROUPPOLICYOVERRIDESORTOPTION_MINUS_ID,
	ORGGROUPPOLICYOVERRIDESORTOPTION_ORG_UUID,
	ORGGROUPPOLICYOVERRIDESORTOPTION_MINUS_ORG_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyOverrideSortOption) GetAllowedValues() []OrgGroupPolicyOverrideSortOption {
	return allowedOrgGroupPolicyOverrideSortOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyOverrideSortOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyOverrideSortOption(value)
	return nil
}

// NewOrgGroupPolicyOverrideSortOptionFromValue returns a pointer to a valid OrgGroupPolicyOverrideSortOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyOverrideSortOptionFromValue(v string) (*OrgGroupPolicyOverrideSortOption, error) {
	ev := OrgGroupPolicyOverrideSortOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyOverrideSortOption: valid values are %v", v, allowedOrgGroupPolicyOverrideSortOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyOverrideSortOption) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyOverrideSortOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyOverrideSortOption value.
func (v OrgGroupPolicyOverrideSortOption) Ptr() *OrgGroupPolicyOverrideSortOption {
	return &v
}
