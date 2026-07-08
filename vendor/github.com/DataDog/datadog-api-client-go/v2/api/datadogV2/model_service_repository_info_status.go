// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceRepositoryInfoStatus The status of the service repository info lookup.
type ServiceRepositoryInfoStatus string

// List of ServiceRepositoryInfoStatus.
const (
	SERVICEREPOSITORYINFOSTATUS_SUCCESS        ServiceRepositoryInfoStatus = "success"
	SERVICEREPOSITORYINFOSTATUS_NOT_FOUND      ServiceRepositoryInfoStatus = "not_found"
	SERVICEREPOSITORYINFOSTATUS_NO_REPOSITORY  ServiceRepositoryInfoStatus = "no_repository"
	SERVICEREPOSITORYINFOSTATUS_INTERNAL_ERROR ServiceRepositoryInfoStatus = "internal_error"
	SERVICEREPOSITORYINFOSTATUS_UNKNOWN        ServiceRepositoryInfoStatus = "unknown"
)

var allowedServiceRepositoryInfoStatusEnumValues = []ServiceRepositoryInfoStatus{
	SERVICEREPOSITORYINFOSTATUS_SUCCESS,
	SERVICEREPOSITORYINFOSTATUS_NOT_FOUND,
	SERVICEREPOSITORYINFOSTATUS_NO_REPOSITORY,
	SERVICEREPOSITORYINFOSTATUS_INTERNAL_ERROR,
	SERVICEREPOSITORYINFOSTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceRepositoryInfoStatus) GetAllowedValues() []ServiceRepositoryInfoStatus {
	return allowedServiceRepositoryInfoStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceRepositoryInfoStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceRepositoryInfoStatus(value)
	return nil
}

// NewServiceRepositoryInfoStatusFromValue returns a pointer to a valid ServiceRepositoryInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceRepositoryInfoStatusFromValue(v string) (*ServiceRepositoryInfoStatus, error) {
	ev := ServiceRepositoryInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceRepositoryInfoStatus: valid values are %v", v, allowedServiceRepositoryInfoStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceRepositoryInfoStatus) IsValid() bool {
	for _, existing := range allowedServiceRepositoryInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceRepositoryInfoStatus value.
func (v ServiceRepositoryInfoStatus) Ptr() *ServiceRepositoryInfoStatus {
	return &v
}
