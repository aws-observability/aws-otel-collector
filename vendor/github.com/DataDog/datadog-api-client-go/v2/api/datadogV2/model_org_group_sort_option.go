// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupSortOption Field to sort org groups by.
type OrgGroupSortOption string

// List of OrgGroupSortOption.
const (
	ORGGROUPSORTOPTION_NAME       OrgGroupSortOption = "name"
	ORGGROUPSORTOPTION_MINUS_NAME OrgGroupSortOption = "-name"
	ORGGROUPSORTOPTION_UUID       OrgGroupSortOption = "uuid"
	ORGGROUPSORTOPTION_MINUS_UUID OrgGroupSortOption = "-uuid"
)

var allowedOrgGroupSortOptionEnumValues = []OrgGroupSortOption{
	ORGGROUPSORTOPTION_NAME,
	ORGGROUPSORTOPTION_MINUS_NAME,
	ORGGROUPSORTOPTION_UUID,
	ORGGROUPSORTOPTION_MINUS_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupSortOption) GetAllowedValues() []OrgGroupSortOption {
	return allowedOrgGroupSortOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupSortOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupSortOption(value)
	return nil
}

// NewOrgGroupSortOptionFromValue returns a pointer to a valid OrgGroupSortOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupSortOptionFromValue(v string) (*OrgGroupSortOption, error) {
	ev := OrgGroupSortOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupSortOption: valid values are %v", v, allowedOrgGroupSortOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupSortOption) IsValid() bool {
	for _, existing := range allowedOrgGroupSortOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupSortOption value.
func (v OrgGroupSortOption) Ptr() *OrgGroupSortOption {
	return &v
}
