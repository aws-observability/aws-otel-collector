// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// ServiceDefinitionV2Dot1OpsgenieRegion Opsgenie instance region.
type ServiceDefinitionV2Dot1OpsgenieRegion string

// List of ServiceDefinitionV2Dot1OpsgenieRegion.
const (
	SERVICEDEFINITIONV2DOT1OPSGENIEREGION_US ServiceDefinitionV2Dot1OpsgenieRegion = "US"
	SERVICEDEFINITIONV2DOT1OPSGENIEREGION_EU ServiceDefinitionV2Dot1OpsgenieRegion = "EU"
)

var allowedServiceDefinitionV2Dot1OpsgenieRegionEnumValues = []ServiceDefinitionV2Dot1OpsgenieRegion{
	SERVICEDEFINITIONV2DOT1OPSGENIEREGION_US,
	SERVICEDEFINITIONV2DOT1OPSGENIEREGION_EU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1OpsgenieRegion) GetAllowedValues() []ServiceDefinitionV2Dot1OpsgenieRegion {
	return allowedServiceDefinitionV2Dot1OpsgenieRegionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1OpsgenieRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1OpsgenieRegion(value)
	return nil
}

// NewServiceDefinitionV2Dot1OpsgenieRegionFromValue returns a pointer to a valid ServiceDefinitionV2Dot1OpsgenieRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1OpsgenieRegionFromValue(v string) (*ServiceDefinitionV2Dot1OpsgenieRegion, error) {
	ev := ServiceDefinitionV2Dot1OpsgenieRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1OpsgenieRegion: valid values are %v", v, allowedServiceDefinitionV2Dot1OpsgenieRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1OpsgenieRegion) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1OpsgenieRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1OpsgenieRegion value.
func (v ServiceDefinitionV2Dot1OpsgenieRegion) Ptr() *ServiceDefinitionV2Dot1OpsgenieRegion {
	return &v
}

// NullableServiceDefinitionV2Dot1OpsgenieRegion handles when a null is used for ServiceDefinitionV2Dot1OpsgenieRegion.
type NullableServiceDefinitionV2Dot1OpsgenieRegion struct {
	value *ServiceDefinitionV2Dot1OpsgenieRegion
	isSet bool
}

// Get returns the associated value.
func (v NullableServiceDefinitionV2Dot1OpsgenieRegion) Get() *ServiceDefinitionV2Dot1OpsgenieRegion {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableServiceDefinitionV2Dot1OpsgenieRegion) Set(val *ServiceDefinitionV2Dot1OpsgenieRegion) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableServiceDefinitionV2Dot1OpsgenieRegion) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableServiceDefinitionV2Dot1OpsgenieRegion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceDefinitionV2Dot1OpsgenieRegion initializes the struct as if Set has been called.
func NewNullableServiceDefinitionV2Dot1OpsgenieRegion(val *ServiceDefinitionV2Dot1OpsgenieRegion) *NullableServiceDefinitionV2Dot1OpsgenieRegion {
	return &NullableServiceDefinitionV2Dot1OpsgenieRegion{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableServiceDefinitionV2Dot1OpsgenieRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableServiceDefinitionV2Dot1OpsgenieRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
