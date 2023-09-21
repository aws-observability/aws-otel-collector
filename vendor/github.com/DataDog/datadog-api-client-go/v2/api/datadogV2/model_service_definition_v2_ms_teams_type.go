// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// ServiceDefinitionV2MSTeamsType Contact type.
type ServiceDefinitionV2MSTeamsType string

// List of ServiceDefinitionV2MSTeamsType.
const (
	SERVICEDEFINITIONV2MSTEAMSTYPE_MICROSOFT_TEAMS ServiceDefinitionV2MSTeamsType = "microsoft-teams"
)

var allowedServiceDefinitionV2MSTeamsTypeEnumValues = []ServiceDefinitionV2MSTeamsType{
	SERVICEDEFINITIONV2MSTEAMSTYPE_MICROSOFT_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2MSTeamsType) GetAllowedValues() []ServiceDefinitionV2MSTeamsType {
	return allowedServiceDefinitionV2MSTeamsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2MSTeamsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2MSTeamsType(value)
	return nil
}

// NewServiceDefinitionV2MSTeamsTypeFromValue returns a pointer to a valid ServiceDefinitionV2MSTeamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2MSTeamsTypeFromValue(v string) (*ServiceDefinitionV2MSTeamsType, error) {
	ev := ServiceDefinitionV2MSTeamsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2MSTeamsType: valid values are %v", v, allowedServiceDefinitionV2MSTeamsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2MSTeamsType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2MSTeamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2MSTeamsType value.
func (v ServiceDefinitionV2MSTeamsType) Ptr() *ServiceDefinitionV2MSTeamsType {
	return &v
}
<<<<<<< HEAD

// NullableServiceDefinitionV2MSTeamsType handles when a null is used for ServiceDefinitionV2MSTeamsType.
type NullableServiceDefinitionV2MSTeamsType struct {
	value *ServiceDefinitionV2MSTeamsType
	isSet bool
}

// Get returns the associated value.
func (v NullableServiceDefinitionV2MSTeamsType) Get() *ServiceDefinitionV2MSTeamsType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableServiceDefinitionV2MSTeamsType) Set(val *ServiceDefinitionV2MSTeamsType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableServiceDefinitionV2MSTeamsType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableServiceDefinitionV2MSTeamsType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDefinitionV2MSTeamsType initializes the struct as if Set has been called.
func NewNullableServiceDefinitionV2MSTeamsType(val *ServiceDefinitionV2MSTeamsType) *NullableServiceDefinitionV2MSTeamsType {
	return &NullableServiceDefinitionV2MSTeamsType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableServiceDefinitionV2MSTeamsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableServiceDefinitionV2MSTeamsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
