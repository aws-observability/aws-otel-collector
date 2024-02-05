// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1SlackType Contact type.
type ServiceDefinitionV2Dot1SlackType string

// List of ServiceDefinitionV2Dot1SlackType.
const (
	SERVICEDEFINITIONV2DOT1SLACKTYPE_SLACK ServiceDefinitionV2Dot1SlackType = "slack"
)

var allowedServiceDefinitionV2Dot1SlackTypeEnumValues = []ServiceDefinitionV2Dot1SlackType{
	SERVICEDEFINITIONV2DOT1SLACKTYPE_SLACK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1SlackType) GetAllowedValues() []ServiceDefinitionV2Dot1SlackType {
	return allowedServiceDefinitionV2Dot1SlackTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1SlackType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1SlackType(value)
	return nil
}

// NewServiceDefinitionV2Dot1SlackTypeFromValue returns a pointer to a valid ServiceDefinitionV2Dot1SlackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1SlackTypeFromValue(v string) (*ServiceDefinitionV2Dot1SlackType, error) {
	ev := ServiceDefinitionV2Dot1SlackType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1SlackType: valid values are %v", v, allowedServiceDefinitionV2Dot1SlackTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1SlackType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1SlackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1SlackType value.
func (v ServiceDefinitionV2Dot1SlackType) Ptr() *ServiceDefinitionV2Dot1SlackType {
	return &v
}
