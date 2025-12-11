// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionType Org connection type.
type OrgConnectionType string

// List of OrgConnectionType.
const (
	ORGCONNECTIONTYPE_ORG_CONNECTION OrgConnectionType = "org_connection"
)

var allowedOrgConnectionTypeEnumValues = []OrgConnectionType{
	ORGCONNECTIONTYPE_ORG_CONNECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgConnectionType) GetAllowedValues() []OrgConnectionType {
	return allowedOrgConnectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgConnectionType(value)
	return nil
}

// NewOrgConnectionTypeFromValue returns a pointer to a valid OrgConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgConnectionTypeFromValue(v string) (*OrgConnectionType, error) {
	ev := OrgConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgConnectionType: valid values are %v", v, allowedOrgConnectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgConnectionType) IsValid() bool {
	for _, existing := range allowedOrgConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgConnectionType value.
func (v OrgConnectionType) Ptr() *OrgConnectionType {
	return &v
}
