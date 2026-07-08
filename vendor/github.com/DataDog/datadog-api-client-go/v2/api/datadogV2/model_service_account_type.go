// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceAccountType Service account resource type.
type ServiceAccountType string

// List of ServiceAccountType.
const (
	SERVICEACCOUNTTYPE_SERVICE_ACCOUNT ServiceAccountType = "service_account"
)

var allowedServiceAccountTypeEnumValues = []ServiceAccountType{
	SERVICEACCOUNTTYPE_SERVICE_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceAccountType) GetAllowedValues() []ServiceAccountType {
	return allowedServiceAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceAccountType(value)
	return nil
}

// NewServiceAccountTypeFromValue returns a pointer to a valid ServiceAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceAccountTypeFromValue(v string) (*ServiceAccountType, error) {
	ev := ServiceAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceAccountType: valid values are %v", v, allowedServiceAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceAccountType) IsValid() bool {
	for _, existing := range allowedServiceAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceAccountType value.
func (v ServiceAccountType) Ptr() *ServiceAccountType {
	return &v
}
