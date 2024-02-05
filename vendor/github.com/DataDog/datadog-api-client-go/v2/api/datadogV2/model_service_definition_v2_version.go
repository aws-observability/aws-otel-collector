// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Version Schema version being used.
type ServiceDefinitionV2Version string

// List of ServiceDefinitionV2Version.
const (
	SERVICEDEFINITIONV2VERSION_V2 ServiceDefinitionV2Version = "v2"
)

var allowedServiceDefinitionV2VersionEnumValues = []ServiceDefinitionV2Version{
	SERVICEDEFINITIONV2VERSION_V2,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Version) GetAllowedValues() []ServiceDefinitionV2Version {
	return allowedServiceDefinitionV2VersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Version) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Version(value)
	return nil
}

// NewServiceDefinitionV2VersionFromValue returns a pointer to a valid ServiceDefinitionV2Version
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2VersionFromValue(v string) (*ServiceDefinitionV2Version, error) {
	ev := ServiceDefinitionV2Version(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Version: valid values are %v", v, allowedServiceDefinitionV2VersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Version) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2VersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Version value.
func (v ServiceDefinitionV2Version) Ptr() *ServiceDefinitionV2Version {
	return &v
}
