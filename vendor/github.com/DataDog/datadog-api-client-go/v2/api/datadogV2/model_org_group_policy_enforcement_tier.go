// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyEnforcementTier The enforcement tier of the policy. `DEFAULT` means the policy is set but member orgs may mutate it. `ENFORCE` means the policy is strictly controlled and mutations are blocked for affected orgs. `DELEGATE` means each member org controls its own value.
type OrgGroupPolicyEnforcementTier string

// List of OrgGroupPolicyEnforcementTier.
const (
	ORGGROUPPOLICYENFORCEMENTTIER_DEFAULT  OrgGroupPolicyEnforcementTier = "DEFAULT"
	ORGGROUPPOLICYENFORCEMENTTIER_ENFORCE  OrgGroupPolicyEnforcementTier = "ENFORCE"
	ORGGROUPPOLICYENFORCEMENTTIER_DELEGATE OrgGroupPolicyEnforcementTier = "DELEGATE"
)

var allowedOrgGroupPolicyEnforcementTierEnumValues = []OrgGroupPolicyEnforcementTier{
	ORGGROUPPOLICYENFORCEMENTTIER_DEFAULT,
	ORGGROUPPOLICYENFORCEMENTTIER_ENFORCE,
	ORGGROUPPOLICYENFORCEMENTTIER_DELEGATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicyEnforcementTier) GetAllowedValues() []OrgGroupPolicyEnforcementTier {
	return allowedOrgGroupPolicyEnforcementTierEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicyEnforcementTier) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicyEnforcementTier(value)
	return nil
}

// NewOrgGroupPolicyEnforcementTierFromValue returns a pointer to a valid OrgGroupPolicyEnforcementTier
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicyEnforcementTierFromValue(v string) (*OrgGroupPolicyEnforcementTier, error) {
	ev := OrgGroupPolicyEnforcementTier(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicyEnforcementTier: valid values are %v", v, allowedOrgGroupPolicyEnforcementTierEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicyEnforcementTier) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicyEnforcementTierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicyEnforcementTier value.
func (v OrgGroupPolicyEnforcementTier) Ptr() *OrgGroupPolicyEnforcementTier {
	return &v
}
