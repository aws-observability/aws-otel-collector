// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomerOrgDisableStatus Resulting lifecycle status of the organization after the disable action.
type CustomerOrgDisableStatus string

// List of CustomerOrgDisableStatus.
const (
	CUSTOMERORGDISABLESTATUS_DISABLED        CustomerOrgDisableStatus = "disabled"
	CUSTOMERORGDISABLESTATUS_PENDING_DISABLE CustomerOrgDisableStatus = "pending_disable"
)

var allowedCustomerOrgDisableStatusEnumValues = []CustomerOrgDisableStatus{
	CUSTOMERORGDISABLESTATUS_DISABLED,
	CUSTOMERORGDISABLESTATUS_PENDING_DISABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomerOrgDisableStatus) GetAllowedValues() []CustomerOrgDisableStatus {
	return allowedCustomerOrgDisableStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomerOrgDisableStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomerOrgDisableStatus(value)
	return nil
}

// NewCustomerOrgDisableStatusFromValue returns a pointer to a valid CustomerOrgDisableStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomerOrgDisableStatusFromValue(v string) (*CustomerOrgDisableStatus, error) {
	ev := CustomerOrgDisableStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomerOrgDisableStatus: valid values are %v", v, allowedCustomerOrgDisableStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomerOrgDisableStatus) IsValid() bool {
	for _, existing := range allowedCustomerOrgDisableStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerOrgDisableStatus value.
func (v CustomerOrgDisableStatus) Ptr() *CustomerOrgDisableStatus {
	return &v
}
