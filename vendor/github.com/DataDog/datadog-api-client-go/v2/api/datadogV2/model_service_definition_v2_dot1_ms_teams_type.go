// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ServiceDefinitionV2Dot1MSTeamsType Contact type.
type ServiceDefinitionV2Dot1MSTeamsType string

// List of ServiceDefinitionV2Dot1MSTeamsType.
const (
	SERVICEDEFINITIONV2DOT1MSTEAMSTYPE_MICROSOFT_TEAMS ServiceDefinitionV2Dot1MSTeamsType = "microsoft-teams"
)

var allowedServiceDefinitionV2Dot1MSTeamsTypeEnumValues = []ServiceDefinitionV2Dot1MSTeamsType{
	SERVICEDEFINITIONV2DOT1MSTEAMSTYPE_MICROSOFT_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1MSTeamsType) GetAllowedValues() []ServiceDefinitionV2Dot1MSTeamsType {
	return allowedServiceDefinitionV2Dot1MSTeamsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1MSTeamsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1MSTeamsType(value)
	return nil
}

// NewServiceDefinitionV2Dot1MSTeamsTypeFromValue returns a pointer to a valid ServiceDefinitionV2Dot1MSTeamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1MSTeamsTypeFromValue(v string) (*ServiceDefinitionV2Dot1MSTeamsType, error) {
	ev := ServiceDefinitionV2Dot1MSTeamsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1MSTeamsType: valid values are %v", v, allowedServiceDefinitionV2Dot1MSTeamsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1MSTeamsType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1MSTeamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1MSTeamsType value.
func (v ServiceDefinitionV2Dot1MSTeamsType) Ptr() *ServiceDefinitionV2Dot1MSTeamsType {
	return &v
}

// NullableServiceDefinitionV2Dot1MSTeamsType handles when a null is used for ServiceDefinitionV2Dot1MSTeamsType.
type NullableServiceDefinitionV2Dot1MSTeamsType struct {
	value *ServiceDefinitionV2Dot1MSTeamsType
	isSet bool
}

// Get returns the associated value.
func (v NullableServiceDefinitionV2Dot1MSTeamsType) Get() *ServiceDefinitionV2Dot1MSTeamsType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableServiceDefinitionV2Dot1MSTeamsType) Set(val *ServiceDefinitionV2Dot1MSTeamsType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableServiceDefinitionV2Dot1MSTeamsType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableServiceDefinitionV2Dot1MSTeamsType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDefinitionV2Dot1MSTeamsType initializes the struct as if Set has been called.
func NewNullableServiceDefinitionV2Dot1MSTeamsType(val *ServiceDefinitionV2Dot1MSTeamsType) *NullableServiceDefinitionV2Dot1MSTeamsType {
	return &NullableServiceDefinitionV2Dot1MSTeamsType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableServiceDefinitionV2Dot1MSTeamsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableServiceDefinitionV2Dot1MSTeamsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
