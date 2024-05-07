// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserResourceType User resource type.
type UserResourceType string

// List of UserResourceType.
const (
	USERRESOURCETYPE_USER UserResourceType = "user"
)

var allowedUserResourceTypeEnumValues = []UserResourceType{
	USERRESOURCETYPE_USER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserResourceType) GetAllowedValues() []UserResourceType {
	return allowedUserResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserResourceType(value)
	return nil
}

// NewUserResourceTypeFromValue returns a pointer to a valid UserResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserResourceTypeFromValue(v string) (*UserResourceType, error) {
	ev := UserResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserResourceType: valid values are %v", v, allowedUserResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserResourceType) IsValid() bool {
	for _, existing := range allowedUserResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserResourceType value.
func (v UserResourceType) Ptr() *UserResourceType {
	return &v
}
