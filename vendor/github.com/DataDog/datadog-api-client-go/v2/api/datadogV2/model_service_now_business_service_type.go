// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowBusinessServiceType Type identifier for ServiceNow business service resources
type ServiceNowBusinessServiceType string

// List of ServiceNowBusinessServiceType.
const (
	SERVICENOWBUSINESSSERVICETYPE_BUSINESS_SERVICES ServiceNowBusinessServiceType = "business_services"
)

var allowedServiceNowBusinessServiceTypeEnumValues = []ServiceNowBusinessServiceType{
	SERVICENOWBUSINESSSERVICETYPE_BUSINESS_SERVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceNowBusinessServiceType) GetAllowedValues() []ServiceNowBusinessServiceType {
	return allowedServiceNowBusinessServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceNowBusinessServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceNowBusinessServiceType(value)
	return nil
}

// NewServiceNowBusinessServiceTypeFromValue returns a pointer to a valid ServiceNowBusinessServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceNowBusinessServiceTypeFromValue(v string) (*ServiceNowBusinessServiceType, error) {
	ev := ServiceNowBusinessServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceNowBusinessServiceType: valid values are %v", v, allowedServiceNowBusinessServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceNowBusinessServiceType) IsValid() bool {
	for _, existing := range allowedServiceNowBusinessServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceNowBusinessServiceType value.
func (v ServiceNowBusinessServiceType) Ptr() *ServiceNowBusinessServiceType {
	return &v
}
