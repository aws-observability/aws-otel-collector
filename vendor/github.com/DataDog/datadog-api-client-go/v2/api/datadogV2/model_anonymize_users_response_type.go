// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnonymizeUsersResponseType Type of the anonymize users response.
type AnonymizeUsersResponseType string

// List of AnonymizeUsersResponseType.
const (
	ANONYMIZEUSERSRESPONSETYPE_ANONYMIZE_USERS_RESPONSE AnonymizeUsersResponseType = "anonymize_users_response"
)

var allowedAnonymizeUsersResponseTypeEnumValues = []AnonymizeUsersResponseType{
	ANONYMIZEUSERSRESPONSETYPE_ANONYMIZE_USERS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnonymizeUsersResponseType) GetAllowedValues() []AnonymizeUsersResponseType {
	return allowedAnonymizeUsersResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnonymizeUsersResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnonymizeUsersResponseType(value)
	return nil
}

// NewAnonymizeUsersResponseTypeFromValue returns a pointer to a valid AnonymizeUsersResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnonymizeUsersResponseTypeFromValue(v string) (*AnonymizeUsersResponseType, error) {
	ev := AnonymizeUsersResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnonymizeUsersResponseType: valid values are %v", v, allowedAnonymizeUsersResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnonymizeUsersResponseType) IsValid() bool {
	for _, existing := range allowedAnonymizeUsersResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnonymizeUsersResponseType value.
func (v AnonymizeUsersResponseType) Ptr() *AnonymizeUsersResponseType {
	return &v
}
