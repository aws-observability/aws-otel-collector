// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// ServiceDefinitionV2Dot2Type The type of service.
type ServiceDefinitionV2Dot2Type string

// List of ServiceDefinitionV2Dot2Type.
const (
	SERVICEDEFINITIONV2DOT2TYPE_WEB      ServiceDefinitionV2Dot2Type = "web"
	SERVICEDEFINITIONV2DOT2TYPE_DB       ServiceDefinitionV2Dot2Type = "db"
	SERVICEDEFINITIONV2DOT2TYPE_CACHE    ServiceDefinitionV2Dot2Type = "cache"
	SERVICEDEFINITIONV2DOT2TYPE_FUNCTION ServiceDefinitionV2Dot2Type = "function"
	SERVICEDEFINITIONV2DOT2TYPE_BROSWER  ServiceDefinitionV2Dot2Type = "browser"
	SERVICEDEFINITIONV2DOT2TYPE_MOBILE   ServiceDefinitionV2Dot2Type = "mobile"
	SERVICEDEFINITIONV2DOT2TYPE_CUSTOM   ServiceDefinitionV2Dot2Type = "custom"
)

var allowedServiceDefinitionV2Dot2TypeEnumValues = []ServiceDefinitionV2Dot2Type{
	SERVICEDEFINITIONV2DOT2TYPE_WEB,
	SERVICEDEFINITIONV2DOT2TYPE_DB,
	SERVICEDEFINITIONV2DOT2TYPE_CACHE,
	SERVICEDEFINITIONV2DOT2TYPE_FUNCTION,
	SERVICEDEFINITIONV2DOT2TYPE_BROSWER,
	SERVICEDEFINITIONV2DOT2TYPE_MOBILE,
	SERVICEDEFINITIONV2DOT2TYPE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceDefinitionV2Dot2Type) GetAllowedValues() []ServiceDefinitionV2Dot2Type {
	return allowedServiceDefinitionV2Dot2TypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceDefinitionV2Dot2Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceDefinitionV2Dot2Type(value)
	return nil
}

// NewServiceDefinitionV2Dot2TypeFromValue returns a pointer to a valid ServiceDefinitionV2Dot2Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceDefinitionV2Dot2TypeFromValue(v string) (*ServiceDefinitionV2Dot2Type, error) {
	ev := ServiceDefinitionV2Dot2Type(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceDefinitionV2Dot2Type: valid values are %v", v, allowedServiceDefinitionV2Dot2TypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceDefinitionV2Dot2Type) IsValid() bool {
	for _, existing := range allowedServiceDefinitionV2Dot2TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceDefinitionV2Dot2Type value.
func (v ServiceDefinitionV2Dot2Type) Ptr() *ServiceDefinitionV2Dot2Type {
	return &v
}
