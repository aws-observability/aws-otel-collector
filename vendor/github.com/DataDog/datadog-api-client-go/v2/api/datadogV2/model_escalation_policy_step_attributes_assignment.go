// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepAttributesAssignment Specifies how this escalation step will assign targets (example `default` or `round-robin`).
type EscalationPolicyStepAttributesAssignment string

// List of EscalationPolicyStepAttributesAssignment.
const (
	ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_DEFAULT     EscalationPolicyStepAttributesAssignment = "default"
	ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_ROUND_ROBIN EscalationPolicyStepAttributesAssignment = "round-robin"
)

var allowedEscalationPolicyStepAttributesAssignmentEnumValues = []EscalationPolicyStepAttributesAssignment{
	ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_DEFAULT,
	ESCALATIONPOLICYSTEPATTRIBUTESASSIGNMENT_ROUND_ROBIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyStepAttributesAssignment) GetAllowedValues() []EscalationPolicyStepAttributesAssignment {
	return allowedEscalationPolicyStepAttributesAssignmentEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyStepAttributesAssignment) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyStepAttributesAssignment(value)
	return nil
}

// NewEscalationPolicyStepAttributesAssignmentFromValue returns a pointer to a valid EscalationPolicyStepAttributesAssignment
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyStepAttributesAssignmentFromValue(v string) (*EscalationPolicyStepAttributesAssignment, error) {
	ev := EscalationPolicyStepAttributesAssignment(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyStepAttributesAssignment: valid values are %v", v, allowedEscalationPolicyStepAttributesAssignmentEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyStepAttributesAssignment) IsValid() bool {
	for _, existing := range allowedEscalationPolicyStepAttributesAssignmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyStepAttributesAssignment value.
func (v EscalationPolicyStepAttributesAssignment) Ptr() *EscalationPolicyStepAttributesAssignment {
	return &v
}
