// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2LinkType Link type.
type ServiceDefinitionV2LinkType string

// List of ServiceDefinitionV2LinkType.
const (
	SERVICEDEFINITIONV2LINKTYPE_DOC       ServiceDefinitionV2LinkType = "doc"
	SERVICEDEFINITIONV2LINKTYPE_WIKI      ServiceDefinitionV2LinkType = "wiki"
	SERVICEDEFINITIONV2LINKTYPE_RUNBOOK   ServiceDefinitionV2LinkType = "runbook"
	SERVICEDEFINITIONV2LINKTYPE_URL       ServiceDefinitionV2LinkType = "url"
	SERVICEDEFINITIONV2LINKTYPE_REPO      ServiceDefinitionV2LinkType = "repo"
	SERVICEDEFINITIONV2LINKTYPE_DASHBOARD ServiceDefinitionV2LinkType = "dashboard"
	SERVICEDEFINITIONV2LINKTYPE_ONCALL    ServiceDefinitionV2LinkType = "oncall"
	SERVICEDEFINITIONV2LINKTYPE_CODE      ServiceDefinitionV2LinkType = "code"
	SERVICEDEFINITIONV2LINKTYPE_LINK      ServiceDefinitionV2LinkType = "link"
)

var allowedServiceDefinitionV2LinkTypeEnumValues = []ServiceDefinitionV2LinkType{
	SERVICEDEFINITIONV2LINKTYPE_DOC,
	SERVICEDEFINITIONV2LINKTYPE_WIKI,
	SERVICEDEFINITIONV2LINKTYPE_RUNBOOK,
	SERVICEDEFINITIONV2LINKTYPE_URL,
	SERVICEDEFINITIONV2LINKTYPE_REPO,
	SERVICEDEFINITIONV2LINKTYPE_DASHBOARD,
	SERVICEDEFINITIONV2LINKTYPE_ONCALL,
	SERVICEDEFINITIONV2LINKTYPE_CODE,
	SERVICEDEFINITIONV2LINKTYPE_LINK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2LinkType) GetAllowedValues() []ServiceDefinitionV2LinkType {
	return allowedServiceDefinitionV2LinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2LinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2LinkType(value)
	return nil
}

// NewServiceDefinitionV2LinkTypeFromValue returns a pointer to a valid ServiceDefinitionV2LinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2LinkTypeFromValue(v string) (*ServiceDefinitionV2LinkType, error) {
	ev := ServiceDefinitionV2LinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2LinkType: valid values are %v", v, allowedServiceDefinitionV2LinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2LinkType) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2LinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2LinkType value.
func (v ServiceDefinitionV2LinkType) Ptr() *ServiceDefinitionV2LinkType {
	return &v
}
