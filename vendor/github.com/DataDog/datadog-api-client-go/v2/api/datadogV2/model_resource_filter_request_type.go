// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResourceFilterRequestType Constant string to identify the request type.
type ResourceFilterRequestType string

// List of ResourceFilterRequestType.
const (
	RESOURCEFILTERREQUESTTYPE_CSM_RESOURCE_FILTER ResourceFilterRequestType = "csm_resource_filter"
)

var allowedResourceFilterRequestTypeEnumValues = []ResourceFilterRequestType{
	RESOURCEFILTERREQUESTTYPE_CSM_RESOURCE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ResourceFilterRequestType) GetAllowedValues() []ResourceFilterRequestType {
	return allowedResourceFilterRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ResourceFilterRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ResourceFilterRequestType(value)
	return nil
}

// NewResourceFilterRequestTypeFromValue returns a pointer to a valid ResourceFilterRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewResourceFilterRequestTypeFromValue(v string) (*ResourceFilterRequestType, error) {
	ev := ResourceFilterRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ResourceFilterRequestType: valid values are %v", v, allowedResourceFilterRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ResourceFilterRequestType) IsValid() bool {
	for _, existing := range allowedResourceFilterRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceFilterRequestType value.
func (v ResourceFilterRequestType) Ptr() *ResourceFilterRequestType {
	return &v
}
