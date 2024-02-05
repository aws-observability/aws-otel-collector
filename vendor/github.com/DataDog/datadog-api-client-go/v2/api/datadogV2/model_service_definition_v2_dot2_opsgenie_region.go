// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot2OpsgenieRegion Opsgenie instance region.
type ServiceDefinitionV2Dot2OpsgenieRegion string

// List of ServiceDefinitionV2Dot2OpsgenieRegion.
const (
	SERVICEDEFINITIONV2DOT2OPSGENIEREGION_US ServiceDefinitionV2Dot2OpsgenieRegion = "US"
	SERVICEDEFINITIONV2DOT2OPSGENIEREGION_EU ServiceDefinitionV2Dot2OpsgenieRegion = "EU"
)

var allowedServiceDefinitionV2Dot2OpsgenieRegionEnumValues = []ServiceDefinitionV2Dot2OpsgenieRegion{
	SERVICEDEFINITIONV2DOT2OPSGENIEREGION_US,
	SERVICEDEFINITIONV2DOT2OPSGENIEREGION_EU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot2OpsgenieRegion) GetAllowedValues() []ServiceDefinitionV2Dot2OpsgenieRegion {
	return allowedServiceDefinitionV2Dot2OpsgenieRegionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot2OpsgenieRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot2OpsgenieRegion(value)
	return nil
}

// NewServiceDefinitionV2Dot2OpsgenieRegionFromValue returns a pointer to a valid ServiceDefinitionV2Dot2OpsgenieRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot2OpsgenieRegionFromValue(v string) (*ServiceDefinitionV2Dot2OpsgenieRegion, error) {
	ev := ServiceDefinitionV2Dot2OpsgenieRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot2OpsgenieRegion: valid values are %v", v, allowedServiceDefinitionV2Dot2OpsgenieRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot2OpsgenieRegion) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot2OpsgenieRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot2OpsgenieRegion value.
func (v ServiceDefinitionV2Dot2OpsgenieRegion) Ptr() *ServiceDefinitionV2Dot2OpsgenieRegion {
	return &v
}
