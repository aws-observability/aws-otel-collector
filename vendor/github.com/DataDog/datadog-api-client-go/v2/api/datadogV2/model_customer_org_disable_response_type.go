// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomerOrgDisableResponseType JSON:API resource type for a customer org disable response.
type CustomerOrgDisableResponseType string

// List of CustomerOrgDisableResponseType.
const (
	CUSTOMERORGDISABLERESPONSETYPE_ORG_DISABLE CustomerOrgDisableResponseType = "org_disable"
)

var allowedCustomerOrgDisableResponseTypeEnumValues = []CustomerOrgDisableResponseType{
	CUSTOMERORGDISABLERESPONSETYPE_ORG_DISABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomerOrgDisableResponseType) GetAllowedValues() []CustomerOrgDisableResponseType {
	return allowedCustomerOrgDisableResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomerOrgDisableResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomerOrgDisableResponseType(value)
	return nil
}

// NewCustomerOrgDisableResponseTypeFromValue returns a pointer to a valid CustomerOrgDisableResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomerOrgDisableResponseTypeFromValue(v string) (*CustomerOrgDisableResponseType, error) {
	ev := CustomerOrgDisableResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomerOrgDisableResponseType: valid values are %v", v, allowedCustomerOrgDisableResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomerOrgDisableResponseType) IsValid() bool {
	for _, existing := range allowedCustomerOrgDisableResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerOrgDisableResponseType value.
func (v CustomerOrgDisableResponseType) Ptr() *CustomerOrgDisableResponseType {
	return &v
}
