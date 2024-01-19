// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IPAllowlistEntryType IP allowlist Entry type.
type IPAllowlistEntryType string

// List of IPAllowlistEntryType.
const (
	IPALLOWLISTENTRYTYPE_IP_ALLOWLIST_ENTRY IPAllowlistEntryType = "ip_allowlist_entry"
)

var allowedIPAllowlistEntryTypeEnumValues = []IPAllowlistEntryType{
	IPALLOWLISTENTRYTYPE_IP_ALLOWLIST_ENTRY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IPAllowlistEntryType) GetAllowedValues() []IPAllowlistEntryType {
	return allowedIPAllowlistEntryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IPAllowlistEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IPAllowlistEntryType(value)
	return nil
}

// NewIPAllowlistEntryTypeFromValue returns a pointer to a valid IPAllowlistEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIPAllowlistEntryTypeFromValue(v string) (*IPAllowlistEntryType, error) {
	ev := IPAllowlistEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IPAllowlistEntryType: valid values are %v", v, allowedIPAllowlistEntryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IPAllowlistEntryType) IsValid() bool {
	for _, existing := range allowedIPAllowlistEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAllowlistEntryType value.
func (v IPAllowlistEntryType) Ptr() *IPAllowlistEntryType {
	return &v
}
