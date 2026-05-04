// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuardrailTriggerAction Action to perform when a guardrail threshold is triggered.
type GuardrailTriggerAction string

// List of GuardrailTriggerAction.
const (
	GUARDRAILTRIGGERACTION_PAUSE GuardrailTriggerAction = "PAUSE"
	GUARDRAILTRIGGERACTION_ABORT GuardrailTriggerAction = "ABORT"
)

var allowedGuardrailTriggerActionEnumValues = []GuardrailTriggerAction{
	GUARDRAILTRIGGERACTION_PAUSE,
	GUARDRAILTRIGGERACTION_ABORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuardrailTriggerAction) GetAllowedValues() []GuardrailTriggerAction {
	return allowedGuardrailTriggerActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuardrailTriggerAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuardrailTriggerAction(value)
	return nil
}

// NewGuardrailTriggerActionFromValue returns a pointer to a valid GuardrailTriggerAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuardrailTriggerActionFromValue(v string) (*GuardrailTriggerAction, error) {
	ev := GuardrailTriggerAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuardrailTriggerAction: valid values are %v", v, allowedGuardrailTriggerActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuardrailTriggerAction) IsValid() bool {
	for _, existing := range allowedGuardrailTriggerActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuardrailTriggerAction value.
func (v GuardrailTriggerAction) Ptr() *GuardrailTriggerAction {
	return &v
}
