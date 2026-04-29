// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnonymizeUsersRequestType Type of the anonymize users request.
type AnonymizeUsersRequestType string

// List of AnonymizeUsersRequestType.
const (
	ANONYMIZEUSERSREQUESTTYPE_ANONYMIZE_USERS_REQUEST AnonymizeUsersRequestType = "anonymize_users_request"
)

var allowedAnonymizeUsersRequestTypeEnumValues = []AnonymizeUsersRequestType{
	ANONYMIZEUSERSREQUESTTYPE_ANONYMIZE_USERS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnonymizeUsersRequestType) GetAllowedValues() []AnonymizeUsersRequestType {
	return allowedAnonymizeUsersRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnonymizeUsersRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnonymizeUsersRequestType(value)
	return nil
}

// NewAnonymizeUsersRequestTypeFromValue returns a pointer to a valid AnonymizeUsersRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnonymizeUsersRequestTypeFromValue(v string) (*AnonymizeUsersRequestType, error) {
	ev := AnonymizeUsersRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnonymizeUsersRequestType: valid values are %v", v, allowedAnonymizeUsersRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnonymizeUsersRequestType) IsValid() bool {
	for _, existing := range allowedAnonymizeUsersRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnonymizeUsersRequestType value.
func (v AnonymizeUsersRequestType) Ptr() *AnonymizeUsersRequestType {
	return &v
}
