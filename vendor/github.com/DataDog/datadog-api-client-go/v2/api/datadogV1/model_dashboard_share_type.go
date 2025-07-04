// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardShareType Type of sharing access (either open to anyone who has the public URL or invite-only).
type DashboardShareType string

// List of DashboardShareType.
const (
	DASHBOARDSHARETYPE_OPEN   DashboardShareType = "open"
	DASHBOARDSHARETYPE_INVITE DashboardShareType = "invite"
	DASHBOARDSHARETYPE_EMBED  DashboardShareType = "embed"
)

var allowedDashboardShareTypeEnumValues = []DashboardShareType{
	DASHBOARDSHARETYPE_OPEN,
	DASHBOARDSHARETYPE_INVITE,
	DASHBOARDSHARETYPE_EMBED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardShareType) GetAllowedValues() []DashboardShareType {
	return allowedDashboardShareTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardShareType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardShareType(value)
	return nil
}

// NewDashboardShareTypeFromValue returns a pointer to a valid DashboardShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardShareTypeFromValue(v string) (*DashboardShareType, error) {
	ev := DashboardShareType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardShareType: valid values are %v", v, allowedDashboardShareTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardShareType) IsValid() bool {
	for _, existing := range allowedDashboardShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardShareType value.
func (v DashboardShareType) Ptr() *DashboardShareType {
	return &v
}

// NullableDashboardShareType handles when a null is used for DashboardShareType.
type NullableDashboardShareType struct {
	value *DashboardShareType
	isSet bool
}

// Get returns the associated value.
func (v NullableDashboardShareType) Get() *DashboardShareType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDashboardShareType) Set(val *DashboardShareType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDashboardShareType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDashboardShareType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardShareType initializes the struct as if Set has been called.
func NewNullableDashboardShareType(val *DashboardShareType) *NullableDashboardShareType {
	return &NullableDashboardShareType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDashboardShareType) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDashboardShareType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
