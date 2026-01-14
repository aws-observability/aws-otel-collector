// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceListDataType Services list resource type.
type ServiceListDataType string

// List of ServiceListDataType.
const (
	SERVICELISTDATATYPE_SERVICES_LIST ServiceListDataType = "services_list"
)

var allowedServiceListDataTypeEnumValues = []ServiceListDataType{
	SERVICELISTDATATYPE_SERVICES_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ServiceListDataType) GetAllowedValues() []ServiceListDataType {
	return allowedServiceListDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ServiceListDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ServiceListDataType(value)
	return nil
}

// NewServiceListDataTypeFromValue returns a pointer to a valid ServiceListDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewServiceListDataTypeFromValue(v string) (*ServiceListDataType, error) {
	ev := ServiceListDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ServiceListDataType: valid values are %v", v, allowedServiceListDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ServiceListDataType) IsValid() bool {
	for _, existing := range allowedServiceListDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceListDataType value.
func (v ServiceListDataType) Ptr() *ServiceListDataType {
	return &v
}
