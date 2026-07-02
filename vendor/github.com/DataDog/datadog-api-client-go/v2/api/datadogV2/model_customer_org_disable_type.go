// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomerOrgDisableType JSON:API resource type for a customer org disable request.
type CustomerOrgDisableType string

// List of CustomerOrgDisableType.
const (
	CUSTOMERORGDISABLETYPE_CUSTOMER_ORG_DISABLE CustomerOrgDisableType = "customer_org_disable"
)

var allowedCustomerOrgDisableTypeEnumValues = []CustomerOrgDisableType{
	CUSTOMERORGDISABLETYPE_CUSTOMER_ORG_DISABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomerOrgDisableType) GetAllowedValues() []CustomerOrgDisableType {
	return allowedCustomerOrgDisableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomerOrgDisableType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomerOrgDisableType(value)
	return nil
}

// NewCustomerOrgDisableTypeFromValue returns a pointer to a valid CustomerOrgDisableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomerOrgDisableTypeFromValue(v string) (*CustomerOrgDisableType, error) {
	ev := CustomerOrgDisableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomerOrgDisableType: valid values are %v", v, allowedCustomerOrgDisableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomerOrgDisableType) IsValid() bool {
	for _, existing := range allowedCustomerOrgDisableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerOrgDisableType value.
func (v CustomerOrgDisableType) Ptr() *CustomerOrgDisableType {
	return &v
}
