// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesUserType Users resource type.
type StatusPagesUserType string

// List of StatusPagesUserType.
const (
	STATUSPAGESUSERTYPE_USERS StatusPagesUserType = "users"
)

var allowedStatusPagesUserTypeEnumValues = []StatusPagesUserType{
	STATUSPAGESUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatusPagesUserType) GetAllowedValues() []StatusPagesUserType {
	return allowedStatusPagesUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatusPagesUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatusPagesUserType(value)
	return nil
}

// NewStatusPagesUserTypeFromValue returns a pointer to a valid StatusPagesUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatusPagesUserTypeFromValue(v string) (*StatusPagesUserType, error) {
	ev := StatusPagesUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatusPagesUserType: valid values are %v", v, allowedStatusPagesUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatusPagesUserType) IsValid() bool {
	for _, existing := range allowedStatusPagesUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusPagesUserType value.
func (v StatusPagesUserType) Ptr() *StatusPagesUserType {
	return &v
}
