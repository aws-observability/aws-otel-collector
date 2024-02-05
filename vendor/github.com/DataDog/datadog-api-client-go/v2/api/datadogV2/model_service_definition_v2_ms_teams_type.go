// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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
	err := datadog.Unmarshal(src, &value)
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
