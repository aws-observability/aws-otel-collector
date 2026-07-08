// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityServiceType The type of the resource. The value should always be `service_env`.
type ApplicationSecurityServiceType string

// List of ApplicationSecurityServiceType.
const (
	APPLICATIONSECURITYSERVICETYPE_SERVICE_ENV ApplicationSecurityServiceType = "service_env"
)

var allowedApplicationSecurityServiceTypeEnumValues = []ApplicationSecurityServiceType{
	APPLICATIONSECURITYSERVICETYPE_SERVICE_ENV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApplicationSecurityServiceType) GetAllowedValues() []ApplicationSecurityServiceType {
	return allowedApplicationSecurityServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApplicationSecurityServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApplicationSecurityServiceType(value)
	return nil
}

// NewApplicationSecurityServiceTypeFromValue returns a pointer to a valid ApplicationSecurityServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApplicationSecurityServiceTypeFromValue(v string) (*ApplicationSecurityServiceType, error) {
	ev := ApplicationSecurityServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApplicationSecurityServiceType: valid values are %v", v, allowedApplicationSecurityServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApplicationSecurityServiceType) IsValid() bool {
	for _, existing := range allowedApplicationSecurityServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationSecurityServiceType value.
func (v ApplicationSecurityServiceType) Ptr() *ApplicationSecurityServiceType {
	return &v
}
