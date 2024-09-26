// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConfigType Data type of an Org Config.
type OrgConfigType string

// List of OrgConfigType.
const (
	ORGCONFIGTYPE_ORG_CONFIGS OrgConfigType = "org_configs"
)

var allowedOrgConfigTypeEnumValues = []OrgConfigType{
	ORGCONFIGTYPE_ORG_CONFIGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgConfigType) GetAllowedValues() []OrgConfigType {
	return allowedOrgConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgConfigType(value)
	return nil
}

// NewOrgConfigTypeFromValue returns a pointer to a valid OrgConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgConfigTypeFromValue(v string) (*OrgConfigType, error) {
	ev := OrgConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgConfigType: valid values are %v", v, allowedOrgConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgConfigType) IsValid() bool {
	for _, existing := range allowedOrgConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgConfigType value.
func (v OrgConfigType) Ptr() *OrgConfigType {
	return &v
}
