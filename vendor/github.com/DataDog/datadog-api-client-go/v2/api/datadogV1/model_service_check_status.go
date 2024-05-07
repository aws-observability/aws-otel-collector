// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceCheckStatus The status of a service check. Set to `0` for OK, `1` for warning, `2` for critical, and `3` for unknown.
type ServiceCheckStatus int32

// List of ServiceCheckStatus.
const (
	SERVICECHECKSTATUS_OK       ServiceCheckStatus = 0
	SERVICECHECKSTATUS_WARNING  ServiceCheckStatus = 1
	SERVICECHECKSTATUS_CRITICAL ServiceCheckStatus = 2
	SERVICECHECKSTATUS_UNKNOWN  ServiceCheckStatus = 3
)

var allowedServiceCheckStatusEnumValues = []ServiceCheckStatus{
	SERVICECHECKSTATUS_OK,
	SERVICECHECKSTATUS_WARNING,
	SERVICECHECKSTATUS_CRITICAL,
	SERVICECHECKSTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceCheckStatus) GetAllowedValues() []ServiceCheckStatus {
	return allowedServiceCheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceCheckStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceCheckStatus(value)
	return nil
}

// NewServiceCheckStatusFromValue returns a pointer to a valid ServiceCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceCheckStatusFromValue(v int32) (*ServiceCheckStatus, error) {
	ev := ServiceCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceCheckStatus: valid values are %v", v, allowedServiceCheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceCheckStatus) IsValid() bool {
	for _, existing := range allowedServiceCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceCheckStatus value.
func (v ServiceCheckStatus) Ptr() *ServiceCheckStatus {
	return &v
}
