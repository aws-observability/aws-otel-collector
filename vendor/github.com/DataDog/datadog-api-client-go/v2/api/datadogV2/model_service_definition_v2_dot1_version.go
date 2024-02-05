// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1Version Schema version being used.
type ServiceDefinitionV2Dot1Version string

// List of ServiceDefinitionV2Dot1Version.
const (
	SERVICEDEFINITIONV2DOT1VERSION_V2_1 ServiceDefinitionV2Dot1Version = "v2.1"
)

var allowedServiceDefinitionV2Dot1VersionEnumValues = []ServiceDefinitionV2Dot1Version{
	SERVICEDEFINITIONV2DOT1VERSION_V2_1,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1Version) GetAllowedValues() []ServiceDefinitionV2Dot1Version {
	return allowedServiceDefinitionV2Dot1VersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1Version) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1Version(value)
	return nil
}

// NewServiceDefinitionV2Dot1VersionFromValue returns a pointer to a valid ServiceDefinitionV2Dot1Version
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1VersionFromValue(v string) (*ServiceDefinitionV2Dot1Version, error) {
	ev := ServiceDefinitionV2Dot1Version(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1Version: valid values are %v", v, allowedServiceDefinitionV2Dot1VersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1Version) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1VersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1Version value.
func (v ServiceDefinitionV2Dot1Version) Ptr() *ServiceDefinitionV2Dot1Version {
	return &v
}
