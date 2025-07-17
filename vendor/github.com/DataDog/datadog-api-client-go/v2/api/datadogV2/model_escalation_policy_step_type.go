// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStepType Indicates that the resource is of type `steps`.
type EscalationPolicyStepType string

// List of EscalationPolicyStepType.
const (
	ESCALATIONPOLICYSTEPTYPE_STEPS EscalationPolicyStepType = "steps"
)

var allowedEscalationPolicyStepTypeEnumValues = []EscalationPolicyStepType{
	ESCALATIONPOLICYSTEPTYPE_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyStepType) GetAllowedValues() []EscalationPolicyStepType {
	return allowedEscalationPolicyStepTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyStepType(value)
	return nil
}

// NewEscalationPolicyStepTypeFromValue returns a pointer to a valid EscalationPolicyStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyStepTypeFromValue(v string) (*EscalationPolicyStepType, error) {
	ev := EscalationPolicyStepType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyStepType: valid values are %v", v, allowedEscalationPolicyStepTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyStepType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyStepType value.
func (v EscalationPolicyStepType) Ptr() *EscalationPolicyStepType {
	return &v
}
