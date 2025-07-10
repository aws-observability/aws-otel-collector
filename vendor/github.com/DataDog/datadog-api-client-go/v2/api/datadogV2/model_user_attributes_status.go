// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAttributesStatus The user's status.
type UserAttributesStatus string

// List of UserAttributesStatus.
const (
	USERATTRIBUTESSTATUS_ACTIVE      UserAttributesStatus = "active"
	USERATTRIBUTESSTATUS_DEACTIVATED UserAttributesStatus = "deactivated"
	USERATTRIBUTESSTATUS_PENDING     UserAttributesStatus = "pending"
)

var allowedUserAttributesStatusEnumValues = []UserAttributesStatus{
	USERATTRIBUTESSTATUS_ACTIVE,
	USERATTRIBUTESSTATUS_DEACTIVATED,
	USERATTRIBUTESSTATUS_PENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserAttributesStatus) GetAllowedValues() []UserAttributesStatus {
	return allowedUserAttributesStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserAttributesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserAttributesStatus(value)
	return nil
}

// NewUserAttributesStatusFromValue returns a pointer to a valid UserAttributesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserAttributesStatusFromValue(v string) (*UserAttributesStatus, error) {
	ev := UserAttributesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserAttributesStatus: valid values are %v", v, allowedUserAttributesStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserAttributesStatus) IsValid() bool {
	for _, existing := range allowedUserAttributesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAttributesStatus value.
func (v UserAttributesStatus) Ptr() *UserAttributesStatus {
	return &v
}
