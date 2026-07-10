// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceAccessTokensType Service access tokens resource type.
type ServiceAccessTokensType string

// List of ServiceAccessTokensType.
const (
	SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS ServiceAccessTokensType = "service_access_tokens"
)

var allowedServiceAccessTokensTypeEnumValues = []ServiceAccessTokensType{
	SERVICEACCESSTOKENSTYPE_SERVICE_ACCESS_TOKENS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceAccessTokensType) GetAllowedValues() []ServiceAccessTokensType {
	return allowedServiceAccessTokensTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceAccessTokensType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceAccessTokensType(value)
	return nil
}

// NewServiceAccessTokensTypeFromValue returns a pointer to a valid ServiceAccessTokensType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceAccessTokensTypeFromValue(v string) (*ServiceAccessTokensType, error) {
	ev := ServiceAccessTokensType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceAccessTokensType: valid values are %v", v, allowedServiceAccessTokensTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceAccessTokensType) IsValid() bool {
	for _, existing := range allowedServiceAccessTokensTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceAccessTokensType value.
func (v ServiceAccessTokensType) Ptr() *ServiceAccessTokensType {
	return &v
}
