// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IPAllowlistType IP allowlist type.
type IPAllowlistType string

// List of IPAllowlistType.
const (
	IPALLOWLISTTYPE_IP_ALLOWLIST IPAllowlistType = "ip_allowlist"
)

var allowedIPAllowlistTypeEnumValues = []IPAllowlistType{
	IPALLOWLISTTYPE_IP_ALLOWLIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IPAllowlistType) GetAllowedValues() []IPAllowlistType {
	return allowedIPAllowlistTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IPAllowlistType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IPAllowlistType(value)
	return nil
}

// NewIPAllowlistTypeFromValue returns a pointer to a valid IPAllowlistType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIPAllowlistTypeFromValue(v string) (*IPAllowlistType, error) {
	ev := IPAllowlistType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IPAllowlistType: valid values are %v", v, allowedIPAllowlistTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IPAllowlistType) IsValid() bool {
	for _, existing := range allowedIPAllowlistTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAllowlistType value.
func (v IPAllowlistType) Ptr() *IPAllowlistType {
	return &v
}
