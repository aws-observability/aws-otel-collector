// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceRepositoryInfoDataType The resource type for service repository info objects.
type ServiceRepositoryInfoDataType string

// List of ServiceRepositoryInfoDataType.
const (
	SERVICEREPOSITORYINFODATATYPE_SERVICE_REPOSITORY_INFO ServiceRepositoryInfoDataType = "service_repository_info"
)

var allowedServiceRepositoryInfoDataTypeEnumValues = []ServiceRepositoryInfoDataType{
	SERVICEREPOSITORYINFODATATYPE_SERVICE_REPOSITORY_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceRepositoryInfoDataType) GetAllowedValues() []ServiceRepositoryInfoDataType {
	return allowedServiceRepositoryInfoDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceRepositoryInfoDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceRepositoryInfoDataType(value)
	return nil
}

// NewServiceRepositoryInfoDataTypeFromValue returns a pointer to a valid ServiceRepositoryInfoDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceRepositoryInfoDataTypeFromValue(v string) (*ServiceRepositoryInfoDataType, error) {
	ev := ServiceRepositoryInfoDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceRepositoryInfoDataType: valid values are %v", v, allowedServiceRepositoryInfoDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceRepositoryInfoDataType) IsValid() bool {
	for _, existing := range allowedServiceRepositoryInfoDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceRepositoryInfoDataType value.
func (v ServiceRepositoryInfoDataType) Ptr() *ServiceRepositoryInfoDataType {
	return &v
}
