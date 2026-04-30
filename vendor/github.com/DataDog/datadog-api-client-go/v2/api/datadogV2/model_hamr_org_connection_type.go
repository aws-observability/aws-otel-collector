// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HamrOrgConnectionType Type of the HAMR organization connection resource.
type HamrOrgConnectionType string

// List of HamrOrgConnectionType.
const (
	HAMRORGCONNECTIONTYPE_HAMR_ORG_CONNECTIONS HamrOrgConnectionType = "hamr_org_connections"
)

var allowedHamrOrgConnectionTypeEnumValues = []HamrOrgConnectionType{
	HAMRORGCONNECTIONTYPE_HAMR_ORG_CONNECTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HamrOrgConnectionType) GetAllowedValues() []HamrOrgConnectionType {
	return allowedHamrOrgConnectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HamrOrgConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HamrOrgConnectionType(value)
	return nil
}

// NewHamrOrgConnectionTypeFromValue returns a pointer to a valid HamrOrgConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHamrOrgConnectionTypeFromValue(v string) (*HamrOrgConnectionType, error) {
	ev := HamrOrgConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HamrOrgConnectionType: valid values are %v", v, allowedHamrOrgConnectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HamrOrgConnectionType) IsValid() bool {
	for _, existing := range allowedHamrOrgConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HamrOrgConnectionType value.
func (v HamrOrgConnectionType) Ptr() *HamrOrgConnectionType {
	return &v
}
