// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepTargetType Specifies the type of escalation target (example `users`, `schedules`, or `teams`).
type EscalationPolicyStepTargetType string

// List of EscalationPolicyStepTargetType.
const (
	ESCALATIONPOLICYSTEPTARGETTYPE_USERS     EscalationPolicyStepTargetType = "users"
	ESCALATIONPOLICYSTEPTARGETTYPE_SCHEDULES EscalationPolicyStepTargetType = "schedules"
	ESCALATIONPOLICYSTEPTARGETTYPE_TEAMS     EscalationPolicyStepTargetType = "teams"
)

var allowedEscalationPolicyStepTargetTypeEnumValues = []EscalationPolicyStepTargetType{
	ESCALATIONPOLICYSTEPTARGETTYPE_USERS,
	ESCALATIONPOLICYSTEPTARGETTYPE_SCHEDULES,
	ESCALATIONPOLICYSTEPTARGETTYPE_TEAMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyStepTargetType) GetAllowedValues() []EscalationPolicyStepTargetType {
	return allowedEscalationPolicyStepTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyStepTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyStepTargetType(value)
	return nil
}

// NewEscalationPolicyStepTargetTypeFromValue returns a pointer to a valid EscalationPolicyStepTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyStepTargetTypeFromValue(v string) (*EscalationPolicyStepTargetType, error) {
	ev := EscalationPolicyStepTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyStepTargetType: valid values are %v", v, allowedEscalationPolicyStepTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyStepTargetType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyStepTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyStepTargetType value.
func (v EscalationPolicyStepTargetType) Ptr() *EscalationPolicyStepTargetType {
	return &v
}
