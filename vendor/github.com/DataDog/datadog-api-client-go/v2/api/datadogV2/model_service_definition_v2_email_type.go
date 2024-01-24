// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2EmailType Contact type.
type ServiceDefinitionV2EmailType string

// List of ServiceDefinitionV2EmailType.
const (
	SERVICEDEFINITIONV2EMAILTYPE_EMAIL ServiceDefinitionV2EmailType = "email"
)

var allowedServiceDefinitionV2EmailTypeEnumValues = []ServiceDefinitionV2EmailType{
	SERVICEDEFINITIONV2EMAILTYPE_EMAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2EmailType) GetAllowedValues() []ServiceDefinitionV2EmailType {
	return allowedServiceDefinitionV2EmailTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2EmailType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2EmailType(value)
	return nil
}

// NewServiceDefinitionV2EmailTypeFromValue returns a pointer to a valid ServiceDefinitionV2EmailType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2EmailTypeFromValue(v string) (*ServiceDefinitionV2EmailType, error) {
	ev := ServiceDefinitionV2EmailType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2EmailType: valid values are %v", v, allowedServiceDefinitionV2EmailTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2EmailType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2EmailTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2EmailType value.
func (v ServiceDefinitionV2EmailType) Ptr() *ServiceDefinitionV2EmailType {
	return &v
}
