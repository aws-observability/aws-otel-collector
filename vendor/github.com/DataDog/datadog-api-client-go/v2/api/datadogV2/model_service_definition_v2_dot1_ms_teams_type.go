// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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
	err := datadog.Unmarshal(src, &value)
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
