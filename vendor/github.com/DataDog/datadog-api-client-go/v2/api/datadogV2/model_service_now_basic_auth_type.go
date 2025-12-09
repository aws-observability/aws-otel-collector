// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowBasicAuthType The definition of the `ServiceNowBasicAuth` object.
type ServiceNowBasicAuthType string

// List of ServiceNowBasicAuthType.
const (
	SERVICENOWBASICAUTHTYPE_SERVICENOWBASICAUTH ServiceNowBasicAuthType = "ServiceNowBasicAuth"
)

var allowedServiceNowBasicAuthTypeEnumValues = []ServiceNowBasicAuthType{
	SERVICENOWBASICAUTHTYPE_SERVICENOWBASICAUTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceNowBasicAuthType) GetAllowedValues() []ServiceNowBasicAuthType {
	return allowedServiceNowBasicAuthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceNowBasicAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceNowBasicAuthType(value)
	return nil
}

// NewServiceNowBasicAuthTypeFromValue returns a pointer to a valid ServiceNowBasicAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceNowBasicAuthTypeFromValue(v string) (*ServiceNowBasicAuthType, error) {
	ev := ServiceNowBasicAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceNowBasicAuthType: valid values are %v", v, allowedServiceNowBasicAuthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceNowBasicAuthType) IsValid() bool {
	for _, existing := range allowedServiceNowBasicAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceNowBasicAuthType value.
func (v ServiceNowBasicAuthType) Ptr() *ServiceNowBasicAuthType {
	return &v
}
