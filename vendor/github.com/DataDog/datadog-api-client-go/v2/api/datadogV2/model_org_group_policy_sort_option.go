// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicySortOption Field to sort policies by.
type OrgGroupPolicySortOption string

// List of OrgGroupPolicySortOption.
const (
	ORGGROUPPOLICYSORTOPTION_ID         OrgGroupPolicySortOption = "id"
	ORGGROUPPOLICYSORTOPTION_MINUS_ID   OrgGroupPolicySortOption = "-id"
	ORGGROUPPOLICYSORTOPTION_NAME       OrgGroupPolicySortOption = "name"
	ORGGROUPPOLICYSORTOPTION_MINUS_NAME OrgGroupPolicySortOption = "-name"
)

var allowedOrgGroupPolicySortOptionEnumValues = []OrgGroupPolicySortOption{
	ORGGROUPPOLICYSORTOPTION_ID,
	ORGGROUPPOLICYSORTOPTION_MINUS_ID,
	ORGGROUPPOLICYSORTOPTION_NAME,
	ORGGROUPPOLICYSORTOPTION_MINUS_NAME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicySortOption) GetAllowedValues() []OrgGroupPolicySortOption {
	return allowedOrgGroupPolicySortOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicySortOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicySortOption(value)
	return nil
}

// NewOrgGroupPolicySortOptionFromValue returns a pointer to a valid OrgGroupPolicySortOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicySortOptionFromValue(v string) (*OrgGroupPolicySortOption, error) {
	ev := OrgGroupPolicySortOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicySortOption: valid values are %v", v, allowedOrgGroupPolicySortOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicySortOption) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicySortOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicySortOption value.
func (v OrgGroupPolicySortOption) Ptr() *OrgGroupPolicySortOption {
	return &v
}
