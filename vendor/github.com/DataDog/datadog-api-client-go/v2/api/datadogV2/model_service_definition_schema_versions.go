// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionSchemaVersions Schema versions
type ServiceDefinitionSchemaVersions string

// List of ServiceDefinitionSchemaVersions.
const (
	SERVICEDEFINITIONSCHEMAVERSIONS_V1   ServiceDefinitionSchemaVersions = "v1"
	SERVICEDEFINITIONSCHEMAVERSIONS_V2   ServiceDefinitionSchemaVersions = "v2"
	SERVICEDEFINITIONSCHEMAVERSIONS_V2_1 ServiceDefinitionSchemaVersions = "v2.1"
	SERVICEDEFINITIONSCHEMAVERSIONS_V2_2 ServiceDefinitionSchemaVersions = "v2.2"
)

var allowedServiceDefinitionSchemaVersionsEnumValues = []ServiceDefinitionSchemaVersions{
	SERVICEDEFINITIONSCHEMAVERSIONS_V1,
	SERVICEDEFINITIONSCHEMAVERSIONS_V2,
	SERVICEDEFINITIONSCHEMAVERSIONS_V2_1,
	SERVICEDEFINITIONSCHEMAVERSIONS_V2_2,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionSchemaVersions) GetAllowedValues() []ServiceDefinitionSchemaVersions {
	return allowedServiceDefinitionSchemaVersionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionSchemaVersions) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionSchemaVersions(value)
	return nil
}

// NewServiceDefinitionSchemaVersionsFromValue returns a pointer to a valid ServiceDefinitionSchemaVersions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionSchemaVersionsFromValue(v string) (*ServiceDefinitionSchemaVersions, error) {
	ev := ServiceDefinitionSchemaVersions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionSchemaVersions: valid values are %v", v, allowedServiceDefinitionSchemaVersionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionSchemaVersions) IsValid() bool {
	for _, existing := range allowedServiceDefinitionSchemaVersionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionSchemaVersions value.
func (v ServiceDefinitionSchemaVersions) Ptr() *ServiceDefinitionSchemaVersions {
	return &v
}
