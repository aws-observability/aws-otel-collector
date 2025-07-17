// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUserType Users resource type.
type EscalationPolicyUserType string

// List of EscalationPolicyUserType.
const (
	ESCALATIONPOLICYUSERTYPE_USERS EscalationPolicyUserType = "users"
)

var allowedEscalationPolicyUserTypeEnumValues = []EscalationPolicyUserType{
	ESCALATIONPOLICYUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyUserType) GetAllowedValues() []EscalationPolicyUserType {
	return allowedEscalationPolicyUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyUserType(value)
	return nil
}

// NewEscalationPolicyUserTypeFromValue returns a pointer to a valid EscalationPolicyUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyUserTypeFromValue(v string) (*EscalationPolicyUserType, error) {
	ev := EscalationPolicyUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyUserType: valid values are %v", v, allowedEscalationPolicyUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyUserType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyUserType value.
func (v EscalationPolicyUserType) Ptr() *EscalationPolicyUserType {
	return &v
}
