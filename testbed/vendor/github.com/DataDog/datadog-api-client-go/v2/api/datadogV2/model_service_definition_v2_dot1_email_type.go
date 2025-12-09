// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1EmailType Contact type.
type ServiceDefinitionV2Dot1EmailType string

// List of ServiceDefinitionV2Dot1EmailType.
const (
	SERVICEDEFINITIONV2DOT1EMAILTYPE_EMAIL ServiceDefinitionV2Dot1EmailType = "email"
)

var allowedServiceDefinitionV2Dot1EmailTypeEnumValues = []ServiceDefinitionV2Dot1EmailType{
	SERVICEDEFINITIONV2DOT1EMAILTYPE_EMAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot1EmailType) GetAllowedValues() []ServiceDefinitionV2Dot1EmailType {
	return allowedServiceDefinitionV2Dot1EmailTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot1EmailType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot1EmailType(value)
	return nil
}

// NewServiceDefinitionV2Dot1EmailTypeFromValue returns a pointer to a valid ServiceDefinitionV2Dot1EmailType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot1EmailTypeFromValue(v string) (*ServiceDefinitionV2Dot1EmailType, error) {
	ev := ServiceDefinitionV2Dot1EmailType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot1EmailType: valid values are %v", v, allowedServiceDefinitionV2Dot1EmailTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot1EmailType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot1EmailTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot1EmailType value.
func (v ServiceDefinitionV2Dot1EmailType) Ptr() *ServiceDefinitionV2Dot1EmailType {
	return &v
}
