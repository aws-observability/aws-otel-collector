// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowUserType Type identifier for ServiceNow user resources
type ServiceNowUserType string

// List of ServiceNowUserType.
const (
	SERVICENOWUSERTYPE_USERS ServiceNowUserType = "users"
)

var allowedServiceNowUserTypeEnumValues = []ServiceNowUserType{
	SERVICENOWUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceNowUserType) GetAllowedValues() []ServiceNowUserType {
	return allowedServiceNowUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceNowUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceNowUserType(value)
	return nil
}

// NewServiceNowUserTypeFromValue returns a pointer to a valid ServiceNowUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceNowUserTypeFromValue(v string) (*ServiceNowUserType, error) {
	ev := ServiceNowUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceNowUserType: valid values are %v", v, allowedServiceNowUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceNowUserType) IsValid() bool {
	for _, existing := range allowedServiceNowUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceNowUserType value.
func (v ServiceNowUserType) Ptr() *ServiceNowUserType {
	return &v
}
