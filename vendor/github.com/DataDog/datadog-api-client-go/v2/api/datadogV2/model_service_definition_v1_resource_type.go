// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV1ResourceType Link type.
type ServiceDefinitionV1ResourceType string

// List of ServiceDefinitionV1ResourceType.
const (
	SERVICEDEFINITIONV1RESOURCETYPE_DOC       ServiceDefinitionV1ResourceType = "doc"
	SERVICEDEFINITIONV1RESOURCETYPE_WIKI      ServiceDefinitionV1ResourceType = "wiki"
	SERVICEDEFINITIONV1RESOURCETYPE_RUNBOOK   ServiceDefinitionV1ResourceType = "runbook"
	SERVICEDEFINITIONV1RESOURCETYPE_URL       ServiceDefinitionV1ResourceType = "url"
	SERVICEDEFINITIONV1RESOURCETYPE_REPO      ServiceDefinitionV1ResourceType = "repo"
	SERVICEDEFINITIONV1RESOURCETYPE_DASHBOARD ServiceDefinitionV1ResourceType = "dashboard"
	SERVICEDEFINITIONV1RESOURCETYPE_ONCALL    ServiceDefinitionV1ResourceType = "oncall"
	SERVICEDEFINITIONV1RESOURCETYPE_CODE      ServiceDefinitionV1ResourceType = "code"
	SERVICEDEFINITIONV1RESOURCETYPE_LINK      ServiceDefinitionV1ResourceType = "link"
)

var allowedServiceDefinitionV1ResourceTypeEnumValues = []ServiceDefinitionV1ResourceType{
	SERVICEDEFINITIONV1RESOURCETYPE_DOC,
	SERVICEDEFINITIONV1RESOURCETYPE_WIKI,
	SERVICEDEFINITIONV1RESOURCETYPE_RUNBOOK,
	SERVICEDEFINITIONV1RESOURCETYPE_URL,
	SERVICEDEFINITIONV1RESOURCETYPE_REPO,
	SERVICEDEFINITIONV1RESOURCETYPE_DASHBOARD,
	SERVICEDEFINITIONV1RESOURCETYPE_ONCALL,
	SERVICEDEFINITIONV1RESOURCETYPE_CODE,
	SERVICEDEFINITIONV1RESOURCETYPE_LINK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV1ResourceType) GetAllowedValues() []ServiceDefinitionV1ResourceType {
	return allowedServiceDefinitionV1ResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV1ResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV1ResourceType(value)
	return nil
}

// NewServiceDefinitionV1ResourceTypeFromValue returns a pointer to a valid ServiceDefinitionV1ResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV1ResourceTypeFromValue(v string) (*ServiceDefinitionV1ResourceType, error) {
	ev := ServiceDefinitionV1ResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV1ResourceType: valid values are %v", v, allowedServiceDefinitionV1ResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV1ResourceType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV1ResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV1ResourceType value.
func (v ServiceDefinitionV1ResourceType) Ptr() *ServiceDefinitionV1ResourceType {
	return &v
}
