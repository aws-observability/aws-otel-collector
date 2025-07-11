// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationType Represents the resource type for individual steps in an escalation policy used during incident response.
type EscalationType string

// List of EscalationType.
const (
	ESCALATIONTYPE_ESCALATION_POLICY_STEPS EscalationType = "escalation_policy_steps"
)

var allowedEscalationTypeEnumValues = []EscalationType{
	ESCALATIONTYPE_ESCALATION_POLICY_STEPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationType) GetAllowedValues() []EscalationType {
	return allowedEscalationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationType(value)
	return nil
}

// NewEscalationTypeFromValue returns a pointer to a valid EscalationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationTypeFromValue(v string) (*EscalationType, error) {
	ev := EscalationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationType: valid values are %v", v, allowedEscalationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationType) IsValid() bool {
	for _, existing := range allowedEscalationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationType value.
func (v EscalationType) Ptr() *EscalationType {
	return &v
}
