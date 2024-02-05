// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV1Version Schema version being used.
type ServiceDefinitionV1Version string

// List of ServiceDefinitionV1Version.
const (
	SERVICEDEFINITIONV1VERSION_V1 ServiceDefinitionV1Version = "v1"
)

var allowedServiceDefinitionV1VersionEnumValues = []ServiceDefinitionV1Version{
	SERVICEDEFINITIONV1VERSION_V1,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV1Version) GetAllowedValues() []ServiceDefinitionV1Version {
	return allowedServiceDefinitionV1VersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV1Version) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV1Version(value)
	return nil
}

// NewServiceDefinitionV1VersionFromValue returns a pointer to a valid ServiceDefinitionV1Version
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV1VersionFromValue(v string) (*ServiceDefinitionV1Version, error) {
	ev := ServiceDefinitionV1Version(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV1Version: valid values are %v", v, allowedServiceDefinitionV1VersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV1Version) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV1VersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV1Version value.
func (v ServiceDefinitionV1Version) Ptr() *ServiceDefinitionV1Version {
	return &v
}
